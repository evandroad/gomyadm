package context

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type Context struct {
	ConnectionID string `json:"connection,omitempty"`
	Database     string `json:"database,omitempty"`
}

type Store struct {
	mu      sync.RWMutex
	path    string
	context Context
}

func NewStore() (*Store, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(configDir, "gomyadm")

	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, err
	}

	store := &Store{
		path: filepath.Join(dir, "context.json"),
	}

	if err := store.load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) load() error {
	data, err := os.ReadFile(s.path)

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, &s.context)
}

func (s *Store) save() error {
	data, err := json.MarshalIndent(s.context, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(
		s.path,
		data,
		0600,
	)
}

func (s *Store) Get() Context {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.context
}

func (s *Store) SetConnection(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.context.ConnectionID = id

	// Ao trocar conexão, database e table deixam de ser válidos.
	s.context.Database = ""

	return s.save()
}

func (s *Store) SetDatabase(database string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.context.Database = database

	return s.save()
}

func (s *Store) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.context = Context{}

	return s.save()
}