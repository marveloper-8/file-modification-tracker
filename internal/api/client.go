package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"net/http"
)

type Stats struct {
	Directory string `json:"directory"`
	Modified time.Time `json:"modified"`
	Files     []string `json:"files"`
}

func reportStatsToAPI(stats Stats) error {
	data, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/stats", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.Status)
	}

	return nil
}