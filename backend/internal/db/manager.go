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