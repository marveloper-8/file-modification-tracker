package tests

import (
	"file-modification-tracker/internal/adapters/config"
	"file-modification-tracker/internal/adapters/logs"
	"file-modification-tracker/internal/adapters/osquery"
	"file-modification-tracker/internal/core"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegration_ServiceStartsSuccessfully(t *testing.T) {
	// Initialize adapters
	configAdapter := config.NewConfigAdapter()
	loggerAdapter := &logs.MockLogger{}             // Use the mock logger here
	osqueryAdapter := &osquery.MockOsqueryAdapter{} // Mock the osquery
	commandQueueAdapter := &MockCommandQueue{Commands: []string{"cmd1"}}

	// Inject dependencies into the core service
	service := core.NewService(configAdapter, loggerAdapter, osqueryAdapter, commandQueueAdapter)

	// Start service
	go service.StartService()

	// Let it run for a short time
	time.Sleep(2 * time.Second)

	// Assert that the command was executed and logs captured
	assert.Equal(t, 1, len(loggerAdapter.LoggedInfo))
	assert.Contains(t, loggerAdapter.LoggedInfo[0], "Executing command: cmd1")
}
