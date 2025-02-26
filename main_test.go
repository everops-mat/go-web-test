package main

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunServer_MissingCommand(t *testing.T) {
	server, err := runServer("", "", ":9991")
	if err == nil {
		t.Errorf("Expected error when cmd is empty, but got nil")
	}
	if server != nil {
		t.Errorf("Expected server to be nil, but got %v", server)
	}
}

func TestRunServer_ValidCGI(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello from CGI"))
		if err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer ts.Close()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Errorf("Failed to listen on port: %v", err)
	}
	runServerAddr := listener.Addr().String()
	listener.Close()

	server, err := runServer("/bin/echo", "", runServerAddr)
	if err != nil {
		t.Errorf("runServer failed with error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		t.Errorf("Failed to shutdown server: %v", err)
	}
}
