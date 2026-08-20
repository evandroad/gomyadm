package commands

import (
	"fmt"
	"gomyadm/internal/models"

	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Gerencia bancos de dados",
}

var databaseListCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lista os bancos de dados",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		databases, err := app.Database.GetAll()
		if err != nil {
			return err
		}

		for _, db := range databases.Databases {
			fmt.Println(db)
		}

		return nil
	},
}

var databaseSelectCmd = &cobra.Command{
	Use:   "use <name>",
	Short: "Seleciona um banco de dados",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()
		
		if err := app.Database.Select(name); err != nil {
			return fmt.Errorf("não foi possível usar: %w", err)
		}

		if err := app.Context.SetDatabase(name); err != nil {
			return err
		}

		fmt.Printf("Selected database: %s\n", name)

		return nil
	},
}

var databaseCreateCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Cria um novo banco de dados",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		database := models.DatabaseRequest{
			Name:     name,
		}

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		if err := app.Database.Create(database); err != nil {
			return fmt.Errorf("%w", err)
		}

		fmt.Printf("Banco de dados %q criado com sucesso.\n", name)

		return nil
	},
}

var databaseUpdateCmd = &cobra.Command{
	Use:   "edit <old-name> <new-name>",
	Short: "Edita um banco de dados",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		oldName := args[0]
		newName := args[1]

		database := models.DatabaseRequest{
			OldName: oldName,
			NewName: newName,
		}

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		if err := app.Database.Update(database); err != nil {
			return fmt.Errorf("%w", err)
		}

		fmt.Printf("Banco de dados %q renomeado para %q com sucesso.\n", oldName, newName)

		return nil
	},
}

var databaseRemoveCmd = &cobra.Command{
	Use:   "rm <name>",
	Short: "Apaga um banco de dados",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := connectSelected(); err != nil {
			return err
		}
		defer app.Session.Disconnect()

		if err := app.Database.Delete(name); err != nil {
			return err
		}

		fmt.Printf("Banco de dados %q apagado com sucesso.\n", name)

		return nil
	},
}

func init() {
	databaseCmd.AddCommand(databaseListCmd)
	databaseCmd.AddCommand(databaseSelectCmd)
	databaseCmd.AddCommand(databaseCreateCmd)
	databaseCmd.AddCommand(databaseUpdateCmd)
	databaseCmd.AddCommand(databaseRemoveCmd)

	rootCmd.AddCommand(databaseCmd)
}