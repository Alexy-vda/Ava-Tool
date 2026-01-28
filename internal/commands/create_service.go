package commands

import (
	"fmt"
	"ava-tool/internal/generators"
	"ava-tool/internal/helpers"
	"log"
	"regexp"

	"github.com/spf13/cobra"
)

func CreateServiceCommand(cmd *cobra.Command) {
	// Check for non-interactive mode
	nonInteractive, _ := cmd.Flags().GetBool("yes")

	// 1. Service Name
	serviceName, _ := cmd.Flags().GetString("name")
	if serviceName == "" {
		if nonInteractive {
			serviceName = "new-service"
		} else {
			serviceName = helpers.PromptWithDefault("Enter the service name", "new-service")
		}
	}

	// Basic validation for service name
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, serviceName); !matched {
		log.Fatalf("Invalid service name '%s'. Use only alphanumeric characters, dashes, and underscores.", serviceName)
	}

	// 2. Database
	includeDB, _ := cmd.Flags().GetBool("with-db")
	if !nonInteractive && !cmd.Flags().Changed("with-db") {
		includeDB = helpers.PromptYesNo("Do you want to include a PostgreSQL database?")
	}

	// 3. Prometheus
	includePrometheus, _ := cmd.Flags().GetBool("with-prometheus")
	if !nonInteractive && !cmd.Flags().Changed("with-prometheus") {
		includePrometheus = helpers.PromptYesNo("Do you want to include Prometheus metrics?")
	}

	// 4. Sentry
	includeSentry, _ := cmd.Flags().GetBool("with-sentry")
	if !nonInteractive && !cmd.Flags().Changed("with-sentry") {
		includeSentry = helpers.PromptYesNo("Do you want to include Sentry for error tracking?")
	}

	// 5. Swagger
	includeSwagger, _ := cmd.Flags().GetBool("with-swagger")
	if !nonInteractive && !cmd.Flags().Changed("with-swagger") {
		includeSwagger = helpers.PromptYesNo("Do you want to include Swagger documentation?")
	}

	// 6. Port
	port, _ := cmd.Flags().GetString("port")
	if !nonInteractive && !cmd.Flags().Changed("port") {
		port = helpers.PromptWithDefault("Enter the port", "8080")
	}

	// Execution
	if err := generators.GenerateService(serviceName, includeDB, includePrometheus, includeSentry, includeSwagger, port); err != nil {
		log.Fatalf("Error generating service: %v", err)
	}

	fmt.Println("Initializing Go module...")
	if err := helpers.ExecGoModInit(serviceName, serviceName); err != nil {
		log.Printf("Warning: Failed to initialize go module (you may need to run 'go mod init' manually): %v", err)
	} else {
		fmt.Println("✔ Go module initialized.")
	}
	
	fmt.Printf("\n\033[1;32m✔ Service %s created successfully!\033[0m\n", serviceName)
	if includeSwagger {
		fmt.Printf("\033[1;33m⚠ Warning: You included Swagger. Don't forget to run 'swag init' in your project root to generate the docs folder, otherwise the code won't compile!\033[0m\n")
		fmt.Printf("  Install swag: go install github.com/swaggo/swag/cmd/swag@latest\n")
		fmt.Printf("  Run init: cd %s && swag init -g cmd/%s/main.go\n", serviceName, serviceName)
	}
	fmt.Printf("To start your service:\n")
	fmt.Printf("  cd %s && docker-compose up --build\n", serviceName)
}
