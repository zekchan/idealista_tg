.PHONY: build
build:
	go build -o bin/bot cmd/bot/main.go

.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: dev
dev:
	air

.PHONY: test
test:
	go test ./... 