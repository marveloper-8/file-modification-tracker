package api

import (
	"bytes"
	"encoding/json"
	"log"
	"time"
	"net/http"
)

type Stats struct {
	Directory string `json:"directory"`
	ModifiedTime string `json:"modified_time"`
	FileName string `json:"file_name"`
	Status string `json:"status"`
}

type APIClient struct {
	client *http.Client
	baseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		client: &http.Client{Timeout: 10 * time.Second},
		baseURL: baseURL,
	}
}

func (api *APIClient) SendStats(stats *Stats) error {
	jsonData, err := json.Marshal(stats)
	if err != nil {
		log.Printf("Error marshalling stats: %v", err)
		return err
	}
	
	req, err := http.NewRequest("POST", api.baseURL+"/file-stats", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := api.client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return err
	}

	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		log.Printf("API responded with status: %v", resp.StatusCode)
		return err
	}

	log.Printf("Stats sent successfully")
	return nil
}