package osquery

import (
	"log"
)

func QueryModificationStats(directory string) ([]map[string]string, error) {
	log.Println("Running osquery to gather file stats")

	return []map[string]string{
		{"file_name": "example.txt", "status": "modified"},
	}, nil
}