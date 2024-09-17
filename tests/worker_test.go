package tests

import (
	"file-modification-tracker/internal/core"
	"file-modification-tracker/internal/adapters/logs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWorkerThread_ProcessesCommands(t *testing.T) {
	config := &MockConfig{}
	logger := &logs.MockLogger{}
	fileChecker := &MockFileChecker{}
	commandQueue := &MockCommandQueue{
		Commands: []string{"cmd1", "cmd2"},
	}

	service := core.NewService(config, logger, fileChecker, commandQueue)
	go service.RunWorker()

	// Allow time for commands to process
	time.Sleep(1 * time.Second)

	assert.Equal(t, 2, len(logger.LoggedInfo))
	assert.Contains(t, logger.LoggedInfo[0], "Executing command: cmd1")
	assert.Contains(t, logger.LoggedInfo[1], "Executing command: cmd2")
}
