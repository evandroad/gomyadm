package drivers

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"gomyadm/internal/models"

	_ "modernc.org/sqlite"
)

type SQLiteDriver struct{}

const sqliteRowIDKey = "__gomyadm_rowid"

func init() {
	Register("sqlite", SQLiteDriver{})
}

func (d SQLiteDriver) DriverName() string {
	return "sqlite"
}

func (d SQLiteDriver) BuildDSN(cfg models.ConnectionConfig) string {
	return (&url.URL{
		Scheme:   "file",
		Path:     cfg.FilePath,
		RawQuery: "mode=rw",
	}).String()
}

func (d SQLiteDriver) GetAllDatabase(db *sql.DB) ([]string, error) {
	var schemaVersion int
	if err := db.QueryRow("PRAGMA schema_version").Scan(&schemaVersion); err != nil {
		return nil, err
	}
	return []string{"main"}, nil
}

func (d SQLiteDriver) CreateDatabase(_ *sql.DB, _ string) error {
	return errors.New("sqlite does not support database creation; each database is a file")
}

func (d SQLiteDriver) UpdateDatabase(_ *sql.DB, _, _ string) error {
	return errors.New("sqlite does not support database renaming; each database is a file")
}

func (d SQLiteDriver) DeleteDatabase(_ *sql.DB, _ string) error {
	return errors.New("sqlite does not support database deletion; each database is a file")
}

