package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/evandroad/gomyadm/internal/drivers"
	"github.com/evandroad/gomyadm/internal/models"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Connection struct {
	Config models.ConnectionConfig
	DB     *sql.DB
}

type ConnectionManager struct {
	mu         sync.RWMutex
	connection *Connection
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
		return models.ConnectionResponse{}, fmt.Errorf("connection already exists")
	}

	driver, ok := drivers.GetDriver(cfg.Driver)
	if !ok {
		return models.ConnectionResponse{}, fmt.Errorf("unsupported driver")
	}

	dsn := driver.BuildDSN(cfg)

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		log.Printf("Failed to open database connection: %v", err)
		return models.ConnectionResponse{}, err
	}

	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
		return models.ConnectionResponse{}, err
	}

	m.connection = &Connection{
		Config: cfg,
		DB:     db,
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
		log.Printf("Failed to close database connection: %v", err)
		return err
	}

	m.connection = nil
	return nil
}

func (m *ConnectionManager) Get() (*Connection, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.connection == nil {
		return nil, fmt.Errorf("no active connection")
	}

	return m.connection, nil
}

func (m *ConnectionManager) Active() (models.ConnectionResponse, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.connection == nil {
		return models.ConnectionResponse{}, fmt.Errorf("no active connection")
	}

	return m.getConnection(), nil
}

func (m *ConnectionManager) getConnection() models.ConnectionResponse {
	return models.ConnectionResponse{
		ID:       m.connection.Config.ID,
		Name:     m.connection.Config.Name,
		Driver:   m.connection.Config.Driver,
		Host:     m.connection.Config.Host,
		Port:     m.connection.Config.Port,
		Database: m.connection.Config.Database,
	}
}