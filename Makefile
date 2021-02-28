TOOLS_BIN_DIR=./tools/bin/
TOOLS_FILE=./tools/tools.go
TOOLS_MIGRATE_FILE=./tools/migrate.go

env-start:
	docker-compose up -d

install-tools:
	go list -f '{{range .Imports}}{{.}} {{end}}' $(TOOLS_FILE) | xargs go build -o $(TOOLS_BIN_DIR)

install-tool-migrate:
	go list -f '{{range .Imports}}{{.}} {{end}}' $(TOOLS_MIGRATE_FILE) | xargs go build -tags 'postgres' -o $(TOOLS_BIN_DIR)
