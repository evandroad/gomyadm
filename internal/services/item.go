package services

import (
	"fmt"

	"gomyadm/internal/db"
	"gomyadm/internal/logger"
	"gomyadm/internal/models"
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

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return nil, fmt.Errorf("No database selected.")
	}

	rows, err := driver.GetAllItem(conn.DB, table)
	if err != nil {
		logger.Error("Failed to select table: %v", err)
		return nil, err
	}

	return rows, nil
}

func (s *ItemService) GetOne(req models.ItemRequest) (map[string]any, error) {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return nil, err
	}

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return nil, fmt.Errorf("No database selected.")
	}

	rows, err := driver.GetOneItem(conn.DB, req.Table, req.Key)
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

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return fmt.Errorf("No database selected.")
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

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return fmt.Errorf("No database selected.")
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

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return fmt.Errorf("No database selected.")
	}

	err = driver.DeleteItem(conn.DB, req.Table, req.Key)
	if err != nil {
		logger.Error("Failed to delete data: %v", err)
		return err
	}

	return nil
}