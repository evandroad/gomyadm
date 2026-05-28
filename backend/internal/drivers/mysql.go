package drivers

import (
	"database/sql"
	"fmt"

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
		SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_KEY
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
		var key string

		err := rows.Scan(&col.Name, &col.Type, &nullable, &key)
		if err != nil {
			return nil, err
		}

		col.Nullable = nullable == "YES"
		col.PrimaryKey = key == "PRI"

		schema.Columns = append(schema.Columns, col)
	}

	return schema, nil
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