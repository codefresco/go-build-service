include .env
export

TOOLS_ROOT:=$(CURDIR)/tools
TOOLS_BIN:=$(TOOLS_ROOT)/bin
PATH:=$(PATH):$(TOOLS_BIN)

MIGRATION_URL=$(POSTGRES_URL)?sslmode=disable

check-migrate:
	@which migrate > /dev/null || make build-tools

check-lint:
	@which revive > /dev/null || make build-tools

.PHONY: migrate-create
migrate-create: check-migrate
ifdef NAME
	migrate create -ext sql -dir migrations $(NAME)
else
	migrate create -ext sql -dir migrations 'placeholder'
endif

.PHONY: migrate-status
migrate-status: check-migrate
	migrate -database $(MIGRATION_URL) -path migrations version

.PHONY: migrate-up
migrate-up: check-migrate
ifdef MIGRATION
	migrate -database $(MIGRATION_URL) -path migrations up $(VERSION)
else
	migrate -database $(MIGRATION_URL) -path migrations up
endif

.PHONY: migrate-down
migrate-down: check-migrate
ifdef MIGRATION
	migrate -database $(MIGRATION_URL) -path migrations down $(VERSION)
else
	migrate -database $(MIGRATION_URL) -path migrations down
endif

.PHONY: lint
lint: check-lint
	revive -config $(TOOLS_ROOT)/lint.toml -formatter friendly ./...

.PHONY: build-tools
build-tools:
	mkdir -p $(TOOLS_BIN) && \
	go build -tags 'postgres' -o $(TOOLS_BIN)/migrate github.com/golang-migrate/migrate/v4/cmd/migrate && \
	go build -o $(TOOLS_BIN)/revive github.com/mgechev/revive

.PHONY: build
build:
	go build -o bin/ main.go

.PHONY: run
run: build
	bin/main

.PHONY: dev-up
dev-up:
	docker compose -f ./devenv/docker-compose.yaml up -d

.PHONY: dev-down
dev-down:
	docker compose -f ./devenv/docker-compose.yaml down

.PHONY: dev-clear
dev-clear:
	docker compose -f ./devenv/docker-compose.yaml down -v --remove-orphans --rmi local
