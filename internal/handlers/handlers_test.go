//go:build unit

package handlers

import (
	"go-web-test/internal/sayings"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestRandomSayingHandler ensures it returns an HTML response with a saying
func TestRandomSayingHandler(t *testing.T) {
	// Set up some test sayings
	sayings.LoadSayingsFromSlice([]string{
		"Test Saying 1",
		"Test Saying 2",
		"Test Saying 3",
	})

	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RandomSayingHandler)

	// Serve HTTP request
	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Ensure response contains at least one test saying
	responseBody := rr.Body.String()
	found := false
	for _, saying := range sayings.GetAllSayings() {
		if strings.Contains(responseBody, saying) {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected one of the sayings in response body, but got:\n%s", responseBody)
	}
}

// TestRandomSayingHandlerEmpty ensures it handles the case where no sayings exist
func TestRandomSayingHandlerEmpty(t *testing.T) {
	// Clear sayings
	sayings.LoadSayingsFromSlice([]string{})

	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RandomSayingHandler)

	// Serve HTTP request
	handler.ServeHTTP(rr, req)

	// Check response status code (should return 500)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusInternalServerError)
	}

	// Check response contains "No sayings available"
	if !strings.Contains(rr.Body.String(), "No sayings available") {
		t.Errorf("Expected 'No sayings available' message, but got:\n%s", rr.Body.String())
	}
}
