package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
)

type ColumnHandler struct {
	Connection *db.ConnectionManager
}

func (h *ColumnHandler) GetSchema(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	driver, conn, err := getDriverAndConnection(h.Connection)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	schema, err := driver.TableStructure(conn.DB, table)
	if err != nil {
		logger.Error("Failed to describe table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to describe table", nil)
		return
	}

	JSON(w, http.StatusOK, schema)
}

func (h *ColumnHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Table  string         			   `json:"table"`
		Values models.ColumnDefinition `json:"values"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	driver, conn, err := getDriverAndConnection(h.Connection)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	err = driver.InsertColumn(conn.DB, req.Table, req.Values)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to insert data", nil)
		return
	}

	Success(w, http.StatusOK, nil)
}