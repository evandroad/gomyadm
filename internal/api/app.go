package api

import (
	"net/http"

	"gomyadm/internal/services"
	. "gomyadm/internal/respond"
)

type AppHandler struct {
	Service *services.AppService
}

func NewAppHandler(service *services.AppService) *AppHandler {
	return &AppHandler{
		Service: service,
	}
}

// @Summary Obtem a versão
// @Description Retorna a versão
// @Tags app
// @Produce json
// @Success 200 {object} map[string]string
// @Router /version [get]
func (h *AppHandler) Version(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, H{ "version": h.Service.Version() } )
}