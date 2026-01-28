package installer

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed templates/SKILL.md.tmpl
var skillTemplate string

func InstallClaudeSkill() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not find user home directory: %w", err)
	}

	skillDir := filepath.Join(homeDir, ".claude", "skills", "ava-create")
	skillFile := filepath.Join(skillDir, "SKILL.md")

	fmt.Printf("Installing Claude Skill to: %s\n", skillDir)

	// Create directory
	if err := os.MkdirAll(skillDir, 0755); err != nil {
		return fmt.Errorf("failed to create skill directory: %w", err)
	}

	// Write file
	file, err := os.Create(skillFile)
	if err != nil {
		return fmt.Errorf("failed to create SKILL.md file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(skillTemplate)
	if err != nil {
		return fmt.Errorf("failed to write skill content: %w", err)
	}

	fmt.Println("\033[1;32m✔ Claude Skill 'ava-create' installed successfully!\033[0m")
	fmt.Println("You can now ask Claude: 'Create a new service with db and metrics'")
	return nil
}
