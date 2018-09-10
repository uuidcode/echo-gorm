package test

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLogrus(t *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
