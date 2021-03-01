TOOLS_BIN_DIR=./tools/bin/
TOOLS_FILE=./tools/tools.go
TOOLS_MIGRATE_FILE=./tools/migrate.go
MIGRATIONS_DIR=./migrations
LOCAL_ENV_FILE=./configs/local.env
POSTGRESQL_URL="postgres://dockerdev:dockerdev@localhost:5432/tasx_app?sslmode=disable"

env-start:
	docker-compose --env-file $(LOCAL_ENV_FILE) up -d

env-stop:
	docker-compose --env-file $(LOCAL_ENV_FILE) down

install-tools:
	go list -f '{{range .Imports}}{{.}} {{end}}' $(TOOLS_FILE) | xargs go build -o $(TOOLS_BIN_DIR)

install-tool-migrate:
	go list -f '{{range .Imports}}{{.}} {{end}}' $(TOOLS_MIGRATE_FILE) | xargs go build -tags 'postgres' -o $(TOOLS_BIN_DIR)

migrate-up:
	./tools/bin/migrate -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) up

migrate-down:
	./tools/bin/migrate -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) down

run:
	go run cmd/web/main.go
