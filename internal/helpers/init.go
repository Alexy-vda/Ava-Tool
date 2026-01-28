package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// execGoModInit initialise un module Go et exécute go mod tidy
func ExecGoModInit(baseDir string, name string) error {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		return fmt.Errorf("directory %s does not exist", baseDir)
	}

	// Exécuter go mod init
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = filepath.Join(".", baseDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("go mod init failed: %s: %w", string(output), err)
	}

	// Exécuter go mod tidy dans le répertoire
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = filepath.Join(".", baseDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("go mod tidy failed: %s: %w", string(output), err)
	}

	return nil
}
