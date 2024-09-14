package main
import (
	"time"
	"github.com/kardianos/service"
	"log"
)

var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}

	p.exit = make(chan struct{})
	
	go p.runWorkerThread()
	go p.runTimerThread()

	return nil
}

func (p *program) runWorkerThread() {
	logger.Info("Worker thread started")
	for {
		select {
			case <-p.exit:
				return
			default:
				logger.Info("Executing shell command")
				time.Sleep(1 * time.Second)
		}
	}
}

func (p *program) runTimerThread() {
	logger.Info("Timer thread started")
	for {
		select {
			case <-p.exit:
				return
			default:
				logger.Info("Checking file modifications...")

				stats := getFileModificationStats(config.Directory)
				if err := reportStatsToAPI(stats); err != nil {
					logger.Error("Failed to report stats to API: ", err)
				}
				time.Sleep(1 * time.Minute)
		}
	}
}

func (p *program) Stop(s service.Service) error {
	close(p.exit)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name: "FileModificationTracker",
		DisplayName: "File Modification Tracker Service",
		Description: "Tracks file modifications in a specified directory",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}