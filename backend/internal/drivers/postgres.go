package drivers

import (
	"database/sql"
	"fmt"

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

func (d PostgresDriver) DescribeTable(db *sql.DB, table string) (*models.TableSchema, error) {
	query := `
		SELECT column_name, data_type, is_nullable
		FROM information_schema.columns
		WHERE table_name = $1
		ORDER BY ordinal_position
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

		err := rows.Scan(&col.Name, &col.Type, &nullable)
		if err != nil {
			return nil, err
		}

		col.Nullable = nullable == "YES"

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