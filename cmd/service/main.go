package main

import (
	"file-modification-tracker/internal/adapters/config"
	"file-modification-tracker/internal/adapters/daemon"
	"file-modification-tracker/internal/adapters/http"
	"file-modification-tracker/internal/adapters/logs"
	"file-modification-tracker/internal/adapters/osquery"
	"file-modification-tracker/internal/adapters/ui"
	"file-modification-tracker/internal/core"
	"os"
)

func main() {

	configAdapter := config.NewConfigAdapter()
	loggerAdapter := logs.NewLoggerAdapter()
	osqueryAdapter := osquery.NewOsqueryAdapter()
	commandQueueAdapter := daemon.NewCommandQueueAdapter()

	service := core.NewService(configAdapter, loggerAdapter, osqueryAdapter, commandQueueAdapter)

	service.StartService()

	go func() {
		err := http.NewHTTPServer(service).Run()
		if err != nil {
			loggerAdapter.LogError(err)
			os.Exit(1)
		}
	}()

	go func() {
		err := ui.NewUI().Run()
		if err != nil {
			loggerAdapter.LogError(err)
			os.Exit(1)
		}
	}()

	select {}
}
