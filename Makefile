# Simple Makefile for a Go project
MIGRATION_DIR = migrations
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Use environment variables for DB_DRIVER and DB_DSN
DB_DRIVER = postgres
DB_DSN = $(GOOSE_DBSTRING)
GOOSE_CMD = goose -dir $(MIGRATION_DIR)
# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

source-goose:
	@bash -c 'source env/goose.env'

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi
	
db-status:
	@DB_DRIVER=$(DB_DRIVER) DB_DSN="$(DB_DSN)" $(GOOSE_CMD) status
migrate:
	@DB_DRIVER=$(DB_DRIVER) DB_DSN="$(DB_DSN)" $(GOOSE_CMD) up

# Rollback the last migration
rollback:
	./goose-custom down

.PHONY: all build run test clean
