package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware returns the Echo middleware.
func LoggerMiddleware(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Perform logging
			logger.Info("HTTP", c.Request().Method, c.Request().URL.Path, c.Request().URL.Query(), c.Response().Status)

			// Continue with the processing
			return next(c)
		}
	}
}
