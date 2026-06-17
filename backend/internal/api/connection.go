package api

import (
	"encoding/json"
	"net/http"

	_ "github.com/evandroad/gomyadm/docs"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/services"
	. "github.com/evandroad/gomyadm/internal/respond"

	"github.com/go-chi/chi/v5"
)

type ConnectionHandler struct {
	Service *services.ConnectionsStore
}

func NewConnectionHandler(service *services.ConnectionsStore) *ConnectionHandler {
	return &ConnectionHandler{
		Service: service,
	}
}

// @Summary Lista conexões salvas
// @Description Retorna todas as conexões cadastradas
// @Tags connection
// @Produce json
// @Success 200 {object} models.ConnectionConfig
// @Router /connection [get]
func (h *ConnectionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, h.Service.GetAll())
}

// @Summary Salva conexão
// @Tags connection
// @Accept json
// @Produce json
// @Param connection body models.ConnectionConfig true "Dados da conexão"
// @Success 201 {object} respond.Response
// @Router /connection [post]
func (h *ConnectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.Service.Create(cfg)
	if err != nil {
		logger.Error("Failed to save connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to save connection", nil)
		return
	}

	Success(w, http.StatusCreated, "Connection saved.", nil)
}

// @Summary Altera conexão
// @Tags connection
// @Accept json
// @Produce json
// @Param connection body models.ConnectionConfig true "Dados da conexão"
// @Success 200 {object} respond.Response
// @Router /connection [put]
func (h *ConnectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.Service.Update(cfg.ID, cfg)
	if err != nil {
		logger.Error("Failed to update connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to update connection", nil)
		return
	}

	Success(w, http.StatusOK, "Connection updated.", nil)
}

// @Summary Remove conexão
// @Tags connection
// @Accept json
// @Produce json
// @Param id path string true "ID da conexão"
// @Success 200 {object} respond.Response
// @Failure 404 {object} respond.Response
// @Router /connection/{id} [delete]
func (h *ConnectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.Service.Delete(id)
	if err != nil {
		logger.Error("Failed to delete connection: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete connection", nil)
		return
	}

	Success(w, http.StatusOK, "Connection deleted.", nil)
}