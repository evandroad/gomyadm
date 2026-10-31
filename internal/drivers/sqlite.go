package drivers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"gomyadm/internal/models"

	_ "modernc.org/sqlite"
)

type SQLiteDriver struct{}

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
		RawQuery: "mode=ro",
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

func (d SQLiteDriver) CreateTable(_ *sql.DB, _ models.Table) error {
	return errors.New("sqlite table creation is not supported yet")
}

func (d SQLiteDriver) UpdateTable(_ *sql.DB, _, _ string) error {
	return errors.New("sqlite table renaming is not supported yet")
}

func (d SQLiteDriver) DeleteTable(_ *sql.DB, _ string) error {
	return errors.New("sqlite table deletion is not supported yet")
}

func (d SQLiteDriver) GetAllItem(db *sql.DB, table string) (*models.TableData, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 100", quoteSQLiteIdentifier(table)))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := &models.TableData{
		Columns: columns,
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
		for i, column := range columns {
			if value, ok := values[i].([]byte); ok {
				row[column] = string(value)
			} else {
				row[column] = values[i]
			}
		}
		result.Rows = append(result.Rows, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (d SQLiteDriver) GetOneItem(_ *sql.DB, _ string, _ map[string]any) (map[string]any, error) {
	return nil, errors.New("sqlite item lookup is not supported yet")
}

func (d SQLiteDriver) CreateItem(_ *sql.DB, _ string, _ map[string]any) error {
	return errors.New("sqlite item creation is not supported yet")
}

func (d SQLiteDriver) UpdateItem(_ *sql.DB, _ string, _ map[string]any, _ map[string]any) error {
	return errors.New("sqlite item updates are not supported yet")
}

func (d SQLiteDriver) DeleteItem(_ *sql.DB, _ string, _ map[string]any) error {
	return errors.New("sqlite item deletion is not supported yet")
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

func (d SQLiteDriver) CreateColumn(_ *sql.DB, _ string, _ models.Column) error {
	return errors.New("sqlite column creation is not supported yet")
}

func (d SQLiteDriver) UpdateColumn(_ *sql.DB, _, _ string, _ models.Column) error {
	return errors.New("sqlite column updates are not supported yet")
}

func (d SQLiteDriver) DeleteColumn(_ *sql.DB, _, _ string) error {
	return errors.New("sqlite column deletion is not supported yet")
}

func quoteSQLiteIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}
