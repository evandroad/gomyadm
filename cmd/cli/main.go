package main

import (
	"gomyadm/cmd/cli/commands"
	"gomyadm/cmd/cli/context"
	"gomyadm/internal/db"
	"gomyadm/internal/logger"
	"gomyadm/internal/services"
	"log"
)

func main() {
	if err := services.LoadConnections(); err != nil {
		log.Fatal(err)
	}

	manager := db.NewConnectionManager()
	logger.SetActive(false)

	contextStore, err := context.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	app := &commands.App{
		Context:    contextStore,
		Connection: services.NewConnectionStore(),
		Session:    services.NewSessionService(manager),
		Database:   services.NewDatabaseService(manager),
	}

	commands.Execute(app)
}

