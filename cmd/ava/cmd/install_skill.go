package cmd

import (
	"ava-tool/internal/installer"
	"log"

	"github.com/spf13/cobra"
)

var installSkillCmd = &cobra.Command{
	Use:   "install-skill",
	Short: "Install the 'ava-create' skill for Claude Code",
	Long:  `Installs the 'ava-create' skill into ~/.claude/skills/. This allows Claude Code to natively understand and use Ava-Tool to generate services.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := installer.InstallClaudeSkill(); err != nil {
			log.Fatalf("Failed to install skill: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installSkillCmd)
}
