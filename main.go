/*
 * @File: main.go
 * @Description: Creates HTTP server
 * @Author: Furkan Kızmaz (dev.furkan@outlook.com)
 */
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/devfurkankizmaz/weather-app/bootstrap"
	_ "github.com/devfurkankizmaz/weather-app/docs"
	"github.com/devfurkankizmaz/weather-app/handler"
	mw "github.com/devfurkankizmaz/weather-app/middleware"
	"github.com/devfurkankizmaz/weather-app/repository"
	"github.com/devfurkankizmaz/weather-app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Weather App
// @version 1.0
// @description This service is a weather web server that utilizes an external API.
// @host localhost:7070
// @BasePath /
func main() {
	// App instance
	app := bootstrap.App()
	timeout := time.Duration(app.Env.CONTEXT_TIMEOUT_SEC) * time.Second
	expiry := time.Duration(app.Env.REDIS_EXPIRY_MIN) * time.Minute

	repository := repository.NewWeatherRepository(
		app.Client,
		expiry,
	)
	service := service.NewWeatherService(
		repository,
		timeout,
	)
	handler := handler.NewWeatherHandler(
		service,
		app.Env,
	)

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(mw.LoggerMiddleware(app.Logger))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	log.Printf("Server is running on %s", app.Env.SERVER_PORT)
	e.GET("/weather", handler.Weather)
	e.GET("/", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.JSON(http.StatusOK, "OK")
	})
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// Graceful Shutdown
	go func() {
		if err := e.Start(app.Env.SERVER_PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
