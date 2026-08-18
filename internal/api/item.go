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

type ItemHandler struct {
	Service *services.ItemService
}

func NewItemHandler(service *services.ItemService) *ItemHandler {
	return &ItemHandler{
		Service: service,
	}
}

// @Summary Lista colunas e itens
// @Description Retorna uma lista com as colunas e uma com os itens
// @Tags item
// @Produce json
// @Param table path string true "Tabela"
// @Success 200 {object} models.TableData
// @Router /table/item/{table} [get]
func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	rows, err := h.Service.GetAll(table)
	if err != nil {
		logger.Error("Failed to select table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to select table", nil)
		return
	}

	JSON(w, http.StatusOK, rows)
}

// @Summary Cria item
// @Description Cria um item em uma tabela
// @Tags item
// @Accept json
// @Produce json
// @Param item body models.ItemRequest true "Item novo"
// @Success 201 {object} respond.Response
// @Router /table/item [post]
func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.ItemRequest	

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.Service.Create(req)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to insert data: " + err.Error(), nil)
		return
	}

	Success(w, http.StatusCreated, "Item inserido com sucesso.", nil)
}

// @Summary Altera item
// @Description Altera um item em uma tabela
// @Tags item
// @Accept json
// @Produce json
// @Param item body models.ItemRequest true "Item"
// @Success 200 {object} respond.Response
// @Router /table/item [put]
func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req models.ItemRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.Service.Update(req)
	if err != nil {
		logger.Error("Failed to update data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to update data", nil)
		return
	}

	Success(w, http.StatusOK, "Item alterado com sucesso.", nil)
}

// @Summary Remove item
// @Description Remove um item de uma tabela
// @Tags item
// @Produce json
// @Param item body models.ItemRequest true "Item"
// @Success 200 {object} models.TableData
// @Router /table/item [delete]
func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var req models.ItemRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.Service.Delete(req)
	if err != nil {
		logger.Error("Failed to delete data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete data", nil)
		return
	}

	Success(w, http.StatusOK, "Item removido com sucesso.", nil)
}