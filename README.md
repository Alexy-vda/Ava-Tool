# Ava

Ava is a CLI utility designed to accelerate the creation of Go services. It generates a ready-to-use service structure, including a Gin-based router, optional Gorm ORM setup (with PostgreSQL), Docker & Docker Compose integration, and a GitLab CI configuration.

[![Documentation](https://img.shields.io/badge/docs-online-blue)](https://alexy-vda.github.io/Ava-Tool/)
[![GitHub Release](https://img.shields.io/github/v/release/alexy-vda/Ava-Tool?color=blue)](https://github.com/alexy-vda/Ava-Tool/releases)
[![Go Report](https://goreportcard.com/badge/github.com/alexy-vda/Ava-Tool)](https://goreportcard.com/report/github.com/alexy-vda/Ava-Tool)
---

## 🚀 Installation

### Download the latest release

Go to the [Releases page](https://github.com/alexy-vda/Ava-Tool/releases) and download the binary for your platform:

- **Linux**: `ava-linux-amd64` or `ava-linux-arm64`
- **macOS**: `ava-darwin-amd64` or `ava-darwin-arm64`
- **Windows**: `ava-windows-amd64.exe` or `ava-windows-arm64.exe`

Rename the file to `ava` (or `ava.exe` on Windows) and move it to a directory in your `PATH` (e.g., `/usr/local/bin` on Linux/macOS).

```bash
# Example for Linux/macOS
chmod +x ava
sudo mv ava /usr/local/bin/
```

### (Optional) Install via Homebrew

_Coming soon!_

---

## 🛠️ Usage

```bash
ava create
```

For more details, see the [documentation](https://alexy-vda.github.io/Ava-Tool/).

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
