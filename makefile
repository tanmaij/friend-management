.PHONY: setup down run run-prod db db-migrate db-migrate-drop gen-migration-file gen-mocks gen-orm

setup: db db-migrate

down:
	@docker-compose rm --force --stop -v

run:
	@docker-compose run --rm --service-ports app go run -mod=vendor cmd/main.go

run-build:
	@docker-compose up build

db:
	@docker-compose up -d postgres

db-migrate:
	@docker-compose run --rm app migrate -path data/migrations -database postgres://friendmgt:friendmgt@postgres:5432/friend-management?sslmode=disable up

db-migrate-drop:
	@docker-compose run --rm app migrate -path data/migrations -database postgres://friendmgt:friendmgt@postgres:5432/friend-management?sslmode=disable drop -f

gen-migration-file:
	@docker-compose run --rm app migrate create -ext sql -dir data/migrations -seq new_migration_file

gen-mocks:
	@docker-compose run --rm app mockery --with-expecter=true --dir ./internal/controller --all --inpackage
	@docker-compose run --rm app mockery --with-expecter=true --dir ./internal/repository --all --inpackage

gen-orm:
	@docker-compose run --rm app sqlboiler psql -c sqlboiler.yaml
