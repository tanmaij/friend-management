.PHONY: setup teardown run start run-build db db-migrate db-migrate-down gen-migration-file gen-mocks gen-orm test help vendor init-env

# App props
APP_VERSION := $$APP_VERSION

# Variable
DEV_SERVICE := app
DB_SERVICE := postgres
DB_URL := $$DB_URL

# Help
help:
	@echo "Available targets:"
	@echo "  setup              : Setup the environment (start db and migrate)"
	@echo "  teardown           : Stop and remove containers"
	@echo "  run                : Run the application"
	@echo "  run-build          : Build the application"
	@echo "  vendor          	: Update vendor"
	@echo "  test               : Run tests"
	@echo "  db                 : Start the database"
	@echo "  db-migrate         : Migrate database up"
	@echo "  db-migrate-down    : Migrate database down"
	@echo "  gen-migration-file : Generate new migration file"
	@echo "  gen-mocks          : Generate mocks"
	@echo "  gen-model          : Generate models"
	@echo "  init-env           : Initialize the .env file from .env.example"

# Setup and Teardown
setup: init-env db db-migrate

teardown:
	@echo "Teardown: stopping and removing containers..."
	@docker compose down --volumes
	@docker compose rm --force --stop -v

# App
start: init-env db db-migrate run

run:
	@echo "Running application..."
	@docker compose run --rm --service-ports $(DEV_SERVICE) go run -mod=vendor cmd/server/main.go

build:
	@echo "Building application..."
	@docker compose up build

clear-build:
	@echo "Clearing old builds and images..."
	@docker compose down --volumes build
	@docker compose rm --force --stop -v build
	@docker rmi -f friend-management-api:v${APP_VERSION}

test:
	@echo "Running tests..."
	@docker compose run --rm $(DEV_SERVICE) go test -mod=vendor -p 1 -v ./...

vendor: 
	@echo "Updating vendor..."
	@docker compose run --rm $(DEV_SERVICE) go mod tidy && go mod vendor

# Database
db:
	@echo "Starting database..."
	@docker compose up -d $(DB_SERVICE)

db-migrate:
	@echo "Migrating database up..."
	@docker compose run --rm $(DEV_SERVICE) sh -c 'migrate -path data/migrations -database $(DB_URL) up || { echo "Migration failed"; exit 1; }'

db-migrate-down:
	@echo "Migrating database down..."
	@docker compose run --rm $(DEV_SERVICE) sh -c 'migrate -path data/migrations -database $(DB_URL) down || { echo "Migration down failed"; exit 1; }'

# Generation tool
gen-migration-file:
	@echo "Generating new migration file..."
	@docker compose run --rm $(DEV_SERVICE) migrate create -ext sql -dir data/migrations -seq new_migration_file

gen-mocks:
	@echo "Generating mocks..."
	@docker compose run --rm $(DEV_SERVICE) mockery --with-expecter=true --dir ./internal/controller --all --inpackage
	@docker compose run --rm $(DEV_SERVICE) mockery --with-expecter=true --dir ./internal/repository --all --inpackage

gen-model:
	@echo "Generating models..."
	@docker compose run --rm $(DEV_SERVICE) sqlboiler psql -c sqlboiler.yaml

# environment
init-env:
	@docker compose --env-file .env.example run --rm $(DEV_SERVICE) sh -c 'if [ ! -f .env ]; then cp .env.example .env && echo ".env created from .env.example"; else echo ".env already exists"; fi'
