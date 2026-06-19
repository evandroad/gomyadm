package services

import (
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
	"github.com/evandroad/gomyadm/internal/models"
)

type DatabaseService struct {
	Connection *db.ConnectionManager
}

func NewDatabaseService(conn *db.ConnectionManager) *DatabaseService {
	return &DatabaseService{
		Connection: conn,
	}
}

func (s *DatabaseService) Select(database string) error {
	err := s.Connection.SelectDatabase(database)
	if err != nil {
		logger.Error("Failed to select database: %v", err)
		return err
	}

	return nil
}

func (s *DatabaseService) GetAll() (models.DatabaseResponse, error) {
	var resp models.DatabaseResponse

	conn, err := s.Connection.Active()
	if err != nil {
		logger.Error("Failed to get all database: %v", err)
		return resp, err
	}

	resp.Active = conn.Database
	resp.Databases = conn.Databases

	return resp, nil
}

func (s *DatabaseService) Create(req models.DatabaseRequest) error {
	driver, conn, err := s.Connection.GetDriverAndConnection()
	if err != nil {
		logger.Error("Failed to get driver and connection: %v", err)
		return err
	}

	err = driver.CreateDatabase(conn.DB, req.Name)
	if err != nil {
		logger.Error("Failed to create database: %v", err)
		return err
	}

	return nil
}