package drivers

import (
	"database/sql"

	"github.com/evandroad/gomyadm/internal/models"
)

type Driver interface {
	DriverName() string
	BuildDSN(cfg models.ConnectionConfig) string

	GetAllDatabase(db *sql.DB) ([]string, error)
	CreateDatabase(db *sql.DB, name string) error
	UpdateDatabase(db *sql.DB, oldName, newName string) error
	DeleteDatabase(db *sql.DB, name string) error

	GetAllTable(db *sql.DB) ([]string, error)
	CreateTable(db *sql.DB, table models.Table) error
	UpdateTable(db *sql.DB, oldName, newName string) error
	DeleteTable(db *sql.DB, table string) error

	GetAllItem(db *sql.DB, table string) (*models.TableData, error)
	CreateItem(db *sql.DB, table string, data map[string]any) error
	UpdateItem(db *sql.DB, table string, key map[string]any, data map[string]any) error
	DeleteItem(db *sql.DB, table string, key map[string]any) error

	GetAllColumn(db *sql.DB, table string) (*models.Table, error)
	CreateColumn(db *sql.DB, table string, column models.Column) error
	UpdateColumn(db *sql.DB, table, oldName string, column models.Column) error
	DeleteColumn(db *sql.DB, table string, column string) error
}