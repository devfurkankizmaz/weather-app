dev:
	docker-compose up -d --build

dev-down:
	docker-compose down
	
run:
	go run main.go

mock:
	mockery --name=WeatherService --filename=weather_mock.go --inpackage
