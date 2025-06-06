BINARY_NAME=main.go

ENV_FILES=.env .env.dev

default: run

build:
	go build -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)

run:
	@echo "▶️  Running $(BINARY_NAME)..."
	@for env in $(ENV_FILES); do \
		if [ -f $$env ]; then \
			export $$(cat $$env | xargs); \
		fi; \
	done; \
	go run ./cmd/$(BINARY_NAME)

fmt:
	go fmt ./...

tidy:
	go mod tidy

test:
	go test ./...

# Удалить бинарник
clean:
	rm -f $(BINARY_NAME)
