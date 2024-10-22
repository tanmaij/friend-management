.PHONY: setup teardown run run-build db db-migrate db-migrate-down gen-migration-file gen-mocks gen-orm

# Setup and Teardown
setup: db db-migrate

teardown:
	@docker-compose down --volumes
	@docker-compose rm --force --stop -v

# Application
run:
	@docker-compose run --rm --service-ports app go run -mod=vendor cmd/main.go

run-build:
	@docker-compose up build

# Database
db:
	@docker-compose up -d postgres

db-migrate:
	@docker-compose run --rm app migrate -path data/migrations -database postgres://friendmgt:friendmgt@postgres:5432/friend-management?sslmode=disable up

db-migrate-down:
	@docker-compose run --rm app migrate -path data/migrations -database postgres://friendmgt:friendmgt@postgres:5432/friend-management?sslmode=disable down

# Generation tools
gen-migration-file:
	@docker-compose run --rm app migrate create -ext sql -dir data/migrations -seq new_migration_file

gen-mocks:
	@docker-compose run --rm app mockery --with-expecter=true --dir ./internal/controller --all --inpackage
	@docker-compose run --rm app mockery --with-expecter=true --dir ./internal/repository --all --inpackage

gen-model:
	@docker-compose run --rm app sqlboiler psql -c sqlboiler.yaml
