package drivers

import (
	"database/sql"

	"github.com/evandroad/gomyadm/internal/models"
)

type Driver interface {
	DriverName() string
	BuildDSN(cfg models.ConnectionConfig) string
	ListTables(db *sql.DB) ([]string, error)
	ListDatabases(db *sql.DB) ([]string, error)
	
	GetAllItem(db *sql.DB, table string) (*models.TableData, error)
	InsertItem(db *sql.DB, table string, data map[string]any) error
	UpdateItem(db *sql.DB, table string, key map[string]any, data map[string]any) error
	DeleteItem(db *sql.DB, table string, key map[string]any) error
	
	GetAllColumn(db *sql.DB, table string) (*models.TableSchema, error)
	InsertColumn(db *sql.DB, table string, column models.ColumnDefinition) error
	UpdateColumn(db *sql.DB, table string, oldName string, column models.ColumnDefinition) error
}