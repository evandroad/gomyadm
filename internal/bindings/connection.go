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