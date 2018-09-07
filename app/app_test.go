package app

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrus(t *testing.T) {
	logger := logrus.New()
	logger.Info("Hello")
}
