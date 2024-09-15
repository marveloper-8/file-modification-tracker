package daemon
import (
	"file-modification-tracker/config"
	"file-modification-tracker/osquery"
	"file-modification-tracker/logs"
	"time"
)

func Run() error {
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
		files, err := osquery.GetFileModifications(config.Config.Directory)
		if err != nil {
			logs.LogError(err)
		}
		logs.LogFileStats(files)
	}
}