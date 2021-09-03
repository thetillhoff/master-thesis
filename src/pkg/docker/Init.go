package docker

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

var (
	cli *client.Client
	ctx context.Context
)

func Init() {
	var (
		err error
	)

	// Only if not already initialized
	if cli == nil {
		// Create docker client
		cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatalln("ERR Couldn't create docker client:", err)
		}
	}

	// Only if not already initialized
	if ctx == nil {
		ctx = context.Background()
	}
}
