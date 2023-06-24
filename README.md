# Weather App

A simple Weather App built with Go that utilizes an external API to fetch current weather data based on a given city.

## Features

- Retrieves current weather data from the [Weather API](https://www.weatherapi.com/)
- Supports logging with Logrus [Logrus](https://github.com/sirupsen/logrus)
- Caching with Redis [go-redis](https://github.com/redis/go-redis)
- Environment configuration with Viper [Viper](https://github.com/spf13/viper)
- Documentation with Swagger [Swaggo](https://github.com/swaggo/swag)
- Echo Framework [Echo](https://echo.labstack.com/)

## Requirements

- Go 1.16 or higher
- Redis
- Weather API key (sign up at https://www.weatherapi.com/ to obtain a key)
- Docker

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/devfurkankizmaz/weather-app.git
   ```

2. Navigate to the project directory:

   ```bash
   cd weather-app
   ```

3. Install the dependencies:

   ```bash
   go mod download
   ```

4. Set up environment variables:

- Create a `config.yml` file in the project root directory.
- Add the following variables to the `config.yml` file:

  ```yml
  SERVER_PORT: :7070
  REDIS_ADDRESS: localhost:6379
  REDIS_PASS:
  REDIS_DB: 0
  REDIS_EXPIRY_MIN: 30
  CONTEXT_TIMEOUT_SEC: 10
  API_KEY: apikey
  API_URL: api.weatherapi.com/v1/current.json
  ```

- Replace `<your-weather-api-key>` with your actual Weather API key.
- Replace `<redis-address>`, `<redis-password>`, `<redis-database>` with your Redis server information

5. Run the application

   ```bash
   make dev
   make run
   ```

6. Open your web browser and navigate to `localhost:7070/docs/index.html` to access the Swagger documentation.

## Usage

- Open your web browser and navigate to `localhost:7070/weather`.
- Append the `city` query parameter to the URL, for example: `localhost:7070/weather?city=Istanbul` or `localhost:7070/weather?city=New%York`
- The application will fetch the current weather data for the specified city and display it

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

