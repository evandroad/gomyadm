package drivers

import "github.com/evandroad/gomyadm/internal/models"

type Driver interface {
	BuildDSN(cfg models.ConnectionConfig) string
}