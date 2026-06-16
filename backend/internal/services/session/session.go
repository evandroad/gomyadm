package sessionService

import (
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/rs/xid"
)

type SessionService struct {
	Connection *db.ConnectionManager
}

func NewSessionService(conn *db.ConnectionManager) *SessionService {
	return &SessionService{
		Connection: conn,
	}
}

func (s *SessionService) Connect(cfg models.ConnectionConfig) (models.ConnectionResponse, error) {
	if cfg.ID == "" {
		cfg.ID = xid.New().String()
	}
	
	return s.Connection.Connect(cfg)
}

func (s *SessionService) Disconnect() error {
	return s.Connection.Disconnect()
}

func (s *SessionService) Active() (models.ConnectionResponse, error) {
	return s.Connection.Active()
}