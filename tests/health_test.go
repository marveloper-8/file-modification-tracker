package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "file-modification-tracker/internal/core"
    httpServer "file-modification-tracker/internal/adapters/http"
)

func TestHealthHandler(t *testing.T) {
    // Create a new HTTPServer instance
    service := &core.Service{} // You might need to mock this
    server := httpServer.NewHTTPServer(service)

    // Create a request to pass to our handler
    req, err := http.NewRequest("GET", "/health", nil)
    assert.NoError(t, err)

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the handler directly
    server.HealthHandler(rr, req)

    // Check the status code
    assert.Equal(t, http.StatusOK, rr.Code)

    // Check the response body
    assert.Contains(t, rr.Body.String(), `"status":"healthy"`)
}