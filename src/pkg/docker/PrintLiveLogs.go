package docker

import (
	"bufio"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
)

func PrintLiveLogs(containerID string) {
	if debug {
		log.Println("INF Printing live logs for container '" + containerID + "'.")
	}

	go func() {
		reader, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     true,
			Timestamps: false,
		})
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
}
