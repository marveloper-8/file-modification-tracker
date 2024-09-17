package daemon

import (
	"file-modification-tracker/internal/adapters/config"
	"file-modification-tracker/internal/adapters/logs"
	"file-modification-tracker/internal/adapters/osquery"
	"time"
)

type CommandQueueAdapter struct{}

// ReceiveCommands implements core.CommandQueuePort.
func (c *CommandQueueAdapter) ReceiveCommands() <-chan string {
	panic("unimplemented")
}

func NewCommandQueueAdapter() *CommandQueueAdapter {
	commandQueue = make(chan string, 10) // Buffered channel
	return &CommandQueueAdapter{}
}

func (c *CommandQueueAdapter) Run() error {
	go workerThread()
	go timerThread()
	return nil
}

func workerThread() {
	for cmd := range commandQueue {
		executeCommand(cmd)
	}
}

func timerThread() {
	for {
		time.Sleep(time.Duration(config.Config.CheckFreq) * time.Second)
		files, err := osquery.NewOsqueryAdapter().GetFileModifications(config.Config.Directory)
		if err != nil {
			logs.NewLoggerAdapter().LogError(err)
		}
		logs.NewLoggerAdapter().LogFileStats(files)
	}
}
