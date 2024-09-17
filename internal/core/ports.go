package core

import "time"

type ConfigPort interface {
	GetDirectory() string
	GetCheckFrequency() int
}

type LoggerPort interface {
	LogError(err error)
	LogInfo(msg string)
	LogFileStats(stats string)
}

type FileCheckerPort interface {
	GetFileModifications(directory string) (string, error)
}

type CommandQueuePort interface {
	ReceiveCommands() <-chan string
	AddCommand(cmd string)
}

type ConfigAdapter interface {
	GetCheckFrequency() int
	GetDirectory() string
}

type LoggerAdapter interface {
	LogError(err error)
	LogFileStats(stats interface{})
}

type OsqueryAdapter interface {
	GetFileModifications(directory string) ([]FileModification, error)
}

type CommandQueueAdapter interface {
	ReceiveCommand() <-chan string
}

type FileModification struct {
	Filename     string
	LastModified time.Time
}
