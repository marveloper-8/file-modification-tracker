package core

import "time"

// Service defines the core structure of the file modification service.
type Service struct {
	Config       ConfigPort
	Logger       LoggerPort
	FileChecker  FileCheckerPort
	CommandQueue CommandQueuePort
}

// NewService creates a new service instance.
func NewService(config ConfigPort, logger LoggerPort, fileChecker FileCheckerPort, commandQueue CommandQueuePort) *Service {
	return &Service{
		Config:       config,
		Logger:       logger,
		FileChecker:  fileChecker,
		CommandQueue: commandQueue,
	}
}

// StartService runs the service.
func (s *Service) StartService() {
	go s.workerThread()
	go s.timerThread()
}

func (s *Service) workerThread() {
	for cmd := range s.CommandQueue.ReceiveCommands() {
		s.Logger.LogInfo("Executing command: " + cmd)
		// handle command execution
	}
}

func (s *Service) timerThread() {
	for {
		time.Sleep(time.Duration(s.Config.GetCheckFrequency()) * time.Second)
		files, err := s.FileChecker.CheckModifications(s.Config.GetDirectory())
		if err != nil {
			s.Logger.LogError(err)
			continue
		}
		s.Logger.LogFileStats(files)
	}
}
