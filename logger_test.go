package go_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logrus := logrus.New()

	logrus.Println("Hello Logger")
}
