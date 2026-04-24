APP_NAME=go_peregerine
DOCKER_COMPOSE=docker-compose
APP_SERVICE=app
LINTER_SERVICE=lint

## ---------------------------
## GO COMMANDS
## ---------------------------

# Run app locally (without Docker)
run:
	go run main.go

# Format all files
fmt:
	go fmt ./...
	goimports -w .

# Run all unit tests
test:
	go test ./... -cover

# Lint using golangci-lint inside container
lint:
	docker exec -it $(APP_NAME) golangci-lint run ./...

# Start dev environment (Air + auto-reload)
dev:
	$(DOCKER_COMPOSE) up --build

# Start dev environment (Air + auto-reload)
up:
	$(DOCKER_COMPOSE) up

# Stop containers
stop:
	$(DOCKER_COMPOSE) down

# Rebuild container with no cache
rebuild:
	$(DOCKER_COMPOSE) build --no-cache

# Restart container
restart:
	$(DOCKER_COMPOSE) down && $(DOCKER_COMPOSE) up --build

# Run linter using a temporary container
docker-lint:
	$(DOCKER_COMPOSE) run --rm app golangci-lint run ./...

# Shell into running container
sh:
	docker exec -it $(APP_NAME) sh

sh-lint:
	docker compose run --rm $(LINTER_SERVICE) golangci-lint run

worker-run:
	docker exec -it $(APP_NAME) sh && go run jobs/main.job.go

## ---------------------------
## CLEANUP
## ---------------------------

# Remove Go build artifacts
clean:
	rm -rf tmp
	rm -rf bin
	rm -rf dist

# Remove Docker images + cache
docker-clean:
	docker system prune -af --volumes