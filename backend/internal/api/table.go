package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/evandroad/gomyadm/internal/services"
)

type TableHandler struct {
	Service *services.TableService
}

func NewTableHandler(service *services.TableService) *TableHandler {
	return &TableHandler{
		Service: service,
	}
}

// @Summary Lista de tabelas
// @Description Retorna todas as tabelas do banco selecionado
// @Tags table
// @Produce json
// @Success 200 {array} string
// @Router /table [get]
func (h *TableHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tables, err := h.Service.GetAll()
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
// @Param tabela body models.Table true "Tabela nova"
// @Success 201 {object} respond.Response
// @Router /table [post]
func (h *TableHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.Table

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err = h.Service.Create(req)
	if err != nil {
		logger.Error("Failed to create table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to create table: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "Tabela salva com sucesso.", nil)
}