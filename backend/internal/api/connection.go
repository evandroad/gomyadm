package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/drivers"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/evandroad/gomyadm/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/rs/xid"
)

type ConnectionHandler struct {
	Connection *db.ConnectionManager
}

func (ch *ConnectionHandler) Connect(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	if cfg.ID == "" {
		cfg.ID = xid.New().String()
	}

	var data models.ConnectionResponse
	data, err = ch.Connection.Connect(cfg)
	if err != nil {
		logger.Error("Failed to connect: %v", err)
		Error(w, http.StatusBadRequest, "Failed to connect", nil)
		return
	}

	JSON(w, http.StatusOK, data)
}

func (ch *ConnectionHandler) Disconnect(w http.ResponseWriter, r *http.Request) {
	err := ch.Connection.Disconnect()
	if err != nil {
		logger.Error("Failed to disconnect: %v", err)
		Error(w, http.StatusNotFound, "Connection not found", nil)
		return
	}

	Success(w, http.StatusOK, nil)
}

func (ch *ConnectionHandler) Active(w http.ResponseWriter, r *http.Request) {
	conn, err := ch.Connection.Active()
	if err != nil {
		logger.Error("Failed to get active connection: %v", err)
		Error(w, http.StatusNotFound, "No active connection", nil)
		return
	}

	JSON(w, http.StatusOK, conn)
}

func (ch *ConnectionHandler) SelectDatabase(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Database string `json:"database"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	err = ch.Connection.SelectDatabase(req.Database)
	if err != nil {
		logger.Error("Failed to select database: %v", err)
		Error(w, http.StatusBadRequest, "Failed to select database", nil)
		return
	}

	Success(w, http.StatusOK, H{ "message": "Database " + req.Database + " selected successfully" })
}

func (ch *ConnectionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	connections := storage.GetConnectionsStore().List()
	JSON(w, http.StatusOK, connections)
}

func (ch *ConnectionHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = storage.GetConnectionsStore().Create(cfg)
	if err != nil {
		logger.Error("Failed to save connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to save connection", nil)
		return
	}

	JSON(w, http.StatusCreated, map[string]any{
		"message": "connection saved",
	})
}

func (ch *ConnectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = storage.GetConnectionsStore().Update(cfg.ID, cfg)
	if err != nil {
		logger.Error("Failed to update connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to update connection", nil)
		return
	}

	JSON(w, http.StatusOK, map[string]any{"message": "connection updated"})
}

func (ch *ConnectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := storage.GetConnectionsStore().Delete(id)
	if err != nil {
		logger.Error("Failed to delete connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete connection", nil)
		return
	}

	Success(w, http.StatusOK, H{ "message": "connection deleted" })
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