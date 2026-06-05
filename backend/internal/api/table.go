package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/drivers"
	"github.com/evandroad/gomyadm/internal/logger"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
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

func (h *SchemaHandler) SelectTable(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	driver, conn, err := getDriverAndConnection(h.Connection)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	rows, err := driver.SelectTable(conn.DB, table)
	if err != nil {
		logger.Error("Failed to select table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to select table", nil)
		return
	}

	JSON(w, http.StatusOK, rows)
}

func (h *SchemaHandler) TableStructure(w http.ResponseWriter, r *http.Request) {
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

func (h *SchemaHandler) InsertValue(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Table  string         `json:"table"`
		Values map[string]any `json:"values"`
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

	err = driver.InsertValue(conn.DB, req.Table, req.Values)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to insert data", nil)
		return
	}

	Success(w, http.StatusOK, nil)
}

func (h *SchemaHandler) UpdateValue(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Table  string         `json:"table"`
		Key    map[string]any `json:"key"`
		Values map[string]any `json:"values"`
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

	err = driver.UpdateValue(conn.DB, req.Table, req.Key, req.Values)
	if err != nil {
		logger.Error("Failed to update data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to update data", nil)
		return
	}

	Success(w, http.StatusOK, nil)
}

func (h *SchemaHandler) DeleteValue(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Table  string         `json:"table"`
		Key    map[string]any `json:"key"`
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

	err = driver.DeleteValue(conn.DB, req.Table, req.Key)
	if err != nil {
		logger.Error("Failed to delete data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete data", nil)
		return
	}

	Success(w, http.StatusOK, nil)
}

func getDriverAndConnection(cm *db.ConnectionManager) (drivers.Driver, *db.Connection, error) {
	conn, err := cm.Get()
	if err != nil {
		logger.Error("Failed to get active connection: %v", err)
		return nil, nil, err
	}

	driver, ok := drivers.GetDriver(conn.Config.Driver)
	if !ok {
		return nil, nil, fmt.Errorf("unsupported driver: %s", conn.Config.Driver)
	}

	return driver, conn, nil
}