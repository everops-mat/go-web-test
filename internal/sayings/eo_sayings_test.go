package sayings

import (
	"os"
	"testing"
)

// TestLoadSayings ensures sayings are loaded correctly
func TestLoadSayings(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_sayings.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := "Test Saying 1\nTest Saying 2\nTest Saying 3\n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Load sayings
	if err := LoadSayings(tmpFile.Name()); err != nil {
		t.Fatalf("LoadSayings failed: %v", err)
	}

	// Check sayings count
	mu.RLock()
	defer mu.RUnlock()
	if len(sayings) != 3 {
		t.Errorf("Expected 3 sayings, got %d", len(sayings))
	}
}

// TestGetRandomSaying ensures a saying is returned
func TestGetRandomSaying(t *testing.T) {
	mu.Lock()
	sayings = []string{"Saying A", "Saying B", "Saying C"}
	mu.Unlock()

	saying, err := GetRandomSaying()
	if err != nil {
		t.Errorf("Expected a saying but got error: %v", err)
	}

	if saying == "" {
		t.Errorf("Expected a saying but got an empty string")
	}
}

// TestGetRandomSayingEmpty ensures error is returned when there are no sayings
func TestGetRandomSayingEmpty(t *testing.T) {
	mu.Lock()
	sayings = []string{} // Empty list
	mu.Unlock()

	_, err := GetRandomSaying()
	if err == nil {
		t.Errorf("Expected an error for empty sayings list, but got none")
	}
}
