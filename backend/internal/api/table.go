package api

import (
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	. "github.com/evandroad/gomyadm/internal/respond"
)

type SchemaHandler struct {
	Connection *db.ConnectionManager
}

func (h *SchemaHandler) ListTables(w http.ResponseWriter, r *http.Request) {
	driver, conn, err := getDriverAndConnection(h.Connection)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	tables, err := driver.ListTables(conn.DB)
	if err != nil {
		logger.Error("Failed to list tables: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to list tables", nil)
		return
	}

	JSON(w, http.StatusOK, tables)
}