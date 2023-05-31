dev:
	docker-compose up -d --build

dev-down:
	docker-compose down
run:
	go run main.go

.PHONY: dev dev-down run