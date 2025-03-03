package gotodosrv

import (
	"gotodo/internal/api/handlers"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckShouldReturn200(t *testing.T) {
	req := httptest.NewRequest("GET", "/diagnostics/healthcheck", nil)
	w := httptest.NewRecorder()
	handlers.HandleHealthCheck(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("Health check failed with status code %d", resp.StatusCode)
	}
}
