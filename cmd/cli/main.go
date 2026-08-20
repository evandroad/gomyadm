package main

import (
	"gomyadm/cmd/cli/commands"
	"gomyadm/cmd/cli/context"
	"gomyadm/internal/db"
	"gomyadm/internal/services"
	"log"
)

func main() {
	if err := services.LoadConnections(); err != nil {
		log.Fatal(err)
	}

	manager := db.NewConnectionManager()

	contextStore, err := context.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	app := &commands.App{
		Connections: services.NewConnectionStore(),
		Session:     services.NewSessionService(manager),
		Context:     contextStore,
	}

	commands.Execute(app)
}

