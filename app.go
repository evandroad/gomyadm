package main

import (
	"context"
	"gomyadm/internal/bindings"
	"gomyadm/internal/db"
	"gomyadm/internal/services"
)

type App struct {
	ctx   context.Context
	binds []any
}

type ContextAware interface {
	SetContext(context.Context)
}

func NewApp() *App {
	manager := db.NewConnectionManager()

	appService := services.NewAppService()
	connectionService := services.NewConnectionStore()
	databaseService   := services.NewDatabaseService(manager)

	return &App{
		binds: []any{
			bindings.NewApp(appService),
			bindings.NewConnection(connectionService),
			bindings.NewDatabase(databaseService),
		},
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	for _, bind := range a.binds {
		if b, ok := bind.(ContextAware); ok {
			b.SetContext(ctx)
		}
	}
}

func (a *App) Bindings() []any {
	return a.binds
}
