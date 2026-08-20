package commands

import (
	"bufio"
	"fmt"
	"gomyadm/internal/models"
	"os"
	"strconv"
	"strings"

	"github.com/rs/xid"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Gerencia conexões",
}

var connectionListCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lista as conexões",
	RunE: func(cmd *cobra.Command, args []string) error {
		connections := app.Connection.GetAll()
		rows := [][]string{}

		for _, connection := range connections {
			rows = append(rows, []string{
				connection.ID,
				connection.Name,
				connection.Driver,
				fmt.Sprintf("%s:%d", connection.Host, connection.Port),
				connection.Database,
			})
		}

		printTable(
			[]string{"ID", "NOME", "DRIVER", "HOST", "BANCO"},
			rows,
		)

		return nil
	},
}

var connectionGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Obtem detalhes de uma conexão",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		connection, ok := app.Connection.GetByID(id)
		if !ok {
			return fmt.Errorf("conexão não encontrada: %s", args[0])
		}

		fmt.Println("Conexão:")
		fmt.Printf("  ID:       %s\n", connection.ID)
		fmt.Printf("  Nome:     %s\n", connection.Name)
		fmt.Printf("  Driver:   %s\n", connection.Driver)
		fmt.Printf("  Host:     %s:%d\n", connection.Host, connection.Port)
		fmt.Printf("  Database: %s\n", connection.Database)

		return nil
	},
}

var connectionCreateCmd = &cobra.Command{
	Use:   "add",
	Short: "Cria uma nova conexão",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		name, err := prompt(reader, "Nome")
		if err != nil {
			return err
		}

		driver, err := prompt(reader, "Driver (mysql|postgres)")
		if err != nil {
			return err
		}

		host, err := prompt(reader, "Host")
		if err != nil {
			return err
		}

		port, err := prompt(reader, "Porta")
		if err != nil {
			return err
		}

		portNumber, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("porta inválida: %s", port)
		}

		username, err := prompt(reader, "Usuário")
		if err != nil {
			return err
		}

		fmt.Print("Senha: ")
		passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}

		fmt.Println()
		password := string(passwordBytes)

		database, err := prompt(reader, "Database")
		if err != nil {
			return err
		}

		connection := models.ConnectionConfig{
			ID:       xid.New().String(),
			Name:     name,
			Driver:   driver,
			Host:     host,
			Port:     portNumber,
			Username: username,
			Password: password,
			Database: database,
		}

		_, err = app.Session.Connect(connection)
		if err != nil {
			return fmt.Errorf("não foi possível conectar: %w", err)
		}

		if err := app.Session.Disconnect(); err != nil {
			return fmt.Errorf("não foi possível desconectar: %w", err)
		}

		if err := app.Connection.Create(connection); err != nil {
			return err
		}

		fmt.Printf("\nConexão %q criada com sucesso.\n", connection.Name)

		return nil
	},
}

var connectionUpdateCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "Edita uma conexão",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		connection, ok := app.Connection.GetByID(id)
		if !ok {
			return fmt.Errorf("conexão não encontrada: %s", id)
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("\nEditando conexão: %s\n\n", connection.Name)
		var err error

		connection.Name, err = promptDefault(reader, "Nome", connection.Name)
		if err != nil {
			return err
		}

		connection.Driver, err = promptDefault(reader, "Driver", connection.Driver)
		if err != nil {
			return err
		}

		connection.Host, err = promptDefault(reader, "Host", connection.Host)
		if err != nil {
			return err
		}

		port, err := promptDefault(reader, "Porta", strconv.Itoa(connection.Port))
		if err != nil {
			return err
		}

		connection.Port, err = strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("porta inválida: %s", port)
		}

		connection.Username, err = promptDefault(reader, "Usuário", connection.Username)
		if err != nil {
			return err
		}

		fmt.Print("Senha [não alterar]: ")

		passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}

		fmt.Println()

		if len(passwordBytes) > 0 {
			connection.Password = string(passwordBytes)
		}

		connection.Database, err = promptDefault(reader, "Database", connection.Database)
		if err != nil {
			return err
		}

		if err := app.Connection.Update(id, connection); err != nil {
			return fmt.Errorf("não foi possível atualizar a conexão: %w", err)
		}

		fmt.Printf("\nConexão %q atualizada com sucesso.\n", connection.Name)

		return nil
	},
}

var connectionDeleteCmd = &cobra.Command{
	Use:   "rm <id>",
	Short: "Remove uma conexão",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		connection, ok := app.Connection.GetByID(id)
		if !ok {
			return fmt.Errorf("conexão não encontrada: %s", args[0])
		}

		err := app.Connection.Delete(args[0])
		if err != nil {
			return fmt.Errorf("não foi possível remover a conxão: %w", err)
		}
		
		ctx := app.Context.Get()

		if ctx.ConnectionID == connection.ID {
			if err := app.Context.Clear(); err != nil {
				return fmt.Errorf("não foi possível limpar o contexto: %w", err)
			}
		}

		fmt.Printf("Conexão %q removida com sucesso.\n", connection.Name)

		return nil
	},
}

var connectionSelectCmd = &cobra.Command{
	Use:   "use <id>",
	Short: "Seleciona uma conexão",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		connection, ok := app.Connection.GetByID(args[0])
		if !ok {
			return fmt.Errorf("conexão não encontrada: %s", args[0])
		}

		// Testa a conexão antes de alterar o contexto.
		_, err := app.Session.Connect(connection)
		if err != nil {
			return fmt.Errorf("não foi possível conectar: %w", err)
		}

		err = app.Session.Disconnect()
		if err != nil {
			return fmt.Errorf("não foi possível desconectar: %w", err)
		}

		// Só chega aqui se a conexão funcionou.
		if err := app.Context.SetConnection(connection.ID); err != nil {
			return err
		}

		fmt.Printf("Selected connection: %s\n", connection.Name)

		return nil
	},
}

func init() {
	connectionCmd.AddCommand(connectionListCmd)
	connectionCmd.AddCommand(connectionGetCmd)
	connectionCmd.AddCommand(connectionCreateCmd)
	connectionCmd.AddCommand(connectionUpdateCmd)
	connectionCmd.AddCommand(connectionDeleteCmd)
	connectionCmd.AddCommand(connectionSelectCmd)

	rootCmd.AddCommand(connectionCmd)
}

func prompt(reader *bufio.Reader, label string) (string, error) {
	fmt.Printf("%s: ", label)

	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}

func promptDefault(reader *bufio.Reader, label, current string) (string, error) {
	if current != "" {
		fmt.Printf("%s [%s]: ", label, current)
	} else {
		fmt.Printf("%s: ", label)
	}

	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	value = strings.TrimSpace(value)

	if value == "" {
		return current, nil
	}

	return value, nil
}