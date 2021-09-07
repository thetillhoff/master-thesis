package docker

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/docker/docker/api/types/container"
)

func StartWithAutoStop(imageName string, hostconfig *container.HostConfig, env ...string) string {
	var (
		containerID string
	)

	containerID = Start(imageName, hostconfig, env...)

	// Taken from https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in-a-defe
	// Make sure cleanup is also called on unnormal exits (strg-c)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

		fmt.Println("Cleaning remnants...")
		Stop(containerID)

		os.Exit(1)
	}()

	if debug {
		log.Println("INF Autostop for container with ID '" + containerID + "' added.")
	}

	return containerID
}
