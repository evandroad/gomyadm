package drivers

import (
	"fmt"

	"github.com/evandroad/gomyadm/internal/models"
)

type PostgresDriver struct{}

func init() {
	Register("postgres", PostgresDriver{})
}

func (d PostgresDriver) BuildDSN(cfg models.ConnectionConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
}