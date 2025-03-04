package tests

import (
	"io"
	"os"
	"testing"
)

// Read the test file
func TestReadFile(t *testing.T) {

	fileName := "/tmp/testfile.txt"
	content := []byte("Hello, World!")

	if err := os.WriteFile(fileName, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Ensure file gets removed after test
	defer os.Remove(fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Expected file content, but got empty data: %v", data)
	}
}

// Write to the test file
func TestWriteFile(t *testing.T) {
	fileName := "/tmp/testfile.txt"
	content := []byte("Hello, World!")
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	defer os.Remove(fileName) // Cleanup
}

// Test read from standard input
func TestIOCloser(t *testing.T) {
	reader := io.NopCloser(os.Stdin)
	if reader == nil {
		t.Errorf("Expected a non-nil io.ReadCloser")
	}
}
