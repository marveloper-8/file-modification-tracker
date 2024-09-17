package http

import (
    "encoding/json"
    "net/http"
    "file-modification-tracker/internal/core"
    "file-modification-tracker/internal/adapters/daemon"
    "file-modification-tracker/internal/adapters/logs"
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
    http.HandleFunc("/health", s.HealthHandler)
    http.HandleFunc("/commands", s.CommandHandler)
    http.HandleFunc("/logs", s.LogsHandler)

    return http.ListenAndServe(":8080", nil)
}

func (s *HTTPServer) HealthHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func (s *HTTPServer) CommandHandler(w http.ResponseWriter, r *http.Request) {
    var commands []string
    json.NewDecoder(r.Body).Decode(&commands)

    for _, cmd := range commands {
        daemon.AddCommand(cmd)
    }

    json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}

func (s *HTTPServer) LogsHandler(w http.ResponseWriter, r *http.Request) {
    logs := logs.NewLoggerAdapter().RetrieveLogs()
    json.NewEncoder(w).Encode(logs)
}