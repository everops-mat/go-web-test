//go:build unit

package handlers

import (
	"encoding/json"
	"go-web-test/internal/sayings"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TesthealthzHandler tests the healthz handler.
func TestHealthzHandler(t *testing.T) {
	// create some sayings
	sayings.LoadSayingsFromSlice([]string{"Test 1", "Test 2"})

	// create a request
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthzHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response HealthzResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	if response.Status != "ok" {
		t.Errorf("Expected response status to be 'ok', got %v", response.Status)
	}

	if !response.Sayings {
		t.Errorf("Expected sayings to be true, got %v", response.Sayings)
	}
}
