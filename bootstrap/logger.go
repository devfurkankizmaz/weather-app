package bootstrap

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

// Logger represents the configured Logrus logger.
var Logger *logrus.Logger

// setupLogger performs the Logrus configuration.
func setupLogger() {
	// Create a new Logrus logger
	logger := logrus.New()

	// Set the output to standard output and enable coloring
	logger.SetOutput(colorable.NewColorableStdout())

	// Use a custom log format with timestamp and colored level
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05", // Customize the timestamp format as desired
		FullTimestamp:   true,
	})

	// Set the log level (default is Debug)
	logger.SetLevel(logrus.DebugLevel)

	// Set as the global logger
	Logger = logger
}

// InitLogger initializes and returns the logger.
func InitLogger() *logrus.Logger {
	setupLogger()
	return Logger
}
