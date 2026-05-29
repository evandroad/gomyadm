package api

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/storage"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/rs/xid"
)

type ConnectionHandler struct {
	Connection *db.ConnectionManager
}

func (ch *ConnectionHandler) Connect(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	cfg.ID = xid.New().String()

	var data models.ConnectionResponse
	data, err = ch.Connection.Connect(cfg)
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		Error(w, http.StatusBadRequest, "Failed to connect", nil)
		return
	}

	JSON(w, http.StatusOK, data)
}

func (ch *ConnectionHandler) Disconnect(w http.ResponseWriter, r *http.Request) {
	err := ch.Connection.Disconnect()
	if err != nil {
		log.Printf("Failed to disconnect: %v", err)
		Error(w, http.StatusNotFound, "Connection not found", nil)
		return
	}

	Success(w, http.StatusOK, nil)
}

func (ch *ConnectionHandler) Active(w http.ResponseWriter, r *http.Request) {
	conn, err := ch.Connection.Active()
	if err != nil {
		log.Printf("Failed to get active connection: %v", err)
		Error(w, http.StatusNotFound, "Connection not found", nil)
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
		log.Printf("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	err = ch.Connection.SelectDatabase(req.Database)
	if err != nil {
		log.Printf("Failed to select database: %v", err)
		Error(w, http.StatusBadRequest, "Failed to select database", nil)
		return
	}

	Success(w, http.StatusOK, H{ "message": "Database " + req.Database + " selected successfully" })
}

func (ch *ConnectionHandler) ListConnections(w http.ResponseWriter, r *http.Request) {
	connections := storage.GetConnectionsStore().List()
	JSON(w, http.StatusOK, connections)
}

func (ch *ConnectionHandler) SaveConnection(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)

		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = storage.GetConnectionsStore().Create(cfg)
	if err != nil {
		log.Printf("Failed to save connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to save connection", nil)
		return
	}

	JSON(w, http.StatusCreated, map[string]any{
		"message": "connection saved",
	})
}