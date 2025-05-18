.PHONY: up down run test

up:
	docker compose up -d

down:
	docker compose down

run:
	set -a; source .env; set +a; go run ./cmd/main.go

test:
	go test -race -shuffle=on ./...