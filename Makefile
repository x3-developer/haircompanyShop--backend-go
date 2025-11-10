OUTPUT_DIR = bin
BINARY = app
API_MAIN = cmd/api/main.go

.PHONY: all build clean run tidy test cover migration migrate

all: build

build:
	@echo "Building the application..."
	mkdir -p ${OUTPUT_DIR}
	go build -o ${OUTPUT_DIR}/${BINARY} ${API_MAIN}

run: build
	@echo "Running the application..."
	${OUTPUT_DIR}/${BINARY}

clean:
	@echo "Cleaning up..."
	rm -rf ${OUTPUT_DIR}

tidy:
	@echo "Tidying up dependencies..."
	go mod tidy

migration:
	@if [ -z "$(name)" ]; then \
		echo "err: migration name is required name=<migration_name>"; \
		exit 1; \
	fi
	@echo "Creating migration: $(name)..."
	migrate create -ext sql -dir migrations -seq $(name)

migrate:
	@if [ -z "$(direction)" ]; then \
		echo "err: direction is required (direction=up|down)"; \
		exit 1; \
	fi; \
	if [ "$(direction)" != "up" ] && [ "$(direction)" != "down" ]; then \
		echo "err: invalid direction value (use up or down)"; \
		exit 1; \
	fi
	@echo "Running migrations $(direction)..."
	go run migrations/auto.go $(direction)

test:
	@echo "Running unit tests..."
	go test -v ./...

db.test.up:
	docker compose -f docker-compose.test.yml up -d

db.test.down:
	docker compose -f docker-compose.test.yml down -v

migrate.test:
	migrate -path migrations -database "postgres://test:test@localhost:5433/testdb?sslmode=disable" up

test.integration: db.test.up migrate.test
	go test -v -tags=integration ./...

cover:
	@echo "Running tests with coverage..."
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out
