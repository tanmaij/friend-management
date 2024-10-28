# Include environment variables
include .env

.PHONY: setup teardown run start build clear-build db db-migrate db-migrate-down gen-migration-file gen-mocks gen-orm test help vendor init-env

# App props
APP_VERSION := $(APP_VERSION)

# Variables
DEV_SERVICE := app
DB_SERVICE := postgres
DB_URL := $(DB_URL)

# Help
help:
	@echo "Available targets:"
	@echo "  setup              : Setup the environment (start db and migrate)"
	@echo "  teardown           : Stop and remove containers"
	@echo "  start              : Start the app with dependencies"
	@echo "  run                : Run the application in dev mode"
	@echo "  build              : Build the application"
	@echo "  build-run          : Build and run the application"
	@echo "  build-stop         : Stop the running application"
	@echo "  clear-build        : Clear current app version and volumes"
	@echo "  vendor             : Update vendor dependencies"
	@echo "  test               : Run tests"
	@echo "  db                 : Start the database service"
	@echo "  db-migrate         : Migrate database up"
	@echo "  db-migrate-down    : Migrate database down"
	@echo "  gen-migration-file : Generate new migration file"
	@echo "  gen-mocks          : Generate mock files"
	@echo "  gen-model          : Generate database models"
	@echo "  init-env           : Initialize the .env file from .env.example"

# Setup and Teardown
setup: init-env db db-migrate
teardown:
	@echo "Teardown: stopping and removing containers..."
	@docker compose down --volumes
	@docker compose rm --force --stop -v

# App Commands
start: init-env db db-migrate run
run:
	@echo "Running application in dev mode..."
	@docker compose run --rm --service-ports $(DEV_SERVICE) go run -mod=vendor cmd/server/main.go

build:
	@echo "Building application..."
	@docker build -t friend-management-api:${APP_VERSION} -f Dockerfile.release .

build-run:
	@echo "Running application..."
	@docker run -d \
		--name friend-management-api \
		-p ${PORT}:${PORT} \
		-e DB_URL=${DB_URL} \
		-e PORT=${PORT} \
		friend-management-api:${APP_VERSION}

build-stop:
	@echo "Stopping application..."
	@docker stop friend-management-api

clear-build:
	@echo "Clearing old builds and images..."
	@docker rm --force -v friend-management-api
	@docker rmi -f friend-management-api:${APP_VERSION}

test:
	@echo "Running tests..."
	@docker compose run --rm $(DEV_SERVICE) go test -mod=vendor -p 1 -v ./...

vendor:
	@echo "Updating vendor dependencies..."
	@docker compose run --rm $(DEV_SERVICE) go mod tidy && go mod vendor

# Database Commands
db:
	@echo "Starting database..."
	@docker compose up -d $(DB_SERVICE)

db-migrate:
	@echo "Migrating database up..."
	@docker compose run --rm $(DEV_SERVICE) sh -c 'migrate -path data/migrations -database $(DB_URL) up || { echo "Migration failed"; exit 1; }'

db-migrate-down:
	@echo "Migrating database down..."
	@docker compose run --rm $(DEV_SERVICE) sh -c 'migrate -path data/migrations -database $(DB_URL) down || { echo "Migration down failed"; exit 1; }'

# Generation Tools
gen-migration-file:
	@echo "Generating new migration file..."
	@docker compose run --rm $(DEV_SERVICE) migrate create -ext sql -dir data/migrations -seq new_migration_file

gen-mocks:
	@echo "Generating mock files..."
	@docker compose run --rm $(DEV_SERVICE) mockery --with-expecter=true --dir ./internal/controller --all --inpackage
	@docker compose run --rm $(DEV_SERVICE) mockery --with-expecter=true --dir ./internal/repository --all --inpackage

gen-model:
	@echo "Generating database models..."
	@docker compose run --rm $(DEV_SERVICE) sqlboiler psql -c sqlboiler.yaml

# Environment
init-env:
	@docker compose --env-file .env.example run --rm $(DEV_SERVICE) sh -c 'if [ ! -f .env ]; then cp .env.example .env && echo ".env created from .env.example"; else echo ".env already exists"; fi'
