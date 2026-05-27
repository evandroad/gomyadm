package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/go-chi/chi/v5"
)

type SchemaHandler struct {
	Connections *db.ConnectionManager
}

func (h *SchemaHandler) ListTables(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	db, err := h.Connections.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tables []string

	for rows.Next() {
		var table string

		err := rows.Scan(&table)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tables = append(tables, table)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}