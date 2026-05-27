package drivers

import (
	"fmt"

	"github.com/evandroad/gomyadm/internal/models"
)

type MySQLDriver struct{}

func init() {
	Register("mysql", MySQLDriver{})
}

func (d MySQLDriver) BuildDSN(cfg models.ConnectionConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}