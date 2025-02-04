build-dev:
	docker compose up -d --build
run:
	docker compose start api
	docker compose logs -f api
logs-api:
	docker compose logs -f api
restart-api:
	docker compose restart api
	make logs-api
create-docs:
	~/go/bin/swag init -g ./cmd/main.go -o cmd/docs
