package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Mostra a versão do gomyadm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Version.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}