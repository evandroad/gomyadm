package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
)

type ItemHandler struct {
	Connection *db.ConnectionManager
}

func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	driver, conn, err := getDriverAndConnection(h.Connection)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	rows, err := driver.GetAllItem(conn.DB, table)
	if err != nil {
		logger.Error("Failed to select table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to select table", nil)
		return
	}

	JSON(w, http.StatusOK, rows)
}

func (h *ItemHandler) Insert(w http.ResponseWriter, r *http.Request) {
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

	err = driver.InsertItem(conn.DB, req.Table, req.Values)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to insert data", nil)
		return
	}

	Success(w, http.StatusOK, "Item inserido com sucesso.", nil)
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	err = driver.UpdateItem(conn.DB, req.Table, req.Key, req.Values)
	if err != nil {
		logger.Error("Failed to update data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to update data", nil)
		return
	}

	Success(w, http.StatusOK, "Item alterado com sucesso.", nil)
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = driver.DeleteItem(conn.DB, req.Table, req.Key)
	if err != nil {
		logger.Error("Failed to delete data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete data", nil)
		return
	}

	Success(w, http.StatusOK, "Item removido com sucesso.", nil)
}