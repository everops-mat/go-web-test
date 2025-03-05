//go:build intergration

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

// TestSayingExists validates that the generated saying is from config/sayings
func TestSayingExists(t *testing.T) {
	// Load sayings from config/sayings
	data, err := ioutil.ReadFile("../config/sayings.txt")
	if err != nil {
		t.Fatalf("Failed to read sayings file: %v", err)
	}

	// Convert sayings file content into a map for fast lookup
	sayings := strings.Split(string(data), "\n")
	sayingSet := make(map[string]struct{})
	for _, s := range sayings {
		s = strings.TrimSpace(s)
		if s != "" {
			sayingSet[s] = struct{}{}
		}
	}

	// Make a request to the web server
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Failed to get response from server: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response body
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	// Extract the excuse text
	excuse := doc.Find(".excuse").Text()
	excuse = strings.TrimSpace(excuse)

	// Verify the excuse exists in the sayings file
	_, exists := sayingSet[excuse]
	assert.True(t, exists, fmt.Sprintf("Generated saying '%s' not found in config/sayings", excuse))
}
