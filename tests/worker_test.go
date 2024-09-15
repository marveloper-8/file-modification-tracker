package tests

import (
	"file-modification-tracker/daemon"
	"testing"
)

func TestAddCommand(t *testing.T) {
	daemon.AddCommand("echo test")
	select {
	case cmd := <-daemon.CommandQueue:
		if cmd != "echo test" {
			t.Fatalf("Command should be echo test, got %s", cmd)
		}
	default:
		t.Fatal("Command should be added to the queue")
	}
}