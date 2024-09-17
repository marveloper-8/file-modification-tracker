package core

import "time"

// ConfigPort defines the interface for configuration handling.
type ConfigPort interface {
	GetDirectory() string
	GetCheckFrequency() int
}

// LoggerPort defines the interface for logging.
type LoggerPort interface {
	LogError(err error)
	LogInfo(msg string)
	LogFileStats(stats string)
}

// FileCheckerPort defines the interface for checking file modifications.
type FileCheckerPort interface {
	GetFileModifications(directory string) (string, error)
}

// CommandQueuePort defines the interface for command queue management.
type CommandQueuePort interface {
	ReceiveCommands() <-chan string
	AddCommand(cmd string)
}

// ConfigAdapter interface
type ConfigAdapter interface {
    GetCheckFrequency() int
    GetDirectory() string
}

// LoggerAdapter interface
type LoggerAdapter interface {
    LogError(err error)
    LogFileStats(stats interface{})
}

// OsqueryAdapter interface
type OsqueryAdapter interface {
    GetFileModifications(directory string) ([]FileModification, error)
}

// CommandQueueAdapter interface
type CommandQueueAdapter interface {
    ReceiveCommand() <-chan string
}

// FileModification struct (used in OsqueryAdapter)
type FileModification struct {
    Filename    string
    LastModified time.Time
    // Add other relevant fields
}