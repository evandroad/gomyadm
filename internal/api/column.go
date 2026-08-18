package api

import (
	"encoding/json"
	"net/http"

	"gomyadm/internal/logger"
	"gomyadm/internal/models"
	"gomyadm/internal/services"
	. "gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
)

type ColumnHandler struct {
	Service *services.ColumnService
}

func NewColumnHandler(service *services.ColumnService) *ColumnHandler {
	return &ColumnHandler{
		Service: service,
	}
}

// @Summary Listar colunas de uma tabela
// @Description Retorna uma lista com as colunas e o schema de uma tabela
// @Tags column
// @Produce json
// @Param table path string true "Tabela"
// @Success 200 {object} models.Table
// @Router /table/column/{table} [get]
func (h *ColumnHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	schema, err := h.Service.GetAll(table)
	if err != nil {
		logger.Error("Failed to describe table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to describe table: " + err.Error(), nil)
		return
	}

	JSON(w, http.StatusOK, schema)
}

// @Summary Cria coluna
// @Description Cria uma coluna em uma tabela
// @Tags column
// @Accept json
// @Produce json
// @Param column body models.ColumnRequest true "Coluna novo"
// @Success 201 {object} respond.Response
// @Router /table/column [post]
func (h *ColumnHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.ColumnRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err = h.Service.Create(req)
	if err != nil {
		logger.Error("Failed to create column: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to create column: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "Coluna salva com sucesso.", nil)
}

// @Summary Altera coluna
// @Description Altera uma coluna em uma tabela
// @Tags column
// @Accept json
// @Produce json
// @Param column body models.ColumnRequest true "Coluna"
// @Success 200 {object} respond.Response
// @Router /table/column [put]
func (h *ColumnHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req models.ColumnRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.Service.Update(req)
	if err != nil {
		logger.Error("Failed to update column: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to update column: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "Coluna alterada com sucesso.", nil)
}

// @Summary Remove coluna
// @Description Remove uma coluna de uma tabela
// @Tags column
// @Produce json
// @Param table path string true "Tabela"
// @Param column path string true "Coluna"
// @Success 200 {object} respond.Response
// @Router /table/column/{table}/{column} [delete]
func (h *ColumnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")
	column := chi.URLParam(r, "column")

	err := h.Service.Delete(table, column)
	if err != nil {
		logger.Error("Failed to delete column: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete column: " + err.Error(), nil)
		return
	}	

	Success(w, http.StatusOK, "Coluna removida com sucesso.", nil)
}