# Build stage
FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./cmd/main.go

# Deploy stage
FROM alpine:3.21 AS deploy
WORKDIR /app
COPY --from=build /app/app .
USER 1001
CMD ["./app"]