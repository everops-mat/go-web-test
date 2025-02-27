package signals

import (
	"go-web-test/internal/logger"
	"go-web-test/internal/sayings"
	"os"
	"os/signal"
	"syscall"
)

// HandleSignals listens for SIGHUP and reloads sayings
func HandleSignals(filename string) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP)

	go func() {
		for sig := range sigChan {
			if sig == syscall.SIGHUP {
				logger.JSONLogger("info", "Received SIGHUP, reloading sayings...")

				// Reload sayings from file
				if err := sayings.LoadSayings(filename); err != nil {
					logger.JSONLogger("error", "Failed to reload sayings: "+err.Error())
				} else {
					logger.JSONLogger("info", "Successfully reloaded sayings.")
				}
			}
		}
	}()
}
