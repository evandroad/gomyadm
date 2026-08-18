package bindings

import (
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

type DatabaseBinding struct {
	service *services.DatabaseService
}

func NewDatabase(s *services.DatabaseService) *DatabaseBinding {
	return &DatabaseBinding{service: s}
}

func (a *DatabaseBinding) GetAll() (models.DatabaseResponse, error) {
	return a.service.GetAll()
}