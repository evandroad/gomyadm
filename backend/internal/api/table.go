package api

import (
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/drivers"
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

	driver, ok := drivers.GetDriver(conn.Config.Driver)
	if !ok {
		Error(w, http.StatusBadRequest, "unsupported driver", nil)
		return
	}

	tables, err := driver.ListTables(conn.DB)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Failed to list tables", nil)
		return
	}

	JSON(w, http.StatusOK, tables)
}

func (h *SchemaHandler) DescribeTable(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	table := chi.URLParam(r, "table")

	conn, err := h.Connections.Get(id)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	driver, ok := drivers.GetDriver(conn.Config.Driver)
	if !ok {
		Error(w, http.StatusBadRequest, "unsupported driver", nil)
		return
	}

	schema, err := driver.DescribeTable(
		conn.DB,
		table,
	)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Failed to describe table", nil)
		return
	}

	JSON(w, http.StatusOK, schema)
}