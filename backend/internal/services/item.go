package services

import (
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
)

type ItemService struct {
	Connection *db.ConnectionManager
}

func NewItemService(conn *db.ConnectionManager) *ItemService {
	return &ItemService{
		Connection: conn,
	}
}

func (s *ItemService) GetAll(table string) (*models.TableData, error) {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return nil, err
	}

	rows, err := driver.GetAllItem(conn.DB, table)
	if err != nil {
		logger.Error("Failed to select table: %v", err)
		return nil, err
	}

	return rows, nil
}

func (s *ItemService) Create(req models.ItemRequest) error {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	err = driver.CreateItem(conn.DB, req.Table, req.Values)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
		return err
	}

	return nil
}

func (s *ItemService) Update(req models.ItemRequest) error {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	err = driver.UpdateItem(conn.DB, req.Table, req.Key, req.Values)
	if err != nil {
		logger.Error("Failed to update data: %v", err)
		return err
	}

	return nil
}

func (s *ItemService) Delete(req models.ItemRequest) error {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	err = driver.DeleteItem(conn.DB, req.Table, req.Key)
	if err != nil {
		logger.Error("Failed to delete data: %v", err)
		return err
	}

	return nil
}