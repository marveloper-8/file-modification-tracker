package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpServer "file-modification-tracker/internal/adapters/http"
	"file-modification-tracker/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {

	service := &core.Service{}
	server := httpServer.NewHTTPServer(service)

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	server.HealthHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Contains(t, rr.Body.String(), `"status":"healthy"`)
}
