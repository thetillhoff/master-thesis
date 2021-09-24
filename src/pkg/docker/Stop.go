package docker

import (
	"log"
	"time"

	"github.com/docker/docker/api/types"
)

var (
	Timeout = time.Second * 1 // Default kill timeout of 1s
)

// Stop container with provided id
func Stop(containerID string) {
	var (
		err        error
		containers []types.Container
	)

	if debug {
		log.Println("INF Stopping container '" + containerID + "'...")
	}

	containers, err = cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatalln("ERR Can't retrieve list of containers:", err)
	}

	for _, container := range containers {
		if container.ID == containerID {
			// Stop Container
			// Due to autoremove set at ContainerCreate, removal is done automatically -> only stopping it is required
			err = cli.ContainerStop(ctx, containerID, &Timeout)
			if err != nil {
				log.Fatalln("ERR Can't stop docker container:", err)
			}
			break
		}
	}
	if debug {
		log.Println("INF Container already stopped")
	}
}
