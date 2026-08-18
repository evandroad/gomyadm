package services

import (
	"fmt"

	"gomyadm/internal/db"
	"gomyadm/internal/logger"
	"gomyadm/internal/models"
)

type ColumnService struct {
	Connection *db.ConnectionManager
}

func NewColumnService(conn *db.ConnectionManager) *ColumnService {
	return &ColumnService{
		Connection: conn,
	}
}

func (s *ColumnService) GetAll(table string) (*models.Table, error) {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return nil, err
	}

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return nil, fmt.Errorf("No database selected.")
	}

	schema, err := driver.GetAllColumn(conn.DB, table)
	if err != nil {
		logger.Error("Failed to describe table: %v", err)
		return nil, err
	}

	return schema, nil
}

func (s *ColumnService) Create(req models.ColumnRequest) error {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return fmt.Errorf("No database selected.")
	}

	err = driver.CreateColumn(conn.DB, req.Table, req.Column)
	if err != nil {
		logger.Error("Failed to create data: %v", err)
		return err
	}

	return nil
}

func (s *ColumnService) Update(req models.ColumnRequest) error {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return fmt.Errorf("No database selected.")
	}

	err = driver.UpdateColumn(conn.DB, req.Table, req.OldName, req.Column)
	if err != nil {
		logger.Error("Failed to update data: %v", err)
		return err
	}

	return nil
}

func (s *ColumnService) Delete(table, column string) error {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	if s.Connection.GetDatabase() == "" {
		logger.Error("No database selected.")
		return fmt.Errorf("No database selected.")
	}

	err = driver.DeleteColumn(conn.DB, table, column)
	if err != nil {
		logger.Error("Failed to delete column: %v", err)
		return err
	}

	return nil
}