package tests

import (
	"file-modification-tracker/internal/core"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestService_StartService(t *testing.T) {
	config := &MockConfig{
		Directory: "./",
		CheckFreq: 1, // 1 second for testing
	}
	logger := &MockLogger{}
	fileChecker := &MockFileChecker{}
	commandQueue := &MockCommandQueue{
		Commands: []string{"cmd1", "cmd2"},
	}

	service := core.NewService(config, logger, fileChecker, commandQueue)
	assert.NotNil(t, service)

	go service.StartService()

	// Let the service run for a few seconds
	time.Sleep(3 * time.Second)

	// Validate command processing
	assert.Equal(t, 2, len(logger.LoggedInfo))
	assert.Contains(t, logger.LoggedInfo[0], "Executing command: cmd1")
	assert.Contains(t, logger.LoggedInfo[1], "Executing command: cmd2")

	// Validate file checks
	assert.Greater(t, len(logger.FileStats), 0)
	assert.Contains(t, logger.FileStats[0], "mocked file stats")
}

func TestService_ErrorHandlingInFileChecker(t *testing.T) {
	config := &MockConfig{
		Directory: "./",
		CheckFreq: 1, // 1 second for testing
	}
	logger := &MockLogger{}
	fileChecker := &MockFileChecker{ShouldError: true} // Simulate error
	commandQueue := &MockCommandQueue{}

	service := core.NewService(config, logger, fileChecker, commandQueue)
	assert.NotNil(t, service)

	go service.StartService()

	// Let the service run for a few seconds
	time.Sleep(3 * time.Second)

	// Validate error handling
	assert.Greater(t, len(logger.LoggedErrors), 0)
	assert.EqualError(t, logger.LoggedErrors[0], "error fetching file modifications")
}
