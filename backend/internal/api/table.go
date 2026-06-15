package api

import (
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/evandroad/gomyadm/internal/services/table"
)

type SchemaHandler struct {
	Connection *db.ConnectionManager
}

// @Summary Lista de tabelas
// @Description Retorna todas as tabelas do banco selecionado
// @Tags tables
// @Produce json
// @Success 200 {array} string
// @Router /tables [get]
func (h *SchemaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tables, err := tableService.GetAll(h.Connection)
	if err != nil {
		logger.Error("Failed to list tables: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to list tables: " + err.Error(), nil)
		return
	}

	JSON(w, http.StatusOK, tables)
}