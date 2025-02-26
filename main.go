package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
)

// Define flags
var (
	cmd     = flag.String("cmd", "", "CGI program to run.")
	wd      = flag.String("wd", "", "Working directory for CGI.")
	address = flag.String("address", ":9991", "Listen address.")
)

func main() {
	flag.Parse()

	_, err := runServer(*cmd, *wd, *address)
	if err != nil {
		log.Fatal(err)
	}

	// Block main thread to keep server running
	select {}
}

// runServer starts an HTTP server with a CGI handler and returns the server instance.
func runServer(cmd, wd, addr string) (*http.Server, error) {
	if cmd == "" {
		return nil, errors.New("missing required parameter: -cmd")
	}

	if len(cmd) > 0 && cmd[0] != '/' {
		cmd = "./" + cmd
	}

	os.Setenv("PATH", os.Getenv("PATH")+":.")

	cgiHandler := &cgi.Handler{
		Path:       cmd,
		Root:       "/",
		Dir:        wd,
		InheritEnv: []string{"PATH", "PLAN9"},
	}

	log.Println("Starting HTTP server listening on", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: cgiHandler,
	}

	// Run the server in a separate goroutine to avoid blocking execution
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	return server, nil
}
