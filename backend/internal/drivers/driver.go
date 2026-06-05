package drivers

import (
	"database/sql"

	"github.com/evandroad/gomyadm/internal/models"
)

type Driver interface {
	DriverName() string
	BuildDSN(cfg models.ConnectionConfig) string
	ListTables(db *sql.DB) ([]string, error)
	TableStructure(db *sql.DB, table string) (*models.TableSchema, error)
	SelectTable(db *sql.DB, table string) (*models.TableData, error)
	ListDatabases(db *sql.DB) ([]string, error)
	InsertValue(db *sql.DB, table string, data map[string]any) error
	UpdateValue(db *sql.DB, table string, key map[string]any, data map[string]any) error
	DeleteValue(db *sql.DB, table string, key map[string]any) error
}