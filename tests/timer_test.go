package tests

import (
	"file-modification-tracker/daemon"
	"testing"
	"time"
)

func TestTimerThread(t *testing.T) {
	daemon.Run()

	time.Sleep(10 * time.Second)
	if len(daemon.Logs) != 1 {
		t.Errorf("Logs should be added to the queue")
	}
}
