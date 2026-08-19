package bindings

import (
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

type TableBinding struct {
	service *services.TableService
}

func NewTable(s *services.TableService) *TableBinding {
	return &TableBinding{service: s}
}

func (a *TableBinding) GetAll() ([]string, error) {
	return a.service.GetAll()
}

func (a *TableBinding) Create(table models.Table) error {
	return a.service.Create(table)
}