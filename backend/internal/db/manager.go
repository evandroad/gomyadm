package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectionManager struct {
	mu    sync.RWMutex
	pools map[string]*sql.DB
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		pools: make(map[string]*sql.DB),
	}
}

func (m *ConnectionManager) Connect(cfg ConnectionConfig) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.pools[cfg.ID]; exists {
    return fmt.Errorf("connection already exists")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return err
	}

	m.pools[cfg.ID] = db

	return nil
}

func (m *ConnectionManager) Disconnect(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	db, ok := m.pools[id]
	if !ok {
		return fmt.Errorf("connection not found")
	}

	err := db.Close()
	if err != nil {
		return err
	}

	delete(m.pools, id)

	return nil
}

func (m *ConnectionManager) Get(id string) (*sql.DB, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	db, ok := m.pools[id]
	if !ok {
		return nil, fmt.Errorf("connection not found")
	}

	return db, nil
}