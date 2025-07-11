include local.env

LOCAL_BIN = $(CURDIR)/bin

all: install-deps migration-up build

build:
	go build -o bin/main ./cmd/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migration-up:
	$(LOCAL_BIN)/migrate --path ${MIGRATION_DIR} -database ${PG_MIGRATION_DSN} up

migration-down:
	$(LOCAL_BIN)/migrate --path ${MIGRATION_DIR} -database ${PG_MIGRATION_DSN} down 1