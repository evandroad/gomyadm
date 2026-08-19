package bindings

import (
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

type ConnectionBinding struct {
	service *services.ConnectionsStore
}

func NewConnection(s *services.ConnectionsStore) *ConnectionBinding {
	return &ConnectionBinding{service: s}
}

func (a *ConnectionBinding) GetAll() []models.ConnectionConfig {
	return a.service.GetAll()
}

func (a *ConnectionBinding) Create(con models.ConnectionConfig) error {
	return a.service.Create(con)
}

func (a *ConnectionBinding) Update(id string, con models.ConnectionConfig) error {
	return a.service.Update(id, con)
}

func (a *ConnectionBinding) Delete(id string) error {
	return a.service.Delete(id)
}