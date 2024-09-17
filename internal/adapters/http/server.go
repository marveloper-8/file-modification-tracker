package http

import (
    "encoding/json"
    "file-modification-tracker/internal/core"
    "file-modification-tracker/internal/adapters/daemon"
    "file-modification-tracker/internal/adapters/logs"
    "net/http"
)

type HTTPServer struct {
    service *core.Service
}

func NewHTTPServer(service *core.Service) *HTTPServer {
    return &HTTPServer{
        service: service,
    }
}

func (s *HTTPServer) Run() error {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/commands", commandHandler)
    http.HandleFunc("/logs", logsHandler)

    return http.ListenAndServe(":8080", nil)
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
    logs.NewLoggerAdapter().RetrieveLogs()
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
