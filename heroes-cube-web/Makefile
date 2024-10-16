# Simple Makefile for a Go project

# Build the application
export CGO_ENABLED=1
export CC=zig cc
export CXX=zig c++


all: build

build:
	@echo "Building..."
	
	
	@go build -o hc_web.exe cmd/main.go

# Run the application
run:
	@go run cmd/main.go



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
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi


DB_URL=./heroes.db
MIGRATIONS_DIR=./migrations

.PHONY: migrate-up migrate-down migrate-status create-migration

migrate-up:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_URL) up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_URL) down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_URL) status

create-migration:
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $${name} sql
