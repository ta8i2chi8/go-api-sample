.PHONY: up down run test

up:
	docker compose up -d

down:
	docker compose down

run:
	go run ./cmd/main.go

test:
	go test -race -shuffle=on ./...