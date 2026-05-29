package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/evandroad/gomyadm/internal/models"
)

type ConnectionsStore struct {
	mu          sync.RWMutex
	filePath    string
	connections []models.ConnectionConfig
}

var (
	instance *ConnectionsStore
	once     sync.Once
)

func GetConnectionsStore() *ConnectionsStore {
	once.Do(func() {
		instance = &ConnectionsStore{
			filePath: filepath.Join("data", "connections.json"),
		}
	})

	return instance
}

func (s *ConnectionsStore) Init() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// cria pasta data/
	err := os.MkdirAll(filepath.Dir(s.filePath), os.ModePerm)
	if err != nil {
		return err
	}

	// arquivo não existe ainda
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		s.connections = []models.ConnectionConfig{}

		return s.saveToDisk()
	}

	// lê arquivo
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	// arquivo vazio
	if len(data) == 0 {
		s.connections = []models.ConnectionConfig{}
		return nil
	}

	err = json.Unmarshal(data, &s.connections)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConnectionsStore) List() []models.ConnectionConfig {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.ConnectionConfig, len(s.connections))
	copy(result, s.connections)

	return result
}

func (s *ConnectionsStore) Get(id string) (*models.ConnectionConfig, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, conn := range s.connections {
		if conn.ID == id {
			c := conn
			return &c, nil
		}
	}

	return nil, fmt.Errorf("connection not found")
}

func (s *ConnectionsStore) Create(conn models.ConnectionConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, existing := range s.connections {
		if existing.ID == conn.ID {
			return fmt.Errorf("connection already exists")
		}
	}

	s.connections = append(s.connections, conn)

	return s.saveToDisk()
}

func (s *ConnectionsStore) Update(id string, conn models.ConnectionConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, existing := range s.connections {
		if existing.ID == id {
			conn.ID = existing.ID

			s.connections[i] = conn

			return s.saveToDisk()
		}
	}

	return fmt.Errorf("connection not found")
}

func (s *ConnectionsStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, conn := range s.connections {
		if conn.ID == id {
			s.connections = append(
				s.connections[:i],
				s.connections[i+1:]...,
			)

			return s.saveToDisk()
		}
	}

	return fmt.Errorf("connection not found")
}

func (s *ConnectionsStore) saveToDisk() error {
	data, err := json.MarshalIndent(s.connections, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}