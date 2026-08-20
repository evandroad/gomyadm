package commands

import (
	"bufio"
	"fmt"
	"gomyadm/internal/models"
	"os"
	"strconv"

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

var tableCreateCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Cria uma nova tabela",

	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		var columns []models.Column

		tableName, err := prompt(reader, "Nome da tabela")
		if err != nil {
			return err
		}

		for {
			fmt.Printf("\nColuna %d\n", len(columns)+1)

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

			columns = append(columns, models.Column{
				Name:          name,
				Type:          typ,
				Length:        lengthPtr,
				Nullable:      nullable,
				Primary:       primary,
				Unique:        unique,
				AutoIncrement: autoIncrement,
				DefaultValue:  defaultValue,
			})

			fmt.Println()

			more, err := promptBool(reader, "Adicionar outra coluna?", true)
			if err != nil {
				return err
			}

			if !more {
				break
			}
		}

		table := models.Table{
			Name:    tableName,
			Columns: columns,
		}

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		if err := app.Table.Create(table); err != nil {
			return err
		}

		fmt.Printf("Tabela %q criada com sucesso.\n", tableName)

		return nil
	},
}

var tableUpdateCmd = &cobra.Command{
	Use:   "edit <old-name> <new-name>",
	Short: "Edita uma tabela",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		oldName := args[0]
		newName := args[1]

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		if err := app.Table.Update(oldName, newName); err != nil {
			return fmt.Errorf("%w", err)
		}

		fmt.Printf("Tabela %q renomeada para %q com sucesso.\n", oldName, newName)

		return nil
	},
}

var tableRemoveCmd = &cobra.Command{
	Use:   "rm <name>",
	Short: "Apaga uma tabela",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		ctx := app.Context.Get()
		if err := app.Database.Select(ctx.Database); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		if err := app.Table.Delete(name); err != nil {
			return err
		}

		fmt.Printf("Tabela %q apagado com sucesso.\n", name)

		return nil
	},
}

func init() {
	tableCmd.AddCommand(tableListCmd)
	tableCmd.AddCommand(tableCreateCmd)
	tableCmd.AddCommand(tableUpdateCmd)
	tableCmd.AddCommand(tableRemoveCmd)

	rootCmd.AddCommand(tableCmd)
}