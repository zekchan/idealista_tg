.PHONY: build
build:
	go build -o bin/bot cmd/bot/main.go

.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: rerun-scrape
rerun-scrape:
	go run cmd/bot/main.go rerun-scrape

.PHONY: dev
dev:
	air

.PHONY: test
test:
	go test ./... 