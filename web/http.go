package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type HealthResponse struct {
	WorkerStatus string `json:"worker_status"`
	TimerStatus string `json:"timer_status"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		WorkerStatus: "running",
		TimerStatus: "running",
	}
	json.NewEncoder(w).Encode(response)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	logs := []string{"Log 1", "Log 2", "Log 3"}
	json.NewEncoder(w).Encode(logs)
}

func startHttpServer() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")
	r.HandleFunc("/logs", logsHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}