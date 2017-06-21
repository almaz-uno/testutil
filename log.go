package testutil

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func initLog() {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)
	}
}
