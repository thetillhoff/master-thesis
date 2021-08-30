package dhcp_and_tftp

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Cleans the docker stuff.
// Taken from https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in-a-defe
func registerCleanup() {

	// Make sure cleanup is also called on unnormal exits (strg-c)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

		fmt.Println("Cleaning remnants...")
		Stop()

		os.Exit(1)
	}()
}
