package daemon

import (
	"file-modification-tracker/internal/adapters/config"
	"file-modification-tracker/internal/adapters/logs"
	"file-modification-tracker/internal/adapters/osquery"
	"time"
)

type CommandQueueAdapter struct{}

// ReceiveCommand implements core.CommandQueueAdapter.
func (c *CommandQueueAdapter) ReceiveCommand() <-chan string {
	panic("unimplemented")
}

func (c *CommandQueueAdapter) ReceiveCommands() <-chan string {
	return commandQueue
}

func NewCommandQueueAdapter() *CommandQueueAdapter {
	commandQueue = make(chan string, 10)
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
