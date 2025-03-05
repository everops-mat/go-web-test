package main

import (
	"context"
	"fmt"
	"go-web-test/internal/handlers"
	"go-web-test/internal/logger"
	"go-web-test/internal/sayings"
	"go-web-test/internal/signals"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Configuration
const (
	defaultSayingsFile = "config/sayings.txt"
	serverPort         = 8080
)

func main() {
	// Ensure sayings file exists
	if _, err := os.Stat(defaultSayingsFile); os.IsNotExist(err) {
		logger.JSONLogger("fatal", fmt.Sprintf("Sayings file not found: %s", defaultSayingsFile))
		os.Exit(1)
	}

	// Load sayings at startup
	if err := sayings.LoadSayings(defaultSayingsFile); err != nil {
		logger.JSONLogger("fatal", fmt.Sprintf("Failed to load sayings: %v", err))
		os.Exit(1)
	}

	// Handle SIGHUP for reloading
	go signals.HandleSignals(defaultSayingsFile)

	// Add middleware to the handler
	mux := http.NewServeMux()
	mux.Handle("/", handlers.AuthMiddleware(http.HandlerFunc(handlers.RandomSayingHandler)))
	mux.Handle("/healthz", http.HandlerFunc(handlers.HealthzHandler))

	// Create HTTP server with graceful shutdown support
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", serverPort),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,   // Prevents Slowloris attack
		ReadTimeout:       10 * time.Second,  // Limits request body read time
		WriteTimeout:      10 * time.Second,  // Limits response write time
		IdleTimeout:       120 * time.Second, // Keep-alive timeout
	}

	// Run server in a separate goroutine
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		logger.JSONLogger("info", fmt.Sprintf("Server starting on port %d...", serverPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.JSONLogger("fatal", fmt.Sprintf("Server error: %v", err))
			os.Exit(1)
		}
	}()

	// Handle SIGINT and SIGTERM for clean shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop // Wait for a termination signal
	logger.JSONLogger("info", "Shutting down server...")

	// Shutdown server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.JSONLogger("error", fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	wg.Wait()
	logger.JSONLogger("info", "Server stopped gracefully.")
}
