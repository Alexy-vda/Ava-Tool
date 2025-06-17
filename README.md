# Ava Tool

Ava Tool is a CLI utility designed to accelerate the creation of Go services. It generates a ready-to-use service structure, including a Gin-based router, optional Gorm ORM setup (with PostgreSQL), Docker & Docker Compose integration, and a GitLab CI configuration.

---

## 📚 Documentation

The full documentation is available online:  
👉 [See the documentation site](https://alexy-vda.github.io/Ava-Tool/)

---

## Features

- **Go service generator**: Instantly scaffold a Go service with a modern, production-ready structure.
- **Configurable**: Choose to include PostgreSQL (Gorm) or not, and set your service port.
- **DevOps ready**: Docker, Docker Compose, and GitLab CI configuration out of the box.
- **Interactive CLI**: Answer a few questions and your service is ready to go.

## Quick Start

1. **Download the latest release** from [GitHub Releases](https://github.com/alexyvda/ava-tool/releases) and add it to your PATH.
2. Run:
   ```bash
   ava-tool
   ```
3. Answer the prompts:
   - Service name (default: `new service`)
   - Add PostgreSQL? (y/n)
   - Service port (default: 8080)
4. Your new service is generated in a dedicated folder.
5. Start it with:
   ```bash
   docker-compose up
   ```

## Project Structure

- `cmd/cli/main.go` — CLI entrypoint
- `internal/commands/` — User interaction & command logic
- `internal/generators/` — Templates & file generation logic
- `internal/config/` — Configuration management
- `internal/helpers/` — Utility functions

## Contributing & Support

- Found a bug or have a question? [Open an issue](https://github.com/alexyvda/ava-tool/issues)
- Want to contribute? PRs are welcome!

## Credits

- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [GORM](https://gorm.io)
- [Go](https://golang.org)

## License

See [LICENSE](LICENSE) for details.

---

For more information, advanced usage, and troubleshooting, please refer to the [documentation site](https://alexy-vda.github.io/Ava-Tool/).
