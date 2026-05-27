package db

import (
	"database/sql"
	"fmt"
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
	mu    sync.RWMutex
	pools map[string]*Connection
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		pools: make(map[string]*Connection),
	}
}

func (m *ConnectionManager) Connect(cfg models.ConnectionConfig) (models.ConnectionResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.pools[cfg.ID]; exists {
    return models.ConnectionResponse{}, fmt.Errorf("connection already exists")
	}

	driver, ok := drivers.GetDriver(cfg.Driver)
	if !ok {
		return models.ConnectionResponse{}, fmt.Errorf("unsupported driver")
	}

	dsn := driver.BuildDSN(cfg)

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return models.ConnectionResponse{}, err
	}

	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return models.ConnectionResponse{}, err
	}

	m.pools[cfg.ID] = &Connection{
		Config: cfg,
		DB:     db,
	}

	return models.ConnectionResponse{
		ID:       cfg.ID,
		Name:     cfg.Name,
		Driver:   cfg.Driver,
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Database,
	}, nil
}

func (m *ConnectionManager) Disconnect(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	conn, ok := m.pools[id]
	if !ok {
		return fmt.Errorf("connection not found")
	}

	err := conn.DB.Close()
	if err != nil {
		return err
	}

	delete(m.pools, id)

	return nil
}

func (m *ConnectionManager) Get(id string) (*Connection, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	conn, ok := m.pools[id]
	if !ok {
		return nil, fmt.Errorf("connection not found")
	}

	return conn, nil
}

func (m *ConnectionManager) List() []models.ConnectionResponse {
	m.mu.RLock()
	defer m.mu.RUnlock()

	connections:= []models.ConnectionResponse{}

	for _, conn := range m.pools {
		connections = append(connections, models.ConnectionResponse{
			ID:       conn.Config.ID,
			Name:     conn.Config.Name,
			Driver:   conn.Config.Driver,
			Host:     conn.Config.Host,
			Port:     conn.Config.Port,
			Database: conn.Config.Database,
		})
	}

	return connections
}