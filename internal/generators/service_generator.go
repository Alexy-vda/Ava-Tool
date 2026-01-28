package generators

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// Embedding all template files
//
//go:embed templates/main.go.tmpl
var mainTemplate string

//go:embed templates/main-db.go.tmpl
var mainDBTemplate string

//go:embed templates/global.go.tmpl
var globalTemplate string

//go:embed templates/health.go.tmpl
var healthTemplate string

//go:embed templates/request.go.tmpl
var requestTemplate string

//go:embed templates/middleware.go.tmpl
var middlewareTemplate string

//go:embed templates/server.go.tmpl
var serverTemplate string

//go:embed templates/init.go.tmpl
var initTemplate string

//go:embed templates/migration.go.tmpl
var migrationTemplate string

//go:embed templates/base.go.tmpl
var baseTemplate string

//go:embed templates/config.go.tmpl
var configTemplate string

//go:embed templates/.env.tmpl
var envTemplate string

//go:embed templates/.gitignore.tmpl
var gitignoreTemplate string

//go:embed templates/Dockerfile.tmpl
var dockerfileTemplate string

//go:embed templates/docker-compose.yaml.tmpl
var dockerComposeTemplate string

//go:embed templates/docker-compose-db.yaml.tmpl
var dockerComposeDBTemplate string

//go:embed templates/README.md.tmpl
var readmeTemplate string

// GenerateService crée la structure des fichiers pour un nouveau service
func GenerateService(name string, includeDB, includePrometheus, includeSwagger bool, port string) error {
	baseDir := name

	// Vérifier si le dossier existe déjà pour éviter l'écrasement
	if _, err := os.Stat(baseDir); !os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' already exists", baseDir)
	}

	directories := []string{
		filepath.Join("cmd", name), "internal/configs", "internal/api/handlers",
		"internal/api/routers", "internal/server", "internal/services", "internal/helpers",
	}
	if includeDB {
		directories = append(directories, "internal/db/", "internal/models", "internal/repositories")
	}

	fmt.Printf("Creating project structure in '%s'...\n", baseDir)

	for _, dir := range directories {
		dirPath := filepath.Join(baseDir, dir)
		// Utilisation de 0755 (rwxr-xr-x) au lieu de ModePerm (0777) pour la sécurité
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
	}

	files := map[string]string{
		filepath.Join(baseDir, "cmd", name, "main.go"):           selectTemplate(includeDB, mainTemplate, mainDBTemplate),
		filepath.Join(baseDir, "internal/api/routers/global.go"): globalTemplate,
		filepath.Join(baseDir, "internal/api/handlers/health.go"): healthTemplate,
		filepath.Join(baseDir, "internal/configs/config.go"):     configTemplate,
		filepath.Join(baseDir, "internal/helpers/request.go"):    requestTemplate,
		filepath.Join(baseDir, "internal/server/middleware.go"):  middlewareTemplate,
		filepath.Join(baseDir, "internal/server/server.go"):      serverTemplate,
		filepath.Join(baseDir, ".env"):                           envTemplate,
		filepath.Join(baseDir, ".gitignore"):                     gitignoreTemplate,
		filepath.Join(baseDir, "Dockerfile"):                     dockerfileTemplate,
		filepath.Join(baseDir, "docker-compose.yaml"):            selectTemplate(includeDB, dockerComposeTemplate, dockerComposeDBTemplate),
		filepath.Join(baseDir, "README.md"):                       readmeTemplate,
	}

	if includeDB {
		files[filepath.Join(baseDir, "internal/db/migration.go")] = migrationTemplate
		files[filepath.Join(baseDir, "internal/db/init.go")] = initTemplate
		files[filepath.Join(baseDir, "internal/models/base.go")] = baseTemplate
	}

	for path, tmplContent := range files {
		if err := generateFileFromTemplate(path, tmplContent, map[string]interface{}{
			"ServiceName":       name,
			"Port":              port,
			"IncludePrometheus": includePrometheus,
			"IncludeSwagger":    includeSwagger,
		}); err != nil {
			return fmt.Errorf("failed to generate file %s: %w", path, err)
		}
	}

	return nil
}

// selectTemplate sélectionne le template selon la présence ou non de la DB
func selectTemplate(includeDB bool, templateNoDB, templateWithDB string) string {
	if includeDB {
		return templateWithDB
	}
	return templateNoDB
}

// generateFileFromTemplate génère un fichier à partir d'un contenu de template et d'une map de données
func generateFileFromTemplate(outputPath, tmplContent string, data map[string]interface{}) error {
	tmpl, err := template.New("template").Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	return nil
}
