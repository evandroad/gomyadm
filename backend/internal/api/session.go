package api

import (
	"encoding/json"
	"net/http"

	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/services/session"
	. "github.com/evandroad/gomyadm/internal/respond"

	"github.com/rs/xid"
)

type SessionHandler struct {
	Connection *db.ConnectionManager
}

// @Summary Faz a conexão
// @Tags session
// @Accept json
// @Produce json
// @Param connection body models.ConnectionConfig true "Dados da conexão"
// @Success 200 {object} models.ConnectionResponse
// @Router /session [post]
func (h *SessionHandler) Connect(w http.ResponseWriter, r *http.Request) {
	var cfg models.ConnectionConfig

	err := json.NewDecoder(r.Body).Decode(&cfg)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	if cfg.ID == "" {
		cfg.ID = xid.New().String()
	}

	var data models.ConnectionResponse
	data, err = sessionService.Connect(h.Connection, cfg)
	if err != nil {
		logger.Error("Failed to connect: %v", err)
		Error(w, http.StatusBadRequest, "Failed to connect", nil)
		return
	}

	JSON(w, http.StatusOK, data)
}

// @Summary Desfaz a conexão
// @Tags session
// @Accept json
// @Produce json
// @Success 200 {object} respond.Response
// @Router /session [delete]
func (h *SessionHandler) Disconnect(w http.ResponseWriter, r *http.Request) {
	err := sessionService.Disconnect(h.Connection)
	if err != nil {
		logger.Error("Failed to disconnect: %v", err)
		Error(w, http.StatusNotFound, "Connection not found", nil)
		return
	}

	Success(w, http.StatusOK, "Conexão encerrada com sucesso.", nil)
}

// @Summary Conexão ativa
// @Description Retorna a conexão ativa
// @Tags session
// @Produce json
// @Success 200 {object} models.ConnectionResponse
// @Router /session [get]
func (h *SessionHandler) Active(w http.ResponseWriter, r *http.Request) {
	conn, err := sessionService.Active(h.Connection)
	if err != nil {
		logger.Error("Failed to get active connection: %v", err)
		Error(w, http.StatusNotFound, "No active connection", nil)
		return
	}

	JSON(w, http.StatusOK, conn)
}