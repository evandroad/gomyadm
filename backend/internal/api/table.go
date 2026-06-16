package api

import (
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/evandroad/gomyadm/internal/services/table"
)

type TableHandler struct {
	Connection *db.ConnectionManager
}

// @Summary Lista de tabelas
// @Description Retorna todas as tabelas do banco selecionado
// @Tags table
// @Produce json
// @Success 200 {array} string
// @Router /table [get]
func (h *TableHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tables, err := tableService.GetAll(h.Connection)
	if err != nil {
		logger.Error("Failed to list tables: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to list tables: " + err.Error(), nil)
		return
	}

	JSON(w, http.StatusOK, tables)
}

// @Summary Cria tabela
// @Description Cria uma tabela
// @Tags table
// @Accept json
// @Produce json
// @Param tabela body models.TableRequest true "Tabela nova"
// @Success 201 {object} respond.Response
// @Router /table [post]
func (h *TableHandler) Create(w http.ResponseWriter, r *http.Request) {
}