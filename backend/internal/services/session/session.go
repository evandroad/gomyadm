package sessionService

import (
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/models"
)

func Connect(m *db.ConnectionManager, cfg models.ConnectionConfig) (models.ConnectionResponse, error) {
	return m.Connect(cfg)
}

func Disconnect(m *db.ConnectionManager) error {
	return m.Disconnect()
}

func Active(m *db.ConnectionManager) (models.ConnectionResponse, error) {
	return m.Active()
}