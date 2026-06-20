package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	. "github.com/evandroad/gomyadm/internal/respond"
	"github.com/evandroad/gomyadm/internal/services"
	"github.com/go-chi/chi/v5"
)

type DatabaseHandler struct {
	Service *services.DatabaseService
}

func NewDatabaseHandler(service *services.DatabaseService) *DatabaseHandler {
	return &DatabaseHandler{
		Service: service,
	}
}

// @Summary Lista de banco de dados
// @Description Retorna todas os bancos de dados
// @Tags database
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router /database [get]
func (h *DatabaseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tables, err := h.Service.GetAll()
	if err != nil {
		logger.Error("Failed to list tables: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to list tables: " + err.Error(), nil)
		return
	}

	JSON(w, http.StatusOK, tables)
}

// @Summary Seleciona um banco de dados
// @Tags database
// @Accept json
// @Produce json
// @Param database body models.DatabaseRequest true "Banco de Dados"
// @Success 200
// @Router /database/select [post]
func (h *DatabaseHandler) Select(w http.ResponseWriter, r *http.Request) {
	var req models.DatabaseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body: " + err.Error(), nil)
		return
	}
	
	err = h.Service.Select(req.Name)
	if err != nil {
		logger.Error("Failed to select database: %v", err)
		Error(w, http.StatusBadRequest, "Failed to select database: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "Database " + req.Name + " selected successfully.", nil)
}

// @Summary Cria banco de dados
// @Description Cria um banco de dados
// @Tags database
// @Accept json
// @Produce json
// @Param database body models.DatabaseRequest true "Banco de Dados"
// @Success 201
// @Router /database [post]
func (h *DatabaseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.DatabaseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	err = h.Service.Create(req)
	if err != nil {
		logger.Error("Failed to create database: %v", err)
		Error(w, http.StatusBadRequest, "Failed to create database: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "Database " + req.Name + " created successfully.", nil)
}

// @Summary Altera banco de dados
// @Description Altera um banco de dados
// @Tags database
// @Accept json
// @Produce json
// @Param database body models.DatabaseRequest true "Banco de Dados"
// @Success 200
// @Router /database [put]
func (h *DatabaseHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req models.DatabaseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	err = h.Service.Update(req)
	if err != nil {
		logger.Error("Failed to update database: %v", err)
		Error(w, http.StatusBadRequest, "Failed to update database: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "Database " + req.Name + " updated successfully.", nil)
}

// @Summary Remove o banco de dados
// @Description Remove o banco de dados
// @Tags table
// @Accept json
// @Produce json
// @Param name path string true "Nome da tabela"
// @Success 200 {object} respond.Response
// @Router /table/{name} [delete]
func (h *DatabaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	err := h.Service.Delete(name)
	if err != nil {
		logger.Error("Failed to delete database: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete database: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "Banco de dados removido com sucesso.", nil)
}