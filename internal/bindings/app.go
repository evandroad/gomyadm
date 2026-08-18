package bindings

import "gomyadm/internal/services"

type AppBinding struct {
	service *services.AppService
}

func NewApp(s *services.AppService) *AppBinding {
	return &AppBinding{service: s}
}

func (a *AppBinding) Version() string {
	return a.service.Version()
}