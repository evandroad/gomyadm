package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/go-chi/chi/v5"
)

type ConnectionHandler struct {
	Connections *db.ConnectionManager
}

func (h *ConnectionHandler) Connect(w http.ResponseWriter, r *http.Request) {
	var cfg db.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Connections.Connect(cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"success":true}`))
}

func (h *ConnectionHandler) Disconnect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.Connections.Disconnect(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
	})
}