package main

import (
	"file-modification-tracker/internal/adapters/config"
	"file-modification-tracker/internal/adapters/daemon"
	"file-modification-tracker/internal/adapters/http"
	"file-modification-tracker/internal/adapters/logs"
	"file-modification-tracker/internal/adapters/osquery"
	"file-modification-tracker/internal/adapters/ui"
	"file-modification-tracker/internal/core"
	"log"
	"os"
)

func main() {
	// Initialize adapters
	configAdapter := config.NewConfigAdapter()
	loggerAdapter := logs.NewLoggerAdapter()
	osqueryAdapter := osquery.NewOsqueryAdapter()
	commandQueueAdapter := daemon.NewCommandQueueAdapter()

	// Inject dependencies into the core service
	service := core.NewService(configAdapter, loggerAdapter, osqueryAdapter, commandQueueAdapter)

	// Start service
	service.StartService()

	// Initialize HTTP server
	go func() {
		err := http.NewHTTPServer(service).Run()
		if err != nil {
			loggerAdapter.LogError(err)
			os.Exit(1)
		}
	}()

	// Initialize UI
	go func() {
		err := ui.NewUI().Run()
		if err != nil {
			loggerAdapter.LogError(err)
			os.Exit(1)
		}
	}()

	select {} // Keep main running
}
