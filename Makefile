.PHONY: up down run test

up:
	docker compose up -d

down:
	docker compose down

run:
	set -a; source configs/.env; set +a; go run ./cmd/main.go

test:
	go test -race -shuffle=on ./...