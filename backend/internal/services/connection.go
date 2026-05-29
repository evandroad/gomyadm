package services

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/evandroad/gomyadm/internal/models"
)

func SaveConnection(cfg models.ConnectionConfig) error {
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		log.Printf("Failed to create data directory: %v", err)
		return err
	}

	filePath := filepath.Join("data", "connections.json")

	var connections []models.ConnectionConfig

	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			log.Printf("Failed to open connections file: %v", err)
			return err
		}

		defer file.Close()

		err = json.NewDecoder(file).Decode(&connections)
		if err != nil {
			log.Printf("Failed to decode connections file: %v", err)
			return err
		}
	}

	connections = append(connections, cfg)

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create connections file: %v", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(connections)
	if err != nil {
		log.Printf("Failed to encode connections: %v", err)
		return err
	}

	return nil
}