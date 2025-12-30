include .env

run:
	go run $(MAIN)

run-dev:
	go run $(MAIN)

build:
	go build -o bin/$(APP_NAME) $(MAIN)

clean:
	rm -rf bin

# Variables
MIGRATION_DIR=internal/database/migrations

# Create a new migration. Usage: make migrate-create name=add_address
migrate-create:
	goose -dir $(MIGRATION_DIR) create $(name) sql

# Run all migrations
migrate-up:
	goose -dir $(MIGRATION_DIR) mysql "devuser:devpass@tcp(localhost:3306)/app_db" up

# Rollback one migration
migrate-down:
	goose -dir $(MIGRATION_DIR) mysql "devuser:devpass@tcp(localhost:3306)/app_db" down

# Add this to your Makefile (Remember to use a TAB!)
migrate-status:
	goose -dir $(MIGRATION_DIR) mysql "devuser:devpass@tcp(localhost:3306)/app_db" status