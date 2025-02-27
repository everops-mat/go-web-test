package signals

import (
	"go-web-test/internal/sayings"
	"os"
	"syscall"
	"testing"
	"time"
)

// TestSIGHUPReload ensures SIGHUP triggers a file reload
func TestSIGHUPReload(t *testing.T) {
	// Create a temporary sayings file
	tmpFile, err := os.CreateTemp("", "test_sayings.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write initial sayings
	content := "Initial Saying\n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Load initial sayings
	if err := sayings.LoadSayings(tmpFile.Name()); err != nil {
		t.Fatalf("Initial LoadSayings failed: %v", err)
	}

	// Start signal handling in a separate goroutine
	go HandleSignals(tmpFile.Name())

	// Modify the sayings file
	newContent := "New Saying 1\nNew Saying 2\n"
	if err := os.WriteFile(tmpFile.Name(), []byte(newContent), 0644); err != nil {
		t.Fatalf("Failed to modify temp file: %v", err)
	}

	// Send SIGHUP
	if err := syscall.Kill(syscall.Getpid(), syscall.SIGHUP); err != nil {
		t.Fatalf("Failed to send SIGHUP: %v", err)
	}

	// Allow time for the signal to be processed
	time.Sleep(1 * time.Second)

	// Verify that the sayings list has been updated
	saying, err := sayings.GetRandomSaying()
	if err != nil {
		t.Fatalf("Expected a saying but got error: %v", err)
	}

	if saying != "New Saying 1" && saying != "New Saying 2" {
		t.Errorf("Expected new sayings, but got: %s", saying)
	}
}
