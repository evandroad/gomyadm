package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "Gerencia itens das tabelas do bancos de dados",
}

var itemListCmd = &cobra.Command{
	Use:   "ls <table>",
	Short: "Lista os itens das tabelas do bancos de dados",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		table := args[0]

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		tableData, err := app.Item.GetAll(table)
		if err != nil {
			return err
		}

		data := make([][]string, 0, len(tableData.Rows))

		for _, row := range tableData.Rows {
			values := make([]string, 0, len(tableData.Columns))

			for _, column := range tableData.Columns {
				values = append(values, fmt.Sprint(row[column]))
			}

			data = append(data, values)
		}

		printTable(tableData.Columns, data)

		return nil
	},
}

func init() {
	itemCmd.AddCommand(itemListCmd)
	// itemCmd.AddCommand(itemCreateCmd)
	// itemCmd.AddCommand(itemUpdateCmd)
	// itemCmd.AddCommand(itemRemoveCmd)

	rootCmd.AddCommand(itemCmd)
}