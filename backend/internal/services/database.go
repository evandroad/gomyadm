package services

import (
	"github.com/evandroad/gomyadm/internal/db"
	"github.com/evandroad/gomyadm/internal/logger"
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