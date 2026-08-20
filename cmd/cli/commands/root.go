package commands

import (
	"gomyadm/cmd/cli/context"
	"gomyadm/internal/services"
	"os"

	"github.com/spf13/cobra"
)

type App struct {
	Version     *services.AppService
	Connections *services.ConnectionsStore
	Session     *services.SessionService
	Context     *context.Store
}

var app *App

var rootCmd = &cobra.Command{
	Use:   "gomyadm",
	Short: "Database Manager CLI",
}

func Execute(a *App) {
	app = a
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}