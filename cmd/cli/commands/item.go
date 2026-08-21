package commands

import (
	"bufio"
	"fmt"
	"gomyadm/internal/models"
	"os"
	"strings"

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

var itemCreateCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona um item em uma tabela",

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

		values := make(map[string]any)

		fmt.Printf("\nTabela: %s\n\n", table.Name)

		for _, column := range table.Columns {
			value, err := prompt(reader, column.Name)
			if err != nil {
				return err
			}

			if value == "" {
				if column.DefaultValue != "" {
					continue
				}

				if column.Nullable {
					values[column.Name] = nil
					continue
				}

				return fmt.Errorf("a coluna %q não pode ser vazia", column.Name)
			}

			values[column.Name] = value
		}

		req := models.ItemRequest{
			Table:  tableName,
			Values: values,
		}

		if err := app.Item.Create(req); err != nil {
			return err
		}

		fmt.Printf("\nItem adicionado com sucesso na tabela %q.\n", tableName)

		return nil
	},
}

var itemUpdateCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edita um item em uma tabela",

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

		var primaryColumn *models.Column

		for i := range table.Columns {
			if table.Columns[i].Primary {
				primaryColumn = &table.Columns[i]
				break
			}
		}

		if primaryColumn == nil {
			return fmt.Errorf("a tabela %q não possui chave primária", tableName)
		}

		keyValue, err := prompt(reader, fmt.Sprintf("Valor de %s", primaryColumn.Name))
		if err != nil {
			return err
		}

		if keyValue == "" {
			return fmt.Errorf("o valor da chave primária não pode ser vazio")
		}

		key := map[string]any{
			primaryColumn.Name: keyValue,
		}

		reqGet := models.ItemRequest{
			Table: tableName,
			Key:   key,
		}

		item, err := app.Item.GetOne(reqGet)
		if err != nil {
			return err
		}

		values := make(map[string]any)

		fmt.Printf("\nTabela: %s\n\n", table.Name)

		for _, column := range table.Columns {
			currentValue := item[column.Name]

			defaultValue := ""

			if currentValue != nil {
				defaultValue = fmt.Sprint(currentValue)
			}

			value, err := promptDefault(reader, column.Name, defaultValue)
			if err != nil {
				return err
			}

			if value == defaultValue {
				continue
			}

			if value == "" {
				if column.Nullable {
					values[column.Name] = nil
					continue
				}

				return fmt.Errorf("a coluna %q não pode ser vazia", column.Name)
			}

			values[column.Name] = value
		}

		if len(values) == 0 {
			fmt.Println("\nNenhuma alteração realizada.")
			return nil
		}

		req := models.ItemRequest{
			Table:  tableName,
			Key:    key,
			Values: values,
		}

		if err := app.Item.Update(req); err != nil {
			return err
		}

		fmt.Printf("\nItem da tabela %q editado com sucesso.\n", tableName)

		return nil
	},
}

var itemRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove um item de uma tabela",

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

		key := make(map[string]any)

		for _, column := range table.Columns {
			if !column.Primary {
				continue
			}

			value, err := prompt(reader, fmt.Sprintf("Valor de %s", column.Name))
			if err != nil {
				return err
			}

			if value == "" {
				return fmt.Errorf("o valor da chave primária não pode ser vazio")
			}

			key[column.Name] = value
		}

		if len(key) == 0 {
			return fmt.Errorf("a tabela %q não possui chave primária", tableName)
		}

		reqGet := models.ItemRequest{
			Table: tableName,
			Key:   key,
		}

		item, err := app.Item.GetOne(reqGet)
		if err != nil {
			return err
		}

		headers := make([]string, 0, len(table.Columns))
		values := make([]string, 0, len(table.Columns))

		for _, column := range table.Columns {
			headers = append(headers, column.Name)

			value := item[column.Name]

			if value == nil {
				values = append(values, "NULL")
			} else {
				values = append(values, fmt.Sprint(value))
			}
		}

		fmt.Println("\nItem:")
		printTable(headers, [][]string{values})

		keyDescription := make([]string, 0, len(key))

		for column, value := range key {
			keyDescription = append(
				keyDescription,
				fmt.Sprintf("%s=%v", column, value),
			)
		}

		fmt.Printf(
			"\nRemover item da tabela %q com %s? [N/s]: ",
			tableName,
			strings.Join(keyDescription, ", "),
		)

		confirm, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		confirm = strings.TrimSpace(strings.ToLower(confirm))

		if confirm != "s" && confirm != "sim" {
			fmt.Println("Operação cancelada.")
			return nil
		}

		req := models.ItemRequest{
			Table: tableName,
			Key:   key,
		}

		if err := app.Item.Delete(req); err != nil {
			return err
		}

		fmt.Printf(
			"\nItem removido com sucesso da tabela %q.\n",
			tableName,
		)

		return nil
	},
}

func init() {
	itemCmd.AddCommand(itemListCmd)
	itemCmd.AddCommand(itemCreateCmd)
	itemCmd.AddCommand(itemUpdateCmd)
	itemCmd.AddCommand(itemRemoveCmd)

	rootCmd.AddCommand(itemCmd)
}