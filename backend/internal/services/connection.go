package services

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/evandroad/gomyadm/internal/models"
	"github.com/evandroad/gomyadm/internal/logger"
)

func SaveConnection(cfg models.ConnectionConfig) error {
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		logger.Error("Failed to create data directory: %v", err)
		return err
	}

	filePath := filepath.Join("data", "connections.json")

	var connections []models.ConnectionConfig

	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			logger.Error("Failed to open connections file: %v", err)
			return err
		}

		defer file.Close()

		err = json.NewDecoder(file).Decode(&connections)
		if err != nil {
			logger.Error("Failed to decode connections file: %v", err)
			return err
		}
	}

	connections = append(connections, cfg)

	file, err := os.Create(filePath)
	if err != nil {
		logger.Error("Failed to create connections file: %v", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(connections)
	if err != nil {
		logger.Error("Failed to encode connections: %v", err)
		return err
	}

	return nil
}