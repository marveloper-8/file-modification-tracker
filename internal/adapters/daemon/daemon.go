package daemon

import (
	"file-modification-tracker/internal/adapters/logs"
	"file-modification-tracker/internal/adapters/osquery"
	"fmt"
	"time"
)

type CommandQueueAdapter struct{}

var Logs []string

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

func Run() {
	go workerThread()
	go timerThread()
}

func workerThread() {
	for cmd := range commandQueue {
		executeCommand(cmd)
	}
}

func timerThread() {
	for {
		time.Sleep(time.Duration(5) * time.Second) // Adjust frequency for testing
		files, err := osquery.NewOsqueryAdapter().GetFileModifications("/some/directory")
		if err != nil {
			logs.NewLoggerAdapter().LogError(err)
			continue // Skip the rest of the loop if an error occurs
		}

		// Log file modification stats
		logMsg := fmt.Sprintf("File modifications: %v", files)
		Logs = append(Logs, logMsg) // Capture log for testing
		logs.NewLoggerAdapter().LogFileStats(logMsg)
	}
}
