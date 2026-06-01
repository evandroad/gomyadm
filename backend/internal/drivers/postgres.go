package drivers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/evandroad/gomyadm/internal/models"
)

type PostgresDriver struct{}

func init() {
	Register("postgres", PostgresDriver{})
}

func (d PostgresDriver) BuildDSN(cfg models.ConnectionConfig) string {
	if cfg.Database == "" {
		cfg.Database = "postgres"
	}
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
}

func (d PostgresDriver) ListTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`
		SELECT tablename
		FROM pg_tables
		WHERE schemaname = 'public'
	`)
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

func (d PostgresDriver) SelectTable(db *sql.DB, table string) (*models.TableData, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s"`, table)
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

func (d PostgresDriver) DescribeTable(db *sql.DB, table string) (*models.TableSchema, error) {
	query := `
		SELECT
    c.column_name,
    c.data_type,
    c.is_nullable,
    c.column_default,
    CASE
        WHEN tc.constraint_type = 'PRIMARY KEY' THEN 'PRI'
        WHEN tc.constraint_type = 'UNIQUE' THEN 'UNI'
        ELSE ''
    END AS column_key
		FROM information_schema.columns c
		LEFT JOIN information_schema.key_column_usage kcu
				ON c.table_name = kcu.table_name
			AND c.column_name = kcu.column_name
		LEFT JOIN information_schema.table_constraints tc
				ON kcu.constraint_name = tc.constraint_name
		WHERE c.table_schema = 'public'
			AND c.table_name = $1
		ORDER BY c.ordinal_position;
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

		err := rows.Scan(&col.Name, &col.Type, &nullable, &col.Key, &defaultValue)
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

func (d PostgresDriver) ListDatabases(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`
		SELECT datname
		FROM pg_database
		WHERE datistemplate = false
	`)
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