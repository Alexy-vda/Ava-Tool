# Ava

!!! info "Context and Goals"
    Ava is a modern and fast Go service generator, designed to speed up the creation of robust microservices, with or without a database, and ready-to-use DevOps integration.

---

## Installation

### Download the CLI

- Go to the [GitHub Releases page](https://github.com/alexy-vda/Ava-Tool/releases)
- Download the binary for your OS and architecture
- Rename it to `ava` (or `ava.exe` on Windows)
- Move it to a directory in your `PATH`

```bash
# Example for Linux/macOS
chmod +x ava
sudo mv ava /usr/local/bin/
```

### (Optional) Install via Homebrew

_Coming soon!_

---

## Quick Start

```bash
ava create
```

For more usage examples, see the [Usage](utilisation.md) section.

# Quick Overview

Ava automates:
- The Go project structure (cmd, internal, helpers, etc.)
- Docker/Docker Compose integration
- CI/CD configuration (GitLab)
- Adding a PostgreSQL database (optional)

For more details, see the [Architecture](architecture.md) page.

!!! warning "Warning"
    The generated services are ready-to-use bases, but must be adapted to your business needs: security, validation, error handling, etc. Never deploy to production without a code review!

- [Usage](utilisation.md)
- [Architecture](architecture.md)
- [FAQ](faq.md)
- [Contact](contact.md)
