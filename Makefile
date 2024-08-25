# Simple Makefile for a Go project

# Build the application

postgres:
	docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=zarokewinda -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres_db createdb --username=zarokewinda --owner=zarokewinda dashboard

dropdb:
	docker exec -it postgres_db dropdb -U zarokewinda dashboard

migrateup:
	migrate -path db/migration -database "postgresql://zarokewinda:password@localhost:5432/dashboard?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://zarokewinda:password@localhost:5432/dashboard?sslmode=disable" -verbose down

all: build

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go


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
	@go test ./... -v


# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v


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

.PHONY: postgres createdb dropdb all build run test clean migrateup migratedown
