package drivers

import (
	"database/sql"

	"github.com/evandroad/gomyadm/internal/models"
)

type Driver interface {
	DriverName() string
	BuildDSN(cfg models.ConnectionConfig) string
	GetAllDatabase(db *sql.DB) ([]string, error)
	
	GetAllTable(db *sql.DB) ([]string, error)
	CreateTable(db *sql.DB, table string, columns []models.ColumnDefinition) error
	
	GetAllItem(db *sql.DB, table string) (*models.TableData, error)
	CreateItem(db *sql.DB, table string, data map[string]any) error
	UpdateItem(db *sql.DB, table string, key map[string]any, data map[string]any) error
	DeleteItem(db *sql.DB, table string, key map[string]any) error
	
	GetAllColumn(db *sql.DB, table string) (*models.TableSchema, error)
	CreateColumn(db *sql.DB, table string, column models.ColumnDefinition) error
	UpdateColumn(db *sql.DB, table string, oldName string, column models.ColumnDefinition) error
	DeleteColumn(db *sql.DB, table string, column string) error
}