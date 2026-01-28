# Usage

!!! warning "Requirements"
    - Docker must be installed to use the generated `docker-compose.yaml` files.
    - Go 1.22+ must be installed on your machine to compile the generated services.

## Run the tool

```bash
ava create
```

### Interactive Prompts

When you run `ava create` without flags, you will be asked the following questions:

1.  **Service Name**: The name of your new Go service.
2.  **Include PostgreSQL Database?**: (y/n) Adds Gorm, migration scripts, and Docker Compose DB setup.
3.  **Include Prometheus Metrics?**: (y/n) Adds a `/metrics` endpoint for monitoring.
4.  **Include Sentry for error tracking?**: (y/n) Adds Sentry integration and middleware.
5.  **Include Swagger Documentation?**: (y/n) Pre-configures Swagger. *Note: You must run `swag init` manually afterwards.*
6.  **Service Port**: The port your service will listen on (default: 8080).

### Non-Interactive Mode (Flags)

For automation, CI/CD, or AI agents, you can use flags to skip the prompts.

```bash
ava create --name <name> [options]
```

**Available Flags:**

- `--name, -n <string>`: Name of the service (Required for non-interactive mode).
- `--port, -p <string>`: Port to listen on (Default: "8080").
- `--with-db`: Include PostgreSQL database.
- `--with-prometheus`: Include Prometheus metrics.
- `--with-sentry`: Include Sentry error tracking.
- `--with-swagger`: Include Swagger documentation.

**Example:**
```bash
# Create a service named "auth-service" with DB and Prometheus on port 3000
ava create --name auth-service --with-db --with-prometheus --port 3000
```

_You can rerun the generation as many times as needed, each service will be created in a separate folder._

### AI / Claude Code Integration

Each generated service includes a `.claude/CLAUDE.md` file. This file provides instructions for AI agents like Claude Code on how to build, run, test, and understand the project, making it immediately "AI-friendly".

## Start a generated service

```bash
cd <service-name>
docker-compose up --build
```

## Default endpoints

- `/health`: Checks the service availability.
- `/metrics`: (Optional) Prometheus metrics.
- `/swagger/*any`: (Optional) Swagger API documentation.

## Environment variables

- `PORT`: Service listening port (default: 8080)
- `GIN_MODE`: Gin mode (`debug`, `release`, `test`)
- `SENTRY_DSN`: (Optional) Sentry Data Source Name for error tracking.
