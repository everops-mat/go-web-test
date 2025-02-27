package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

// JSONLogger logs messages in JSON format
func JSONLogger(level, message string) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
	}
	log.Println(toJSON(entry))
}

// toJSON converts a struct into a JSON string
func toJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf(`{"error": "failed to log message: %v"}`, err)
	}
	return string(data)
}
