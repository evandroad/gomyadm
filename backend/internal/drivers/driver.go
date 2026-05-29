package drivers

import (
	"database/sql"

	"github.com/evandroad/gomyadm/internal/models"
)

type Driver interface {
	BuildDSN(cfg models.ConnectionConfig) string
	ListTables(db *sql.DB) ([]string, error)
	DescribeTable(db *sql.DB, table string) (*models.TableSchema, error)
	SelectTable(db *sql.DB, table string) (*models.TableData, error)
	ListDatabases(db *sql.DB) ([]string, error)
}