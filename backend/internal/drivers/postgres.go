package drivers

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/evandroad/gomyadm/internal/models"
)

type PostgresDriver struct{}

func init() {
	Register("postgres", PostgresDriver{})
}

func (d PostgresDriver) DriverName() string {
	return "pgx"
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

func (d PostgresDriver) GetAllItem(db *sql.DB, table string) (*models.TableData, error) {
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

func (d PostgresDriver) InsertItem(db *sql.DB, table string, data map[string]any) error {
	columns := make([]string, 0, len(data))
	values := make([]any, 0, len(data))
	placeholders := make([]string, 0, len(data))

	i := 1
	for col, val := range data {
		columns = append(columns, fmt.Sprintf(`"%s"`, col))
		values = append(values, val)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		i++
	}

	query := fmt.Sprintf(
		`INSERT INTO "%s" (%s) VALUES (%s)`,
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := db.Exec(query, values...)
	return err
}

func (d PostgresDriver) UpdateItem(db *sql.DB, table string, key map[string]any, data map[string]any) error {
	setClauses := make([]string, 0, len(data))
	whereClauses := make([]string, 0, len(key))
	values := make([]any, 0, len(data)+len(key))

	i := 1
	for col, val := range data {
		setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, col, i))
		values = append(values, val)
		i++
	}

	for col, val := range key {
		whereClauses = append(whereClauses, fmt.Sprintf(`"%s" = $%d`, col, i))
		values = append(values, val)
		i++
	}

	query := fmt.Sprintf(
		`UPDATE "%s" SET %s WHERE %s`,
		table,
		strings.Join(setClauses, ", "),
		strings.Join(whereClauses, " AND "),
	)

	_, err := db.Exec(query, values...)
	return err
}

func (d PostgresDriver) DeleteItem(db *sql.DB, table string, key map[string]any) error {
	whereClauses := make([]string, 0, len(key))
	values := make([]any, 0, len(key))

	for col, val := range key {
		whereClauses = append(whereClauses, fmt.Sprintf(`"%s" = $%d`, col, len(values)+1))
		values = append(values, val)
	}

	query := fmt.Sprintf(
		`DELETE FROM "%s" WHERE %s`,
		table,
		strings.Join(whereClauses, " AND "),
	)

	_, err := db.Exec(query, values...)
	return err
}

func (d PostgresDriver) GetAllColumn(db *sql.DB, table string) (*models.TableSchema, error) {
	query := `
		SELECT
			c.column_name,
			c.data_type,
			c.character_maximum_length,
			c.is_nullable,
			c.column_default,
			c.is_identity,
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
		var col models.ColumnDefinition
		var nullable string
		var key string
		var isIdentity string
		var length sql.NullInt64
		var defaultValue sql.NullString

		err := rows.Scan(&col.Name, &col.Type, &length, &nullable, &defaultValue, &isIdentity, &key)
		if err != nil {
			return nil, err
		}

		if length.Valid {
			l := int(length.Int64)
			col.Length = &l
		}

		col.Nullable = nullable == "YES"
		col.Primary = key == "PRI"
		col.Unique = key == "UNI"

		if defaultValue.Valid {
			col.DefaultValue = defaultValue.String
		}

		col.AutoIncrement = isIdentity == "YES" || (defaultValue.Valid && strings.Contains(defaultValue.String, "nextval("))

		schema.Columns = append(schema.Columns, col)
	}

	return schema, nil
}

func (d PostgresDriver) InsertColumn(db *sql.DB, table string, column models.ColumnDefinition) error {
	colType := strings.ToUpper(column.Type)

	if column.AutoIncrement {
		colType = "SERIAL"
	} else if column.Length != nil {
		colType = fmt.Sprintf("%s(%d)", colType, *column.Length)
	}

	query := fmt.Sprintf(
		`ALTER TABLE "%s" ADD COLUMN "%s" %s`,
		table,
		column.Name,
		colType,
	)

	if !column.Nullable {
		query += " NOT NULL"
	}

	if column.DefaultValue != "" {
		query += fmt.Sprintf(" DEFAULT '%s'", column.DefaultValue)
	}

	if _, err := db.Exec(query); err != nil {
		return err
	}

	if column.Primary {
		_, err := db.Exec(fmt.Sprintf(
			`ALTER TABLE "%s" ADD PRIMARY KEY ("%s")`,
			table,
			column.Name,
		))
		if err != nil {
			return err
		}
	}

	if column.Unique {
		_, err := db.Exec(fmt.Sprintf(
			`ALTER TABLE "%s" ADD UNIQUE ("%s")`,
			table,
			column.Name,
		))
		if err != nil {
			return err
		}
	}

	return nil
}

func (d PostgresDriver) UpdateColumn(db *sql.DB, table string, oldName string, column models.ColumnDefinition) error {
	colType := strings.ToUpper(column.Type)

	if column.AutoIncrement {
		colType = "SERIAL"
	} else if column.Length != nil {
		colType = fmt.Sprintf("%s(%d)", colType, *column.Length)
	}

	query := fmt.Sprintf(
		`ALTER TABLE "%s" ADD COLUMN "%s" %s`,
		table,
		column.Name,
		colType,
	)

	if !column.Nullable {
		query += " NOT NULL"
	}

	if column.DefaultValue != "" {
		query += fmt.Sprintf(" DEFAULT '%s'", column.DefaultValue)
	}

	if _, err := db.Exec(query); err != nil {
		return err
	}

	if column.Primary {
		_, err := db.Exec(fmt.Sprintf(
			`ALTER TABLE "%s" ADD PRIMARY KEY ("%s")`,
			table,
			column.Name,
		))
		if err != nil {
			return err
		}
	}

	if column.Unique {
		_, err := db.Exec(fmt.Sprintf(
			`ALTER TABLE "%s" ADD UNIQUE ("%s")`,
			table,
			column.Name,
		))
		if err != nil {
			return err
		}
	}

	return nil
}