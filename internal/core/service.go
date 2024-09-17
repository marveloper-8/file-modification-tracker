package core

import (
	"time"
)

type Service struct {
	configAdapter  ConfigAdapter
	loggerAdapter  LoggerAdapter
	osqueryAdapter OsqueryAdapter
	commandQueue   CommandQueueAdapter
}

func NewService(
	configAdapter ConfigAdapter,
	loggerAdapter LoggerAdapter,
	osqueryAdapter OsqueryAdapter,
	commandQueue CommandQueueAdapter,
) *Service {
	return &Service{
		configAdapter:  configAdapter,
		loggerAdapter:  loggerAdapter,
		osqueryAdapter: osqueryAdapter,
		commandQueue:   commandQueue,
	}
}

func (s *Service) StartService() {
	go s.runWorkerThread()
	go s.runTimerThread()
}

func (s *Service) runWorkerThread() {
	for cmd := range s.commandQueue.ReceiveCommand() {
		s.executeCommand(cmd)
	}
}

func (s *Service) runTimerThread() {
	for {
		// Config for check frequency
		checkFreq := s.configAdapter.GetCheckFrequency()
		directory := s.configAdapter.GetDirectory()

		// Wait for the check frequency interval
		time.Sleep(time.Duration(checkFreq) * time.Second)

		// Get file modifications using the osquery adapter
		files, err := s.osqueryAdapter.GetFileModifications(directory)
		if err != nil {
			s.loggerAdapter.LogError(err)
			continue
		}

		// Log file stats
		s.loggerAdapter.LogFileStats(files)
	}
}

func (s *Service) executeCommand(cmd string) {
	// Command execution logic goes here
	s.loggerAdapter.LogFileStats("Executed: " + cmd)
}

func (s *Service) RunWorker() {
    s.runWorkerThread()
}