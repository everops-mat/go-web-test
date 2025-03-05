package handlers

import (
	"encoding/json"
	"go-web-test/internal/logger"
	"go-web-test/internal/sayings"
	"net/http"
	"os"
)

type HealthzResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Sayings  bool   `json:"sayings"`
	Hostname string `json:"hostname"`
}

// HealthzHandler is a handler for the healthz endpoint
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	sayingsLoaded := len(sayings.GetAllSayings()) > 0

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	reponse := HealthzResponse{
		Status:   "ok",
		Message:  "ok",
		Sayings:  sayingsLoaded,
		Hostname: hostname,
	}

	logger.JSONLogger("info", "Health check requested")

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(reponse); err != nil {
		logger.JSONLogger("error", "Error encoding healthz response")
		logger.JSONLogger("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

}
