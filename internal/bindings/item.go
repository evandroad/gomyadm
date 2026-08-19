package bindings

import (
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

type ItemBinding struct {
	service *services.ItemService
}

func NewItem(s *services.ItemService) *ItemBinding {
	return &ItemBinding{service: s}
}

func (a *ItemBinding) GetAll(table string) (*models.TableData, error) {
	return a.service.GetAll(table)
}