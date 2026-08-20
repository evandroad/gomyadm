package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Gerencia tabelas do bancos de dados",
}

var tableListCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lista as tabelas do bancos de dados",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		tables, err := app.Table.GetAll()
		if err != nil {
			return err
		}

		for _, table := range tables {
			fmt.Println(table)
		}

		return nil
	},
}

func init() {
	tableCmd.AddCommand(tableListCmd)
	// tableCmd.AddCommand(tableCreateCmd)
	// tableCmd.AddCommand(tableUpdateCmd)
	// tableCmd.AddCommand(tableRemoveCmd)

	rootCmd.AddCommand(tableCmd)
}