package bindings

import (
	"gomyadm/internal/models"
	"gomyadm/internal/services"
)

type SessionBinding struct {
	service *services.SessionService
}

func NewSession(s *services.SessionService) *SessionBinding {
	return &SessionBinding{service: s}
}

func (a *SessionBinding) Active() (models.ConnectionResponse, error) {
	return a.service.Active()
}