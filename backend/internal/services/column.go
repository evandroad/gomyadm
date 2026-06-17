package services

import (
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
)

type ColumnService struct {
	Connection *db.ConnectionManager
}

func NewColumnService(conn *db.ConnectionManager) *ColumnService {
	return &ColumnService{
		Connection: conn,
	}
}

func (s *ColumnService) GetAll(table string) (*models.TableSchema, error) {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return nil, err
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

	err = driver.DeleteColumn(conn.DB, table, column)
	if err != nil {
		logger.Error("Failed to delete column: %v", err)
		return err
	}

	return nil
}