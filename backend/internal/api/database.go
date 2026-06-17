package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/services"
	. "github.com/evandroad/gomyadm/internal/respond"
)

type DatabaseHandler struct {
	Service *services.DatabaseService
}

func NewDatabaseHandler(service *services.DatabaseService) *DatabaseHandler {
	return &DatabaseHandler{
		Service: service,
	}
}

// @Summary Seleciona um banco de dados
// @Tags database
// @Accept json
// @Produce json
// @Param database body models.DatabaseRequest true "Banco de Dados"
// @Success 200
// @Router /database/select [post]
func (h *DatabaseHandler) SelectDatabase(w http.ResponseWriter, r *http.Request) {
	var req models.DatabaseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	err = h.Service.Select(req.Database)
	if err != nil {
		logger.Error("Failed to select database: %v", err)
		Error(w, http.StatusBadRequest, "Failed to select database", nil)
		return
	}

	Success(w, http.StatusOK, "Database " + req.Database + " selected successfully.", nil)
}