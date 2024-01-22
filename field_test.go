package go_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "rafli").Info("hello world")

	logger.WithField("username", "muhammad").
		WithField("name", "muhammad rafli").
		Info("hello world")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "eko",
		"name":     "eko kurniawan",
	}).Info("hello world")
}
