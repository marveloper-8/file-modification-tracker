package tests

import (
	"testing"
	"github.com/marveloper-8/file_modification_tracker/internal/service"
)

func TestWorkerThread(t *testing.T) {
	go service.StartWorkerThread()
}