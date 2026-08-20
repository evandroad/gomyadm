package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var contextCmd = &cobra.Command{
	Use:   "status",
	Short: "Mostra o status contexto atual",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := app.Context.Get()

		fmt.Println("Context")

		if ctx.ConnectionID == "" {
			fmt.Println("  Connection: -")
			fmt.Println("  Id: -")
			fmt.Println("  Database:   -")
			return nil
		}

		connection, ok := app.Connection.GetByID(ctx.ConnectionID)
		if !ok {
			fmt.Printf("  Connection: %s (não encontrada)\n", ctx.ConnectionID)
			fmt.Printf("  Id:         %s (não encontrada)\n", ctx.ConnectionID)
		} else {
			fmt.Printf("  Connection: %s\n", connection.Name)
			fmt.Printf("  Id:         %s\n", connection.ID)
		}

		if ctx.Database == "" {
			fmt.Println("  Database:   -")
		} else {
			fmt.Printf("  Database:   %s\n", ctx.Database)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)
}