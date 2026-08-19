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

func (a *DatabaseBinding) Select(database string) error {
	return a.service.Select(database)
}

func (a *DatabaseBinding) Create(database models.DatabaseRequest) error {
	return a.service.Create(database)
}

func (a *DatabaseBinding) Update(database models.DatabaseRequest) error {
	return a.service.Update(database)
}

func (a *DatabaseBinding) Delete(name string) error {
	return a.service.Delete(name)
}