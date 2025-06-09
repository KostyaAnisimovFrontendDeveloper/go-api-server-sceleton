BINARY_NAME=main.go
ENV_FILES=.env .env.dev
DB_URL=postgres://postgres:postgres@localhost:5432/apigobox?sslmode=disable
DB_URL_TEST=postgres://postgres:postgres@localhost:5432/apigobox-test?sslmode=disable
MIGRATIONS_DIR=./db/migrations
TESTS_DIR=./tests
SERVER_PORT=8081
default: run

build:
	go build -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)
.PHONY: build

run:
	@echo "▶️  Running $(BINARY_NAME)..."
	APP_ENV=dev go run $(BINARY_NAME)
.PHONY: run

stop:
	@echo "Остановка сервера"
	@kill -SIGINT $(shell lsof -t -i:$(SERVER_PORT))
.PHONY: stop
fmt:
	go fmt ./...
.PHONY: fmt

tidy:
	go mod tidy
.PHONY: tidy

test:
	APP_ENV=test go test $(TESTS_DIR)
.PHONY: test

testv:
	APP_ENV=test go test $(TESTS_DIR) -v
.PHONY: testv

clean:
	rm -f $(BINARY_NAME)
.PHONY: clean

mc:
	goose create -dir $(MIGRATIONS_DIR) $(NAME) sql

mu:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up
.PHONY: mu

md:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down
.PHONY: md

mut:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_TEST)" up
.PHONY: mut

mdt:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL_TEST)" down
.PHONY: mdt
