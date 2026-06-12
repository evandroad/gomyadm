package drivers

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/evandroad/gomyadm/internal/models"
)

type MySQLDriver struct{}

func init() {
	Register("mysql", MySQLDriver{})
}

func (d MySQLDriver) DriverName() string {
	return "mysql"
}

func (d MySQLDriver) BuildDSN(cfg models.ConnectionConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

func (d MySQLDriver) ListTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string

	for rows.Next() {
		var table string

		if err := rows.Scan(&table); err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}

func (d MySQLDriver) ListDatabases(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`SHOW DATABASES`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		databases = append(databases, name)
	}

	return databases, nil
}

func (d MySQLDriver) GetAllItem(db *sql.DB, table string) (*models.TableData, error) {
	query := fmt.Sprintf("SELECT * FROM %s LIMIT 100", table)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]any, 0)

	for rows.Next() {
		values := make([]any, len(columns))
		valuePtrs := make([]any, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]any)

		for i, col := range columns {
			val := values[i]

			switch v := val.(type) {
				case []byte:
					rowMap[col] = string(v)
				case time.Time:
					rowMap[col] = v.Format("2006-01-02 15:04:05")
				default:
					rowMap[col] = v
			}
		}

		results = append(results, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &models.TableData{
		Columns: columns,
		Rows:    results,
	}, nil
}

func (d MySQLDriver) InsertItem(db *sql.DB, table string, data map[string]any) error {
	columns := make([]string, 0, len(data))
	placeholders := make([]string, 0, len(data))
	values := make([]any, 0, len(data))

	for col, val := range data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := db.Exec(query, values...)
	return err
}

func (d MySQLDriver) UpdateItem(db *sql.DB, table string, key map[string]any, data map[string]any) error {
	setClauses := make([]string, 0, len(data))
	whereClauses := make([]string, 0, len(key))
	values := make([]any, 0, len(data)+len(key))

	for col, val := range data {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	for col, val := range key {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		table,
		strings.Join(setClauses, ", "),
		strings.Join(whereClauses, " AND "),
	)

	_, err := db.Exec(query, values...)
	return err
}

func (d MySQLDriver) DeleteItem(db *sql.DB, table string, key map[string]any) error {
	whereClauses := make([]string, 0)
	values := make([]any, 0, len(key))

	for col, val := range key {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s",
		table,
		strings.Join(whereClauses, " AND "),
	)

	_, err := db.Exec(query, values...)
	return err
}

func (d MySQLDriver) GetAllColumn(db *sql.DB, table string) (*models.TableSchema, error) {
	query := `
		SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_KEY, COLUMN_DEFAULT, EXTRA
		FROM information_schema.columns
		WHERE table_schema = DATABASE()
		AND table_name = ?
		ORDER BY ORDINAL_POSITION
	`

	rows, err := db.Query(query, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schema := &models.TableSchema{
		Name: table,
	}

	for rows.Next() {
		var col models.ColumnDefinition
		var nullable string
		var key string
		var extra string
		var columnType string
		var defaultValue sql.NullString

		err := rows.Scan(&col.Name, &columnType, &nullable, &key, &defaultValue, &extra)
		if err != nil {
			return nil, err
		}

		col.Type = extractBaseType(columnType)
		col.Length = extractLength(columnType)
		col.Nullable = nullable == "YES"
		col.Primary = key == "PRI"
		col.Unique = key == "UNI"
		col.AutoIncrement = strings.Contains(extra, "auto_increment")
		if defaultValue.Valid {
			col.DefaultValue = defaultValue.String
		}

		schema.Columns = append(schema.Columns, col)
	}

	return schema, nil
}

func extractBaseType(columnType string) string {
	if idx := strings.Index(columnType, "("); idx != -1 {
		return strings.ToUpper(columnType[:idx])
	}

	return strings.ToUpper(columnType)
}

func extractLength(columnType string) *int {
	re := regexp.MustCompile(`\((\d+)`)
	match := re.FindStringSubmatch(columnType)

	if len(match) < 2 {
		return nil
	}

	n, err := strconv.Atoi(match[1])
	if err != nil {
		return nil
	}

	return &n
}

func (d MySQLDriver) InsertColumn(db *sql.DB, table string, column models.ColumnDefinition) error {
	colType := strings.ToUpper(column.Type)

	if column.Length != nil {
		colType = fmt.Sprintf("%s(%d)", colType, *column.Length)
	}

	query := fmt.Sprintf(
		"ALTER TABLE `%s` ADD COLUMN `%s` %s",
		table,
		column.Name,
		colType,
	)

	if !column.Nullable {
		query += " NOT NULL"
	}

	if column.AutoIncrement {
		query += " AUTO_INCREMENT"
	}

	if column.Unique {
		query += " UNIQUE"
	}

	if column.DefaultValue != "" {
		query += fmt.Sprintf(" DEFAULT '%s'", column.DefaultValue)
	}

	if column.Primary {
		query += ", ADD PRIMARY KEY (`" + column.Name + "`)"
	}

	_, err := db.Exec(query)
	return err
}

func (d MySQLDriver) UpdateColumn(db *sql.DB, table string, oldName string, column models.ColumnDefinition) error {
	colType := strings.ToUpper(column.Type)

	if column.Length != nil {
		colType = fmt.Sprintf("%s(%d)", colType, *column.Length)
	}

	query := fmt.Sprintf(
		"ALTER TABLE `%s` CHANGE COLUMN `%s` `%s` %s",
		table,
		oldName,
		column.Name,
		colType,
	)

	if !column.Nullable {
		query += " NOT NULL"
	}

	if column.AutoIncrement {
		query += " AUTO_INCREMENT"
	}

	if column.Unique {
		query += " UNIQUE"
	}

	if column.DefaultValue != "" {
		query += fmt.Sprintf(" DEFAULT '%s'", column.DefaultValue)
	}

	if column.Primary {
		query += ", ADD PRIMARY KEY (`" + column.Name + "`)"
	}

	_, err := db.Exec(query)
	return err
}