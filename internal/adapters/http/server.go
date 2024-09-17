package http

import (
	"encoding/json"
	"net/http"
	"file-modification-tracker/internal/adapters/daemon"
	"file-modification-tracker/internal/adapters/logs"
)


func Run() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/commands", commandHandler)
	http.HandleFunc("/logs", logsHandler)

	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	var commands []string
	json.NewDecoder(r.Body).Decode(&commands)

	for _, cmd := range commands {
		daemon.AddCommand(cmd)
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	logs.RetrieveLogs()
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

	