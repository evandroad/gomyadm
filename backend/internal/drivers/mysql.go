package drivers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/evandroad/gomyadm/internal/models"
)

type MySQLDriver struct{}

func init() {
	Register("mysql", MySQLDriver{})
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

func (d MySQLDriver) DescribeTable(db *sql.DB, table string) (*models.TableSchema, error) {
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
		var col models.TableColumn

		var nullable string
		var defaultValue sql.NullString

		err := rows.Scan(&col.Name, &col.Type, &nullable, &col.Key, &defaultValue, &col.Extra)
		if err != nil {
			return nil, err
		}

		col.Nullable = nullable == "YES"
		if defaultValue.Valid {
			col.Default = defaultValue.String
		}

		schema.Columns = append(schema.Columns, col)
	}

	return schema, nil
}

func (d MySQLDriver) SelectTable(db *sql.DB, table string) (*models.TableData, error) {
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