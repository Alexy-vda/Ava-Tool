# Usage

!!! warning "Requirements"
    - Docker must be installed to use the generated `docker-compose.yaml` files.
    - Go 1.22+ must be installed on your machine to compile the generated services.

## Run the tool

```bash
ava-tool
```

Answer the interactive questions to generate a new Go service.

!!! info "Tip"
    You can rerun the generation as many times as needed, each service will be created in a separate folder.

## Start a generated service

```bash
cd <service-name>
docker-compose up
```

## Default endpoints

- `/health`: Checks the service availability.

## Environment variables

- `PORT`: Service listening port (default: 8080)
- `GIN_MODE`: Gin mode (`debug`, `release`, `test`)
