package commands

import (
	"fmt"
	"gomyadm/cmd/cli/context"
	"gomyadm/internal/models"
	"gomyadm/internal/services"
	"os"

	"github.com/spf13/cobra"
)

type App struct {
	Version    *services.AppService
	Connection *services.ConnectionsStore
	Session    *services.SessionService
	Context    *context.Store
	Database   *services.DatabaseService
}

var app *App

var rootCmd = &cobra.Command{
	Use:   "gomyadm",
	Short: "Database Manager CLI",
}

func Execute(a *App) {
	app = a

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getSelectedConnection() (models.ConnectionConfig, error) {
	ctx := app.Context.Get()

	if ctx.ConnectionID == "" {
		return models.ConnectionConfig{}, fmt.Errorf(
			"nenhuma conexão selecionada; use 'gomyadm connection use <id>'",
		)
	}

	connection, ok := app.Connection.GetByID(ctx.ConnectionID)
	if !ok {
		return models.ConnectionConfig{}, fmt.Errorf(
			"conexão selecionada não encontrada: %s",
			ctx.ConnectionID,
		)
	}

	return connection, nil
}

func connectSelected() error {
	connection, err := getSelectedConnection()
	if err != nil {
		return err
	}

	_, err = app.Session.Connect(connection)
	return err
}