package go_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	logger.SetOutput(file)

	logger.Info("Hello Logging")
	logger.Warn("Hello Warning")
	logger.Error("Hello Error")
}
