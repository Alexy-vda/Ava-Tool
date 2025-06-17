# Ava Tool

!!! info "Context and Goals"
    Ava Tool is a modern and fast Go service generator, designed to speed up the creation of robust microservices, with or without a database, and ready-to-use DevOps integration.
    
    It is aimed at developers who want to standardize and industrialize the creation of Go services, while keeping the flexibility to add a PostgreSQL database or not.

!!! warning "Warning"
    The generated services are ready-to-use bases, but must be adapted to your business needs: security, validation, error handling, etc. Never deploy to production without a code review!

- [Usage](utilisation.md)
- [Architecture](architecture.md)
- [FAQ](faq.md)
- [Contact](contact.md)

# Quick Overview

Ava Tool automates:
- The Go project structure (cmd, internal, helpers, etc.)
- Docker/Docker Compose integration
- CI/CD configuration (GitLab)
- Adding a PostgreSQL database (optional)

For more details, see the [Architecture](architecture.md) page.
