.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: runapi
runapi:
	go run cmd/api/main.go
