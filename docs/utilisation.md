# Usage

!!! warning "Requirements"
    - Docker must be installed to use the generated `docker-compose.yaml` files.
    - Go 1.22+ must be installed on your machine to compile the generated services.

## Run the tool

```bash
ava create
```

### Interactive Prompts

When you run `ava create`, you will be asked the following questions:

1.  **Service Name**: The name of your new Go service.
2.  **Include PostgreSQL Database?**: (y/n) Adds Gorm, migration scripts, and Docker Compose DB setup.
3.  **Include Prometheus Metrics?**: (y/n) Adds a `/metrics` endpoint for monitoring.
4.  **Include Swagger Documentation?**: (y/n) Pre-configures Swagger. *Note: You must run `swag init` manually afterwards.*
5.  **Service Port**: The port your service will listen on (default: 8080).

_You can rerun the generation as many times as needed, each service will be created in a separate folder._

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
