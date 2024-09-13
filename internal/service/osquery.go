package main

import (
	"fmt"
	"github.com/osquery/osquery-go"
	"time"
	"log"
)

func getModificationStats(dir string) {
	client, err := osquery.NewClient("127.0.0.1:9001", 5*time.Second)
	if err != nil {
		log.Fatalf("Error creating osquery client: %v", err)
	}
	defer client.Close()

	query := fmt.Sprintf("SELECT * FROM file WHERE directory = '%s';", dir)
	resp, err := client.Query(query)
	if err != nil {
		log.Fatalf("Error querying file events: %v", err)
	}

	log.Printf("File modification stats: %v", resp.Response)
}