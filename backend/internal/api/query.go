package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/services"
	"github.com/evandroad/gomyadm/internal/logger"
	. "github.com/evandroad/gomyadm/internal/respond"
)

type QueryHandler struct {
	Connection *db.ConnectionManager
}

func (h *QueryHandler) ExecuteQuery(w http.ResponseWriter, r *http.Request) {
	var req models.QueryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode query request: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	conn, err := h.Connection.Get()
	if err != nil {
		logger.Error("Failed to get active connection: %v", err)
		Error(w, http.StatusNotFound, "Connection not found", nil)
		return
	}

	result, err := services.ExecuteQuery(conn.DB, req.Query)
	if err != nil {
		logger.Error("Failed to execute query: %v", err)
		Error(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	JSON(w, http.StatusOK, result)
}