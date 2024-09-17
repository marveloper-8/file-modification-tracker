package tests

import (
	"errors"
	"file-modification-tracker/internal/core"
)

// Mock ConfigPort for testing.
type MockConfig struct {
	Directory string
	CheckFreq int
}

func (m *MockConfig) GetDirectory() string {
	return m.Directory
}

func (m *MockConfig) GetCheckFrequency() int {
	return m.CheckFreq
}

// Mock LoggerPort for testing.
type MockLogger struct {
	LoggedErrors []error
	LoggedInfo   []string
	FileStats    []string
}

func (m *MockLogger) LogError(err error) {
	m.LoggedErrors = append(m.LoggedErrors, err)
}

func (m *MockLogger) LogInfo(msg string) {
	m.LoggedInfo = append(m.LoggedInfo, msg)
}

func (m *MockLogger) LogFileStats(stats string) {
	m.FileStats = append(m.FileStats, stats)
}

// Mock FileCheckerPort for testing.
type MockFileChecker struct {
	ShouldError bool
}

// GetFileModifications implements core.OsqueryAdapter.
func (m *MockFileChecker) GetFileModifications(directory string) ([]core.FileModification, error) {
	panic("unimplemented")
}

func (m *MockFileChecker) CheckModifications(directory string) (string, error) {
	if m.ShouldError {
		return "", errors.New("error fetching file modifications")
	}
	return "mocked file stats", nil
}

// Mock CommandQueuePort for testing.
type MockCommandQueue struct {
	Commands []string
}

// ReceiveCommand implements core.CommandQueueAdapter.
func (m *MockCommandQueue) ReceiveCommand() <-chan string {
	panic("unimplemented")
}

func (m *MockCommandQueue) ReceiveCommands() <-chan string {
	ch := make(chan string)
	go func() {
		for _, cmd := range m.Commands {
			ch <- cmd
		}
		close(ch)
	}()
	return ch
}
