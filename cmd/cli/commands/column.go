package commands

import (
	"bufio"
	"fmt"
	"gomyadm/internal/models"
	"os"
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

var columnCreateCmd = &cobra.Command{
	Use:   "add",
	Short: "Cria uma nova coluna numa tabela",

	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		tableName, err := prompt(reader, "Nome da tabela")
		if err != nil {
			return err
		}

		fmt.Printf("\nColuna\n")

		name, err := prompt(reader, "Nome")
		if err != nil {
			return err
		}

		typ, err := prompt(reader, `Tipo (varchar|text|int|bigint|decimal|boolean|date|datetime)`)
		if err != nil {
			return err
		}

		length, err := prompt(reader, "Tamanho")
		if err != nil {
			return err
		}

		var lengthPtr *int

		if length != "" {
			n, err := strconv.Atoi(length)
			if err != nil {
				return fmt.Errorf("tamanho inválido: %s", length)
			}

			lengthPtr = &n
		}

		nullable, err := promptBool(reader, "Nullable", false)
		if err != nil {
			return err
		}

		primary, err := promptBool(reader, "Primary", false)
		if err != nil {
			return err
		}

		unique, err := promptBool(reader, "Unique", false)
		if err != nil {
			return err
		}

		autoIncrement, err := promptBool(reader, "Auto Increment", false)
		if err != nil {
			return err
		}

		defaultValue, err := prompt(reader, "Default")
		if err != nil {
			return err
		}

		column := models.Column{
			Name:          name,
			Type:          typ,
			Length:        lengthPtr,
			Nullable:      nullable,
			Primary:       primary,
			Unique:        unique,
			AutoIncrement: autoIncrement,
			DefaultValue:  defaultValue,
		}

		req := models.ColumnRequest{
			Table:   tableName,
			Column:  column,
		}

		fmt.Println()

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		if err := app.Column.Create(req); err != nil {
			return err
		}

		fmt.Printf("Tabela %q criada com sucesso.\n", tableName)

		return nil
	},
}

var columnUpdateCmd = &cobra.Command{
	Use:   "edit",
	Short: "Editaa uma coluna numa tabela",

	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		tableName, err := prompt(reader, "Nome da tabela")
		if err != nil {
			return err
		}

		table, err := app.Column.GetAll(tableName)
		if err != nil {
			return err
		}

		oldName, err := prompt(reader, "Nome atual da coluna")
		if err != nil {
			return err
		}

		var column models.Column

		for _, col := range table.Columns {
			if col.Name == oldName {
				column = col
				break
			}
		}

		fmt.Printf("\nColuna\n")

		column.Name, err = promptDefault(reader, "Nome", column.Name)
		if err != nil {
			return err
		}

		column.Type, err = promptDefault(reader, `Tipo (varchar|text|int|bigint|decimal|boolean|date|datetime)`, column.Type)
		if err != nil {
			return err
		}

		length := ""

		if column.Length != nil {
			length = strconv.Itoa(*column.Length)
		}

		length, err = promptDefault(reader, "Tamanho", length)
		if err != nil {
			return err
		}

		if length == "" {
			column.Length = nil
		} else {
			n, err := strconv.Atoi(length)
			if err != nil {
				return fmt.Errorf("tamanho inválido: %s", length)
			}

			column.Length = &n
		}

		column.Nullable, err = promptBool(reader, "Nullable", column.Nullable)
		if err != nil {
			return err
		}

		column.Primary, err = promptBool(reader, "Primary", column.Primary)
		if err != nil {
			return err
		}

		column.Unique, err = promptBool(reader, "Unique", column.Unique)
		if err != nil {
			return err
		}

		column.AutoIncrement, err = promptBool(reader, "Auto Increment", column.AutoIncrement)
		if err != nil {
			return err
		}

		column.DefaultValue, err = promptDefault(reader, "Default", column.DefaultValue)
		if err != nil {
			return err
		}

		req := models.ColumnRequest{
			Table:   tableName,
			OldName: oldName,
			Column:  column,
		}

		fmt.Println()

		if err := app.Column.Update(req); err != nil {
			return err
		}

		fmt.Printf("Coluna %q da Tabela %q editada com sucesso.\n", oldName, tableName)

		return nil
	},
}

var columnRemoveCmd = &cobra.Command{
	Use:   "rm <table> <column>",
	Short: "Apaga uma coluna de uma tabela",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		table  := args[0]
		column := args[1]

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		if err := app.Column.Delete(table, column); err != nil {
			return err
		}

		fmt.Printf("Coluna %q da tabela %q apagada com sucesso.\n", column, table)

		return nil
	},
}

func init() {
	columnCmd.AddCommand(columnListCmd)
	columnCmd.AddCommand(columnCreateCmd)
	columnCmd.AddCommand(columnUpdateCmd)
	columnCmd.AddCommand(columnRemoveCmd)

	rootCmd.AddCommand(columnCmd)
}