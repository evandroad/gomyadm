package tableService

import (
	"fmt"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
)

func GetAll(m *db.ConnectionManager) ([]string, error) {
	driver, conn, err := m.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return []string{}, fmt.Errorf("Failed to get driver and connection")
	}

	tables, err := driver.GetAllTable(conn.DB)
	if err != nil {
		logger.Error("Failed to list tables: %v", err)
		return []string{}, fmt.Errorf("Failed to list tables")
	}

	return tables, nil
}