package api

import (
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
)

type SchemaHandler struct {
	Connections *db.ConnectionManager
}

func (h *SchemaHandler) ListTables(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	conn, err := h.Connections.Get(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Connection not found", nil)
		return
	}

	rows, err := conn.DB.Query("SHOW TABLES")
	if err != nil {
		Error(w, http.StatusInternalServerError, "Failed to query tables", nil)
		return
	}
	defer rows.Close()

	var tables []string

	for rows.Next() {
		var table string

		err := rows.Scan(&table)
		if err != nil {
			Error(w, http.StatusInternalServerError, "Failed to scan table", nil)
			return
		}

		tables = append(tables, table)
	}

	JSON(w, http.StatusOK, tables)
}