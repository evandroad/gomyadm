package bindings

import (
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

type ColumnBinding struct {
	service *services.ColumnService
}

func NewColumn(s *services.ColumnService) *ColumnBinding {
	return &ColumnBinding{service: s}
}

func (a *ColumnBinding) GetAll(table string) (*models.Table, error) {
	return a.service.GetAll(table)
}

func (a *ColumnBinding) Create(column models.ColumnRequest) error {
	return a.service.Create(column)
}

func (a *ColumnBinding) Update(column models.ColumnRequest) error {
	return a.service.Update(column)
}

func (a *ColumnBinding) Delete(table, column string) error {
	return a.service.Delete(table, column)
}