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

	configAdapter := config.NewConfigAdapter()
	loggerAdapter := &logs.MockLogger{}
	osqueryAdapter := &osquery.MockOsqueryAdapter{}
	commandQueueAdapter := &MockCommandQueue{Commands: []string{"cmd1"}}

	service := core.NewService(configAdapter, loggerAdapter, osqueryAdapter, commandQueueAdapter)

	go service.StartService()

	time.Sleep(2 * time.Second)

	assert.Equal(t, 1, len(loggerAdapter.LoggedInfo))
	assert.Contains(t, loggerAdapter.LoggedInfo[0], "Executing command: cmd1")
}
