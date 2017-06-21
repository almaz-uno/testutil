package testutil

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func InitLog() {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)
	}
}
