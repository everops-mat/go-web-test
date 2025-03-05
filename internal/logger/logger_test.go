//go:build unit

package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"testing"
)

// TestJSONLogger ensures logs are correctly formatted
func TestJSONLogger(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf) // Capture log output

	JSONLogger("info", "Test log message")

	logOutput := buf.String()

	// Extract only the JSON part from the log output
	startIdx := strings.Index(logOutput, "{")
	if startIdx == -1 {
		t.Fatalf("No JSON found in log output: %s", logOutput)
	}

	jsonPart := logOutput[startIdx:]

	// Parse JSON
	var logEntry map[string]string
	err := json.Unmarshal([]byte(jsonPart), &logEntry)
	if err != nil {
		t.Fatalf("Failed to parse log JSON: %v", err)
	}

	// Check log level and message
	if logEntry["level"] != "info" {
		t.Errorf("Expected level 'info', got '%s'", logEntry["level"])
	}

	if logEntry["message"] != "Test log message" {
		t.Errorf("Expected message 'Test log message', got '%s'", logEntry["message"])
	}
}
