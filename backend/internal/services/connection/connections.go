package connectionService

import (
	"encoding/json"
	"errors"
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
	store *ConnectionsStore
	once  sync.Once
)

var (
	ErrConnectionNotFound = errors.New("connection not found")
	ErrConnectionExists   = errors.New("connection already exists")
)

func getStore() *ConnectionsStore {
	once.Do(func() {
		store = &ConnectionsStore{
			filePath: filepath.Join("data", "connections.json"),
		}
	})

	return store
}

func Init() error { return getStore().init() }
func GetAll() []models.ConnectionConfig { return getStore().list() }
func Create(conn models.ConnectionConfig) error { return getStore().create(conn) }
func Update(id string, conn models.ConnectionConfig) error { return getStore().update(id, conn) }
func Delete(id string) error { return getStore().delete(id) }

func (s *ConnectionsStore) init() error {
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

func (s *ConnectionsStore) list() []models.ConnectionConfig {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.ConnectionConfig, len(s.connections))
	copy(result, s.connections)

	return result
}

func (s *ConnectionsStore) create(conn models.ConnectionConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, existing := range s.connections {
		if existing.ID == conn.ID {
			return ErrConnectionExists
		}
	}

	s.connections = append(s.connections, conn)

	return s.saveToDisk()
}

func (s *ConnectionsStore) update(id string, conn models.ConnectionConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, existing := range s.connections {
		if existing.ID == id {
			conn.ID = existing.ID

			s.connections[i] = conn

			return s.saveToDisk()
		}
	}

	return ErrConnectionNotFound
}

func (s *ConnectionsStore) delete(id string) error {
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

	return ErrConnectionNotFound
}

func (s *ConnectionsStore) saveToDisk() error {
	data, err := json.MarshalIndent(s.connections, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}