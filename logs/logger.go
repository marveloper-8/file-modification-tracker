package logs

import (
	"log"
	"sync"
)

var logs = make([]string, 0)
var mu sync.Mutex

func LogError(err error) {
	mu.Lock()
	defer mu.Unlock()
	log.Println("Error: ", err)
	logs = append(logs, err.Error())
}

func LogFileStats(stats string) {
	mu.Lock()
	defer mu.Unlock()
	logs = append(logs, stats)
}

func RetrieveLogs() []string {
	mu.Lock()
	defer mu.Unlock()
	return logs
}