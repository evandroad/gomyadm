package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
)

type Handler struct {
	Manager *db.ConnectionManager
}

func (h *Handler) TestConnection(w http.ResponseWriter, r *http.Request) {
	var cfg db.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Manager.Connect(cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"success":true}`))
}