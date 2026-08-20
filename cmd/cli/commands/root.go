package commands

import (
	"fmt"
	"gomyadm/cmd/cli/context"
	"gomyadm/internal/services"
	"os"

	"github.com/spf13/cobra"
)

type App struct {
	Context    *context.Store
	Version    *services.AppService
	Connection *services.ConnectionsStore
	Session    *services.SessionService
	Database   *services.DatabaseService
	Table      *services.TableService
	Item       *services.ItemService
	Column     *services.ColumnService
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

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}