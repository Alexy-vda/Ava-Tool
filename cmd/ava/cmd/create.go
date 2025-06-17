package cmd

import (
	"go-init-service/internal/commands"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go service interactively",
	Run: func(cmd *cobra.Command, args []string) {
		commands.CreateServiceCommand()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
