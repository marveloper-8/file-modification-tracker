package service

import (
	"time"
	"log"
)

func StartWorkerThread() {
	for {
		log.Println("Worker thread running")
		time.Sleep(10 * time.Second)
	}
}