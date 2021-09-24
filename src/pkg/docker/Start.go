package docker

import (
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

// Start container with provided imagename and hostconfig, returns containerID
func Start(imageName string, hostconfig *container.HostConfig, env ...string) string {
	var (
		err  error
		resp container.ContainerCreateCreatedBody
	)

	// If not initialized
	if cli == nil {
		Init() // Initialize
	}

	log.Println("INF Running container for image '" + imageName + "'.")

	// pwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalln("ERR Can't retrieve working directory:", err)
	// }

	// Create Container
	// Expose ports, set privileged, enable NET_ADMIN, delete after stop, mount ./isos to /http/isos in container
	resp, err = cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		// Cmd:          strslice.StrSlice{"bash"},
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
		AttachStdin:  debug, // This ensures input is possible when in debug mode
		OpenStdin:    debug,
		Env:          env,
	}, hostconfig,
		nil, nil, "")
	if err != nil {
		log.Fatalln("ERR Can't create docker container:", err)
	}

	// Start previously created container
	if err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatalln(err)
	}

	if debug {
		log.Println("INF Container started. ID:", resp.ID)
	}

	if debug {
		PrintLiveLogs(resp.ID)
	}

	return resp.ID
}
