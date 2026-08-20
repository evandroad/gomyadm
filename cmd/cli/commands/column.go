package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var columnCmd = &cobra.Command{
	Use:   "column",
	Short: "Gerencia as colunas de uma tabela",
}

var columnListCmd = &cobra.Command{
	Use:   "ls <table>",
	Short: "Lista as colunas de uma tabela",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		tableName := args[0]

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		table, err := app.Column.GetAll(tableName)
		if err != nil {
			return err
		}

		headers := []string{"NAME", "TYPE", "NULL", "PK", "UNIQUE", "AUTO", "DEFAULT"}

		rows := make([][]string, 0, len(table.Columns))

		for _, column := range table.Columns {
			columnType := column.Type

			if column.Length != nil {
				columnType = fmt.Sprintf("%s(%d)", columnType, *column.Length)
			}

			rows = append(rows, []string{
				column.Name,
				columnType,
				strconv.FormatBool(column.Nullable),
				strconv.FormatBool(column.Primary),
				strconv.FormatBool(column.Unique),
				strconv.FormatBool(column.AutoIncrement),
				column.DefaultValue,
			})
		}

		fmt.Printf("Table: %s\n\n", table.Name)
		printTable(headers, rows)

		return nil
	},
}

func init() {
	columnCmd.AddCommand(columnListCmd)
	// columnCmd.AddCommand(columnCreateCmd)
	// columnCmd.AddCommand(columnUpdateCmd)
	// columnCmd.AddCommand(columnRemoveCmd)

	rootCmd.AddCommand(columnCmd)
}