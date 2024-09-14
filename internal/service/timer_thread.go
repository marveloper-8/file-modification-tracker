package service

import (
	"github.com/marveloper-8/file-modification-tracker/internal/api"
	"time"
	"log"
)

func StartTimerThread(apiClient *api.APIClient) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		stats := &api.Stats{
			Directory: "C:\\Users\\marve\\Downloads",
			FileName:  "example.txt",
			Status:    "modified",
		}
		err := apiClient.SendStats(stats)
		if err != nil {
			log.Printf("Error sending stats: %v", err)
		}
	}
}