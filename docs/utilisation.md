# Usage

!!! warning "Requirements"
    - Docker must be installed to use the generated `docker-compose.yaml` files.
    - Go 1.22+ must be installed on your machine to compile the generated services.

## Run the tool

```bash
ava <command>
```

### Available commands

- `ava create` : Interactively generate a new Go service
- `ava help`   : Show help and available commands

_You can rerun the generation as many times as needed, each service will be created in a separate folder._

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
