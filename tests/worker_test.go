package tests

import (
	"file-modification-tracker/internal/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkerThread_ProcessesCommands(t *testing.T) {
	config := &MockConfig{}
	logger := &MockLogger{}
	fileChecker := &MockFileChecker{}
	commandQueue := &MockCommandQueue{
		Commands: []string{"cmd1", "cmd2"},
	}

	service := core.NewService(config, logger, fileChecker, commandQueue)
	go service.workerThread()

	// Allow time for commands to process
	time.Sleep(1 * time.Second)

	assert.Equal(t, 2, len(logger.LoggedInfo))
	assert.Contains(t, logger.LoggedInfo[0], "Executing command: cmd1")
	assert.Contains(t, logger.LoggedInfo[1], "Executing command: cmd2")
}
