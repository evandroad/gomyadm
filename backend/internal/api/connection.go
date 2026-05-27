package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/models"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
	"github.com/rs/xid"
)

type ConnectionHandler struct {
	Connections *db.ConnectionManager
}

func (ch *ConnectionHandler) Connect(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	cfg.ID = xid.New().String()

	var data models.ConnectionResponse
	data, err = ch.Connections.Connect(cfg)
	if err != nil {
		Error(w, http.StatusBadRequest, "Failed to connect", nil)
		return
	}

	JSON(w, http.StatusOK, data)
}

func (ch *ConnectionHandler) Disconnect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := ch.Connections.Disconnect(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Connection not found", nil)
		return
	}

	Success(w, http.StatusOK, nil)
}

func (ch *ConnectionHandler) List(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, ch.Connections.List())
}