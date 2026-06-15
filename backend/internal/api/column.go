package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/go-chi/chi/v5"
)

type ColumnHandler struct {
	Connection *db.ConnectionManager
}

// @Summary Listar colunas de uma tabela
// @Description Retorna uma lista com as colunas e o schema de uma tabela
// @Tags column
// @Produce json
// @Param table path string true "Tabela"
// @Success 200 {object} models.TableSchema
// @Router /tables/column/{table} [get]
func (h *ColumnHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	schema, err := driver.GetAllColumn(conn.DB, table)
	if err != nil {
		logger.Error("Failed to describe table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to describe table", nil)
		return
	}

	JSON(w, http.StatusOK, schema)
}

// @Summary Insere coluna
// @Description Insere uma coluna em uma tabela
// @Tags column
// @Accept json
// @Produce json
// @Param column body models.ColumnRequest true "Coluna novo"
// @Success 201 {object} respond.Response
// @Router /tables/column [post]
func (h *ColumnHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var req models.ColumnRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	err = driver.InsertColumn(conn.DB, req.Table, req.Column)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to insert data", err)
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
// @Router /tables/column [put]
func (h *ColumnHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req models.ColumnRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	err = driver.UpdateColumn(conn.DB, req.Table, req.OldName, req.Column)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to insert data", nil)
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
// @Router /tables/column/{table}/{column} [delete]
func (h *ColumnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")
	column := chi.URLParam(r, "column")

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to get driver and connection", nil)
		return
	}

	err = driver.DeleteColumn(conn.DB, table, column)
	if err != nil {
		logger.Error("Failed to describe table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to describe table", nil)
		return
	}

	Success(w, http.StatusOK, "Coluna removida com sucesso.", nil)
}