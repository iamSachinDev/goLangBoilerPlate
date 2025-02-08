package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)        // Log to console
	Logger.SetLevel(logrus.DebugLevel) // Set default log level

	// Set log format (JSON or Text)
	Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: false,
	})
}
