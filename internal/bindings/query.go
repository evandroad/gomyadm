package bindings

import (
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

type QueryBinding struct {
	service *services.QueryService
}

func NewQuery(s *services.QueryService) *QueryBinding {
	return &QueryBinding{service: s}
}

func (a *QueryBinding) Query(query string) (*models.QueryResult, error) {
	return a.service.ExecuteQuery(query)
}