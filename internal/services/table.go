package services

import (
	"gomyadm/internal/db"
	"gomyadm/internal/logger"
	"gomyadm/internal/models"
)

type TableService struct {
	Connection *db.ConnectionManager
}

func NewTableService(conn *db.ConnectionManager) *TableService {
	return &TableService{
		Connection: conn,
	}
}

func (s *TableService) GetAll() ([]string, error) {
	driver, conn, err := s.Connection.GetDriverAndConnection(true)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return []string{}, err
	}

	tables, err := driver.GetAllTable(conn.DB)
	if err != nil {
		logger.Error("Failed to list tables: %v", err)
		return []string{}, err
	}

	return tables, nil
}

func (s *TableService) Create(req models.Table) error {
	driver, conn, err := s.Connection.GetDriverAndConnection(true)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	err = driver.CreateTable(conn.DB, req)
	if err != nil {
		logger.Error("Failed to create table: %v", err)
		return err
	}

	return nil
}

func (s *TableService) Update(oldName, newName string) error {
	driver, conn, err := s.Connection.GetDriverAndConnection(true)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	err = driver.UpdateTable(conn.DB, oldName, newName)
	if err != nil {
		logger.Error("Failed to update table: %v", err)
		return err
	}

	return nil
}

func (s *TableService) Delete(table string) error {
	driver, conn, err := s.Connection.GetDriverAndConnection(true)
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	err = driver.DeleteTable(conn.DB, table)
	if err != nil {
		logger.Error("Failed to delete table: %v", err)
		return err
	}

	return nil
}