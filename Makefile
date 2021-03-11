TOOLS_BIN_DIR = ./tools/bin/
TOOLS_MIGRATE_FILE = ./tools/migrate.go
MIGRATIONS_DIR = ./migrations
SCRIPT_GET_PG_URL = ./scripts/get-pg-url.sh

LOCAL_ENV_FILE = .env.local
PROD_ENV_FILE = .env.prod

COMPOSE_LOCAL = docker-compose --env-file $(LOCAL_ENV_FILE)

.PHONY: compose-build compose-up compose-down compose-logs-web
.PHONY: install-tool-migrate migrate-up migrate-down
.PHONY: prod-build prod-run prod-stop prod-migrate-up prod-migrate-down


compose-build:
	$(COMPOSE_LOCAL) build

compose-up:
	$(COMPOSE_LOCAL) up -d

compose-down:
	$(COMPOSE_LOCAL) down

compose-logs-web:
	@$(COMPOSE_LOCAL) logs -f web

compose-pgsql:
	$(COMPOSE_LOCAL) up -d pgsql

install-tool-migrate:
	go list -f '{{range .Imports}}{{.}} {{end}}' $(TOOLS_MIGRATE_FILE) | xargs go build -tags 'postgres' -o $(TOOLS_BIN_DIR)

migrate-up:
	./tools/bin/migrate -database $(shell $(SCRIPT_GET_PG_URL) $(LOCAL_ENV_FILE)) -path $(MIGRATIONS_DIR) up

migrate-down:
	./tools/bin/migrate -database $(shell $(SCRIPT_GET_PG_URL) $(LOCAL_ENV_FILE)) -path $(MIGRATIONS_DIR) down

prod-build:
	docker build -f containers/prod/Dockerfile -t tasx_app .

prod-run:
	./scripts/prod-run.sh $(PROD_ENV_FILE)

prod-stop:
	docker stop tasx

prod-migrate-up:
	./tools/bin/migrate -database $(shell $(SCRIPT_GET_PG_URL) $(PROD_ENV_FILE)) -path $(MIGRATIONS_DIR) up

prod-migrate-down:
	./tools/bin/migrate -database $(shell $(SCRIPT_GET_PG_URL) $(PROD_ENV_FILE)) -path $(MIGRATIONS_DIR) down
