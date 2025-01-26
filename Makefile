include .env

DEFAULT_GOAL := help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: up
up: ## "Statrting containers ..."
	@docker compose up --build -d

.PHONY: down
down: ## "Stop and remove all docker containers."
	@docker compose down --remove-orphans

.PHONY: build
build: # "Build the app without api service in the docker-compose.yml (before building must comment the api service)."
	go build -o ./bin/${APP_NAME} cmd/server/main.go

.PHONY: run
run: build ## "Build and run the app without api service in the docker-compose.yml(before runnig it must comment the api service)."
	@env DB_USERNAME=${DB_USERNAME} DB_PASSWORD=${DB_PASSWORD} ./bin/${APP_NAME}
