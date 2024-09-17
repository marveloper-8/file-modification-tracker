package core

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
	CheckModifications(directory string) (string, error)
}

// CommandQueuePort defines the interface for command queue management.
type CommandQueuePort interface {
	ReceiveCommands() <-chan string
}
