package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/evandroad/gomyadm/internal/drivers"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/logger"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type ConnectionManager struct {
	mu         sync.RWMutex
	connection *models.Connection
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connection: nil,
	}
}

func (m *ConnectionManager) Connect(cfg models.ConnectionConfig) (models.ConnectionResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.connection != nil {
		if m.connection.Matches(cfg) {
			return m.getConnection(), nil
		}

		if err := m.connection.DB.Close(); err != nil {
			return models.ConnectionResponse{}, err
		}

		m.connection = nil
	}

	err := m.createConnection(cfg)
	if err != nil {
		return models.ConnectionResponse{}, err
	}

	return m.getConnection(), nil
}

func (m *ConnectionManager) Disconnect() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.connection == nil {
		return fmt.Errorf("no active connection")
	}

	err := m.connection.DB.Close()
	if err != nil {
		logger.Error("Failed to close database connection: %v", err)
		return err
	}

	m.connection = nil
	return nil
}

func (m *ConnectionManager) Get() (*models.Connection, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.connection == nil {
		return nil, fmt.Errorf("no active connection")
	}

	return m.connection, nil
}

func (m *ConnectionManager) DB() *sql.DB {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.connection.DB
}

func (m *ConnectionManager) Active() (models.ConnectionResponse, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.connection == nil {
		return models.ConnectionResponse{}, fmt.Errorf("no active connection")
	}

	return m.getConnection(), nil
}

func (m *ConnectionManager) GetDatabase() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.connection == nil {
		return ""
	}

	return m.connection.Config.Database
}

func (m *ConnectionManager) SelectDatabase(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.connection == nil {
		return fmt.Errorf("no active connection")
	}

	err := m.connection.DB.Close()
	if err != nil {
		logger.Error("Failed to close database connection: %v", err)
		return err
	}

	m.connection.Config.Database = name

	err = m.createConnection(m.connection.Config)
	if err != nil {
		logger.Error("Failed to connect to selected database: %v", err)
		return err
	}

	return nil
}

func (m *ConnectionManager) GetDriverAndConnection() (drivers.Driver, *models.Connection, error) {
	conn, err := m.Get()
	if err != nil {
		logger.Error("Failed to get connection: %v", err)
		return nil, nil, err
	}

	driver, ok := drivers.GetDriver(conn.Config.Driver)
	if !ok {
		return nil, nil, fmt.Errorf("unsupported driver: %s", conn.Config.Driver)
	}

	return driver, conn, nil
}

func (m *ConnectionManager) createConnection(cfg models.ConnectionConfig) error {
	driver, ok := drivers.GetDriver(cfg.Driver)
	if !ok {
		return fmt.Errorf("unsupported driver")
	}

	dsn := driver.BuildDSN(cfg)

	db, err := sql.Open(driver.DriverName(), dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	databases, err := driver.GetAllDatabase(db)
	if err != nil {
		return fmt.Errorf("failed to list databases: %w", err)
	}

	m.connection = &models.Connection{
		Config: cfg,
		DB:     db,
		DBs:    databases,
	}

	return nil
}

func (m *ConnectionManager) getConnection() models.ConnectionResponse {
	return models.ConnectionResponse{
		ID:        m.connection.Config.ID,
		Name:      m.connection.Config.Name,
		Driver:    m.connection.Config.Driver,
		Host:      m.connection.Config.Host,
		Port:      m.connection.Config.Port,
		Database:  m.connection.Config.Database,
		Databases: m.connection.DBs,
	}
}