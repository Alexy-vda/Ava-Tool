package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ava",
	Short: "Ava Tool is a modern Go service generator",
	Long:  `Ava Tool helps you quickly scaffold Go microservices with best practices and DevOps integration.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
