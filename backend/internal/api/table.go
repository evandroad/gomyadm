package api

import (
	"log"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/drivers"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
)

type SchemaHandler struct {
	Connection *db.ConnectionManager
}

func (h *SchemaHandler) ListTables(w http.ResponseWriter, r *http.Request) {
	conn, err := h.Connection.Get()
	if err != nil {
		log.Printf("Failed to get active connection: %v", err)
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
		log.Printf("Failed to list tables: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to list tables", nil)
		return
	}

	JSON(w, http.StatusOK, tables)
}

func (h *SchemaHandler) SelectTable(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	conn, err := h.Connection.Get()
	if err != nil {
		log.Printf("Failed to get active connection: %v", err)
		Error(w, http.StatusNotFound, "Connection not found", nil)
		return
	}

	driver, ok := drivers.GetDriver(conn.Config.Driver)
	if !ok {
		Error(w, http.StatusBadRequest, "unsupported driver", nil)
		return
	}

	rows, err := driver.SelectTable(conn.DB, table)
	if err != nil {
		log.Printf("Failed to select table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to select table", nil)
		return
	}

	JSON(w, http.StatusOK, rows)
}

func (h *SchemaHandler) DescribeTable(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	conn, err := h.Connection.Get()
	if err != nil {
		log.Printf("Failed to get active connection: %v", err)
		Error(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	driver, ok := drivers.GetDriver(conn.Config.Driver)
	if !ok {
		Error(w, http.StatusBadRequest, "unsupported driver", nil)
		return
	}

	schema, err := driver.DescribeTable(conn.DB, table)
	if err != nil {
		log.Printf("Failed to describe table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to describe table", nil)
		return
	}

	JSON(w, http.StatusOK, schema)
}