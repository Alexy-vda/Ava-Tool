package cmd

import (
	"ava-tool/internal/commands"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go service interactively or via flags",
	Long: `Create a new Go service.
You can use flags to skip the interactive prompts, which is useful for CI/CD or AI agents.

Example:
  ava create --name my-service --with-db --with-prometheus --port 9090`,
	Run: func(cmd *cobra.Command, args []string) {
		commands.CreateServiceCommand(cmd)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Define flags
	createCmd.Flags().StringP("name", "n", "", "Name of the service")
	createCmd.Flags().StringP("port", "p", "8080", "Port to listen on")
	createCmd.Flags().Bool("with-db", false, "Include PostgreSQL database (GORM)")
	createCmd.Flags().Bool("with-prometheus", false, "Include Prometheus metrics endpoint")
	createCmd.Flags().Bool("with-sentry", false, "Include Sentry error tracking")
	createCmd.Flags().Bool("with-swagger", false, "Include Swagger documentation config")
}
