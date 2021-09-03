package docker

import (
	"log"

	"github.com/docker/docker/api/types/container"
)

func Wait(containerID string) {
	if debug {
		log.Println("INF Waiting for container '" + containerID + "' to finish.")
	}

	// Both cases (err and status) mean the container is finished.
	statusCh, errCh := cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Println("ERR Waiting for container to finish failed:", err)
		}
	case _ = <-statusCh:
		// Finished successfully
	}

}
