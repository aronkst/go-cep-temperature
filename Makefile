dev-start:
	docker compose up -d

dev-stop:
	docker compose stop

dev-down:
	docker compose down

dev-run:
	docker compose exec dev-go-cep-temperature go run cmd/server/main.go

dev-run-tests:
	docker compose exec dev-go-cep-temperature go test ./... -v

prod-start:
	docker compose -f docker-compose.prod.yml up -d

prod-stop:
	docker compose -f docker-compose.prod.yml stop

prod-down:
	docker compose -f docker-compose.prod.yml down
