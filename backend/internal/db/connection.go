package db

import (
	"database/sql"

	"github.com/evandroad/gomyadm/internal/models"
)

type Connection struct {
	Config models.ConnectionConfig
	DB     *sql.DB
	DBs		[]string
}

func (c *Connection) Matches(cfg models.ConnectionConfig) bool {
	return c.Config.Driver == cfg.Driver &&
		c.Config.Host == cfg.Host &&
		c.Config.Port == cfg.Port &&
		c.Config.Username == cfg.Username &&
		c.Config.Database == cfg.Database
}