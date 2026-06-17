package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/services"
	. "github.com/evandroad/gomyadm/internal/respond"
)

type QueryHandler struct {
	Service *services.QueryService
}

func NewQueryHandler(service *services.QueryService) *QueryHandler {
	return &QueryHandler{
		Service: service,
	}
}

// @Summary Executa query
// @Tags query
// @Accept json
// @Produce json
// @Param connection body models.QueryRequest true "Dados da query"
// @Success 200 {object} respond.Response
// @Router /query [post]
func (h *QueryHandler) ExecuteQuery(w http.ResponseWriter, r *http.Request) {
	var req models.QueryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode query request: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	result, err := h.Service.ExecuteQuery(req.Query)
	if err != nil {
		logger.Error("Failed to execute query: %v", err)
		Error(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	JSON(w, http.StatusOK, result)
}