func (d SQLiteDriver) GetAllTable(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`
		SELECT name
		FROM sqlite_master
		WHERE type = 'table' AND name NOT LIKE 'sqlite_%'
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}

func (d SQLiteDriver) CreateTable(db *sql.DB, table models.Table) error {
	if table.Name == "" {
		return errors.New("table name is required")
	}
	if len(table.Columns) == 0 {
		return errors.New("at least one column is required")
	}

	definitions := make([]string, 0, len(table.Columns)+1)
	primaryKeys := make([]string, 0)
	autoIncrementColumns := 0
	for _, column := range table.Columns {
		definition, err := buildSQLiteColumnDefinition(column, false)
		if err != nil {
			return fmt.Errorf("column %q: %w", column.Name, err)
		}
		if column.Primary {
			primaryKeys = append(primaryKeys, quoteSQLiteIdentifier(column.Name))
		}
		if column.AutoIncrement {
			autoIncrementColumns++
		}
		definitions = append(definitions, definition)
	}
	if len(primaryKeys) > 1 {
		if autoIncrementColumns > 0 {
			return errors.New("sqlite AUTOINCREMENT cannot be used with a composite primary key")
		}
		for i, column := range table.Columns {
			if column.Primary {
				definitions[i] = strings.Replace(definitions[i], " PRIMARY KEY", "", 1)
			}
		}
		definitions = append(definitions, "PRIMARY KEY ("+strings.Join(primaryKeys, ", ")+")")
	}

	query := fmt.Sprintf(
		"CREATE TABLE %s (%s)",
		quoteSQLiteIdentifier(table.Name),
		strings.Join(definitions, ", "),
	)
	_, err := db.Exec(query)
	return err
}

func (d SQLiteDriver) UpdateTable(db *sql.DB, oldName, newName string) error {
	if oldName == "" || newName == "" {
		return errors.New("old and new table names are required")
	}
	_, err := db.Exec(fmt.Sprintf(
		"ALTER TABLE %s RENAME TO %s",
		quoteSQLiteIdentifier(oldName),
		quoteSQLiteIdentifier(newName),
	))
	return err
}

func (d SQLiteDriver) DeleteTable(db *sql.DB, table string) error {
	if table == "" {
		return errors.New("table name is required")
	}
	_, err := db.Exec("DROP TABLE " + quoteSQLiteIdentifier(table))
	return err
}

func (d SQLiteDriver) GetAllItem(db *sql.DB, table string) (*models.TableData, error) {
	rowIDColumn, err := sqliteRowIDColumn(db, table)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s LIMIT 100", quoteSQLiteIdentifier(table))
	if rowIDColumn != "" {
		query = fmt.Sprintf(
			"SELECT %s AS %s, * FROM %s LIMIT 100",
			quoteSQLiteIdentifier(rowIDColumn),
			quoteSQLiteIdentifier(sqliteRowIDKey),
			quoteSQLiteIdentifier(table),
		)
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	visibleColumns := columns
	if rowIDColumn != "" {
		visibleColumns = columns[1:]
	}
	result := &models.TableData{
		Columns: visibleColumns,
		Rows:    make([]map[string]any, 0),
	}

	for rows.Next() {
		values := make([]any, len(columns))
		valuePointers := make([]any, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		row := make(map[string]any, len(columns))
		if rowIDColumn != "" {
			row[sqliteRowIDKey] = values[0]
		}
		offset := 0
		if rowIDColumn != "" {
			offset = 1
		}
		for i, column := range visibleColumns {
			value := values[i+offset]
			if bytes, ok := value.([]byte); ok {
				row[column] = string(bytes)
			} else {
				row[column] = value
			}
		}
		result.Rows = append(result.Rows, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (d SQLiteDriver) GetOneItem(db *sql.DB, table string, key map[string]any) (map[string]any, error) {
	where, args, err := buildSQLiteWhere(db, table, key)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE %s LIMIT 1",
		quoteSQLiteIdentifier(table),
		where,
	)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return nil, errors.New("item not found")
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]any, len(columns))
	valuePointers := make([]any, len(columns))
	for i := range values {
		valuePointers[i] = &values[i]
	}
	if err := rows.Scan(valuePointers...); err != nil {
		return nil, err
	}

	result := make(map[string]any, len(columns))
	for i, column := range columns {
		if value, ok := values[i].([]byte); ok {
			result[column] = string(value)
		} else {
			result[column] = values[i]
		}
	}
	return result, nil
}

func (d SQLiteDriver) CreateItem(db *sql.DB, table string, data map[string]any) error {
	if table == "" {
		return errors.New("table name is required")
	}
	if len(data) == 0 {
		_, err := db.Exec("INSERT INTO " + quoteSQLiteIdentifier(table) + " DEFAULT VALUES")
		return err
	}

	columns := sortedSQLiteKeys(data)
	quotedColumns := make([]string, len(columns))
	placeholders := make([]string, len(columns))
	values := make([]any, len(columns))
	for i, column := range columns {
		quotedColumns[i] = quoteSQLiteIdentifier(column)
		placeholders[i] = "?"
		values[i] = data[column]
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		quoteSQLiteIdentifier(table),
		strings.Join(quotedColumns, ", "),
		strings.Join(placeholders, ", "),
	)
	_, err := db.Exec(query, values...)
	return err
}

func (d SQLiteDriver) UpdateItem(db *sql.DB, table string, key, data map[string]any) error {
	if table == "" {
		return errors.New("table name is required")
	}
	if len(data) == 0 {
		return errors.New("at least one value is required to update an item")
	}
	where, keyArgs, err := buildSQLiteWhere(db, table, key)
	if err != nil {
		return err
	}

	columns := sortedSQLiteKeys(data)
	assignments := make([]string, len(columns))
	args := make([]any, 0, len(columns)+len(keyArgs))
	for i, column := range columns {
		assignments[i] = quoteSQLiteIdentifier(column) + " = ?"
		args = append(args, data[column])
	}
	args = append(args, keyArgs...)

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		quoteSQLiteIdentifier(table),
		strings.Join(assignments, ", "),
		where,
	)
	_, err = db.Exec(query, args...)
	return err
}

func (d SQLiteDriver) DeleteItem(db *sql.DB, table string, key map[string]any) error {
	if table == "" {
		return errors.New("table name is required")
	}
	where, args, err := buildSQLiteWhere(db, table, key)
	if err != nil {
		return err
	}
	_, err = db.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE %s", quoteSQLiteIdentifier(table), where),
		args...,
	)
	return err
}

func (d SQLiteDriver) GetAllColumn(db *sql.DB, table string) (*models.Table, error) {
	var tableDefinition sql.NullString
	if err := db.QueryRow(
		"SELECT sql FROM sqlite_master WHERE type = 'table' AND name = ?",
		table,
	).Scan(&tableDefinition); err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT name, type, "notnull", dflt_value, pk
		FROM pragma_table_info(?)
		ORDER BY cid
	`, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schema := &models.Table{Name: table, Columns: make([]models.Column, 0)}
	columnIndexes := make(map[string]int)
	for rows.Next() {
		var column models.Column
		var notNull, primary int
		var defaultValue sql.NullString

		if err := rows.Scan(&column.Name, &column.Type, &notNull, &defaultValue, &primary); err != nil {
			return nil, err
		}
		column.Nullable = notNull == 0 && primary == 0
		column.Primary = primary > 0
		column.AutoIncrement = column.Primary &&
			strings.EqualFold(column.Type, "INTEGER") &&
			tableDefinition.Valid &&
			strings.Contains(strings.ToUpper(tableDefinition.String), "AUTOINCREMENT")
		if defaultValue.Valid {
			column.DefaultValue = defaultValue.String
		}
		columnIndexes[column.Name] = len(schema.Columns)
		schema.Columns = append(schema.Columns, column)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	indexes, err := db.Query(`SELECT name FROM pragma_index_list(?) WHERE "unique" = 1`, table)
	if err != nil {
		return nil, err
	}
	defer indexes.Close()

	for indexes.Next() {
		var index string
		if err := indexes.Scan(&index); err != nil {
			return nil, err
		}

		indexColumns, err := db.Query(`SELECT name FROM pragma_index_info(?)`, index)
		if err != nil {
			return nil, err
		}
		var columnName string
		hasSingleColumn, hasMultipleColumns := false, false
		if indexColumns.Next() {
			if err := indexColumns.Scan(&columnName); err != nil {
				indexColumns.Close()
				return nil, err
			}
			hasSingleColumn = true
		}
		if indexColumns.Next() {
			hasMultipleColumns = true
		}
		if err := indexColumns.Err(); err != nil {
			indexColumns.Close()
			return nil, err
		}
		if err := indexColumns.Close(); err != nil {
			return nil, err
		}
		if hasSingleColumn && !hasMultipleColumns {
			if columnIndex, exists := columnIndexes[columnName]; exists {
				schema.Columns[columnIndex].Unique = true
			}
		}
	}
	if err := indexes.Err(); err != nil {
		return nil, err
	}

	return schema, nil
}

func (d SQLiteDriver) CreateColumn(db *sql.DB, table string, column models.Column) error {
	if table == "" {
		return errors.New("table name is required")
	}
	if column.Primary || column.AutoIncrement {
		return errors.New("sqlite cannot add a primary key or AUTOINCREMENT column without rebuilding the table")
	}
	definition, err := buildSQLiteColumnDefinition(column, true)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(fmt.Sprintf(
		"ALTER TABLE %s ADD COLUMN %s",
		quoteSQLiteIdentifier(table),
		definition,
	)); err != nil {
		return err
	}
	if column.Unique {
		indexName := sqliteUniqueIndexName(table, column.Name)
		if _, err := tx.Exec(fmt.Sprintf(
			"CREATE UNIQUE INDEX %s ON %s (%s)",
			quoteSQLiteIdentifier(indexName),
			quoteSQLiteIdentifier(table),
			quoteSQLiteIdentifier(column.Name),
		)); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (d SQLiteDriver) UpdateColumn(db *sql.DB, table, oldName string, column models.Column) error {
	if table == "" || oldName == "" || column.Name == "" {
		return errors.New("table and column names are required")
	}

	schema, err := d.GetAllColumn(db, table)
	if err != nil {
		return err
	}
	var existing *models.Column
	for i := range schema.Columns {
		if schema.Columns[i].Name == oldName {
			existing = &schema.Columns[i]
			break
		}
	}
	if existing == nil {
		return fmt.Errorf("column %q not found", oldName)
	}
	if existing.Type != column.Type ||
		existing.Nullable != column.Nullable ||
		existing.Primary != column.Primary ||
		existing.Unique != column.Unique ||
		existing.AutoIncrement != column.AutoIncrement ||
		existing.DefaultValue != column.DefaultValue {
		return errors.New("sqlite currently supports column renaming only; other column changes require rebuilding the table")
	}
	if oldName == column.Name {
		return nil
	}

	_, err = db.Exec(fmt.Sprintf(
		"ALTER TABLE %s RENAME COLUMN %s TO %s",
		quoteSQLiteIdentifier(table),
		quoteSQLiteIdentifier(oldName),
		quoteSQLiteIdentifier(column.Name),
	))
	return err
}

func (d SQLiteDriver) DeleteColumn(db *sql.DB, table, column string) error {
	if table == "" || column == "" {
		return errors.New("table and column names are required")
	}

	rows, err := db.Query(`
		SELECT DISTINCT il.name
		FROM pragma_index_list(?) AS il
		JOIN pragma_index_info(il.name) AS ii
		WHERE ii.name = ?
	`, table, column)
	if err != nil {
		return err
	}
	var indexes []string
	for rows.Next() {
		var index string
		if err := rows.Scan(&index); err != nil {
			rows.Close()
			return err
		}
		indexes = append(indexes, index)
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return err
	}
	if err := rows.Close(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, index := range indexes {
		if _, err := tx.Exec("DROP INDEX " + quoteSQLiteIdentifier(index)); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(fmt.Sprintf(
		"ALTER TABLE %s DROP COLUMN %s",
		quoteSQLiteIdentifier(table),
		quoteSQLiteIdentifier(column),
	)); err != nil {
		return err
	}
	return tx.Commit()
}

var sqliteTypePattern = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_ ]*(\([0-9]+(,[0-9]+)?\))?$`)
var sqliteNumberPattern = regexp.MustCompile(`^[+-]?[0-9]+(\.[0-9]+)?$`)

func buildSQLiteColumnDefinition(column models.Column, forAlter bool) (string, error) {
	if column.Name == "" {
		return "", errors.New("column name is required")
	}
	columnType := strings.TrimSpace(column.Type)
	if columnType == "" {
		columnType = "TEXT"
	}
	if !sqliteTypePattern.MatchString(columnType) {
		return "", fmt.Errorf("invalid SQLite type %q", column.Type)
	}
	if column.AutoIncrement && (!column.Primary || !strings.EqualFold(columnType, "INTEGER") || forAlter) {
		return "", errors.New("AUTOINCREMENT requires an INTEGER PRIMARY KEY and cannot be added to an existing table")
	}
	if forAlter && column.Primary {
		return "", errors.New("a primary key cannot be added to an existing table")
	}

	definition := quoteSQLiteIdentifier(column.Name) + " " + columnType
	if !column.Nullable {
		definition += " NOT NULL"
	}
	if column.Unique && !forAlter {
		definition += " UNIQUE"
	}
	if column.DefaultValue != "" {
		definition += " DEFAULT " + sqliteDefaultExpression(column.DefaultValue)
	}
	if column.Primary && !forAlter {
		definition += " PRIMARY KEY"
	}
	if column.AutoIncrement && !forAlter {
		definition += " AUTOINCREMENT"
	}
	return definition, nil
}

func sqliteDefaultExpression(value string) string {
	trimmed := strings.TrimSpace(value)
	switch strings.ToUpper(trimmed) {
	case "NULL", "TRUE", "FALSE", "CURRENT_TIMESTAMP", "CURRENT_DATE", "CURRENT_TIME":
		return strings.ToUpper(trimmed)
	}
	if sqliteNumberPattern.MatchString(trimmed) {
		return trimmed
	}
	return "'" + strings.ReplaceAll(trimmed, "'", "''") + "'"
}

func buildSQLiteWhere(db *sql.DB, table string, key map[string]any) (string, []any, error) {
	if len(key) == 0 {
		return "", nil, errors.New("at least one key is required")
	}
	columns := sortedSQLiteKeys(key)
	clauses := make([]string, len(columns))
	args := make([]any, len(columns))
	for i, column := range columns {
		identifier := column
		if column == sqliteRowIDKey {
			rowIDColumn, err := sqliteRowIDColumn(db, table)
			if err != nil {
				return "", nil, err
			}
			if rowIDColumn == "" {
				return "", nil, errors.New("table does not expose a SQLite rowid; use a primary key to edit rows")
			}
			identifier = rowIDColumn
		}
		clauses[i] = quoteSQLiteIdentifier(identifier) + " = ?"
		args[i] = key[column]
	}
	return strings.Join(clauses, " AND "), args, nil
}

func sqliteRowIDColumn(db *sql.DB, table string) (string, error) {
	rows, err := db.Query("SELECT name, pk FROM pragma_table_info(?)", table)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columns := make(map[string]bool)
	hasPrimaryKey := false
	for rows.Next() {
		var name string
		var primaryKeyPosition int
		if err := rows.Scan(&name, &primaryKeyPosition); err != nil {
			return "", err
		}
		columns[strings.ToLower(name)] = true
		hasPrimaryKey = hasPrimaryKey || primaryKeyPosition > 0
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	if hasPrimaryKey {
		return "", nil
	}

	for _, alias := range []string{"rowid", "_rowid_", "oid"} {
		if !columns[alias] {
			return alias, nil
		}
	}
	return "", nil
}

func sortedSQLiteKeys(values map[string]any) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sqliteUniqueIndexName(table, column string) string {
	sum := sha256.Sum256([]byte(table + "\x00" + column))
	return fmt.Sprintf("gomyadm_unique_%x", sum[:])
}

func quoteSQLiteIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}
