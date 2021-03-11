TOOLS_BIN_DIR = ./tools/bin/
TOOLS_MIGRATE_FILE = ./tools/migrate.go
MIGRATIONS_DIR = ./migrations
LOCAL_ENV_FILE = .env.local
SCRIPT_GET_PG_URL = ./scripts/get-pg-url.sh

COMPOSE_LOCAL = docker-compose --env-file $(LOCAL_ENV_FILE)

.PHONY: compose-build compose-up compose-down compose-logs-web install-tool-migrate migrate-up migrate-down run

compose-build:
	$(COMPOSE_LOCAL) build

compose-up:
	$(COMPOSE_LOCAL) up -d

compose-down:
	$(COMPOSE_LOCAL) down

compose-logs-web:
	@$(COMPOSE_LOCAL) logs -f web

install-tool-migrate:
	go list -f '{{range .Imports}}{{.}} {{end}}' $(TOOLS_MIGRATE_FILE) | xargs go build -tags 'postgres' -o $(TOOLS_BIN_DIR)

migrate-up:
	./tools/bin/migrate -database $(shell $(SCRIPT_GET_PG_URL) $(LOCAL_ENV_FILE)) -path $(MIGRATIONS_DIR) up

migrate-down:
	./tools/bin/migrate -database $(shell $(SCRIPT_GET_PG_URL) $(LOCAL_ENV_FILE)) -path $(MIGRATIONS_DIR) down

run:
	go run cmd/web/main.go
