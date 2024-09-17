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

		checkFreq := s.configAdapter.GetCheckFrequency()
		directory := s.configAdapter.GetDirectory()

		time.Sleep(time.Duration(checkFreq) * time.Second)

		files, err := s.osqueryAdapter.GetFileModifications(directory)
		if err != nil {
			s.loggerAdapter.LogError(err)
			continue
		}

		s.loggerAdapter.LogFileStats(files)
	}
}

func (s *Service) executeCommand(cmd string) {

	s.loggerAdapter.LogFileStats("Executed: " + cmd)
}

func (s *Service) RunWorker() {
	s.runWorkerThread()
}
