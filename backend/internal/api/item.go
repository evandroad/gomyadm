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

type ItemHandler struct {
	Connection *db.ConnectionManager
}

// @Summary Lista colunas e itens
// @Description Retorna uma lista com as colunas e uma com os itens
// @Tags item
// @Produce json
// @Param table path string true "Tabela"
// @Success 200 {object} models.TableData
// @Router /tables/item/{table} [get]
func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	table := chi.URLParam(r, "table")

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "No connection established.", nil)
		return
	}

	rows, err := driver.GetAllItem(conn.DB, table)
	if err != nil {
		logger.Error("Failed to select table: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to select table", nil)
		return
	}

	JSON(w, http.StatusOK, rows)
}

// @Summary Insere item
// @Description Insere um item em uma tabela
// @Tags item
// @Accept json
// @Produce json
// @Param item body models.ItemRequest true "Item novo"
// @Success 201 {object} respond.Response
// @Router /tables/item [post]
func (h *ItemHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var req models.ItemRequest	

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "No connection established.", nil)
		return
	}

	err = driver.InsertItem(conn.DB, req.Table, req.Values)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to insert data", nil)
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
// @Router /tables/item [put]
func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req models.ItemRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "No connection established.", nil)
		return
	}

	err = driver.UpdateItem(conn.DB, req.Table, req.Key, req.Values)
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
// @Router /tables/item [delete]
func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var req models.ItemRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Failed to decode request body: %v", err)
		Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	driver, conn, err := h.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		Error(w, http.StatusInternalServerError, "No connection established.", nil)
		return
	}

	err = driver.DeleteItem(conn.DB, req.Table, req.Key)
	if err != nil {
		logger.Error("Failed to delete data: %v", err)
		Error(w, http.StatusInternalServerError, "Failed to delete data", nil)
		return
	}

	Success(w, http.StatusOK, "Item removido com sucesso.", nil)
}