package dhcp_and_tftp

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

var (
	buildSrc    string = "dnsmasq/"
	imageName   string = "dnsmasq" // TODO: This could include the dockerRegistry-user-id so it can be pushed & pulled -> reduce build time.
	containerID string
	bindIp      string        = "0.0.0.0"
	timeout     time.Duration = 1 * time.Second

	ctx context.Context = context.Background()
)

// Start dnsmasq in container
// Later on, include own dhcp server with proxyDHCP capability.
// Additionally, add another package for a dedicated tftp and http server for PXE-boot.
func Start() {
	var (
		err        error
		respBuild  types.ImageBuildResponse
		respCreate container.ContainerCreateCreatedBody

		scanner *bufio.Scanner
	)

	log.Println("INF Starting dhcp and tftp server...")

	// Ensure the docker container is stopped after this app is stopped.
	registerCleanup()

	// Create docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalln("ERR Couldn't create docker client:", err)
	}

	// TODO:
	// Only create build context and build image, when image doesn't yet exist.
	// Add parameter for force build

	// TODO:
	// When a container with said image is already running, stop and remove it first.

	// Creating build context
	log.Println("INF Creating build context...")
	buildContextFilePath := createBuildContext(buildSrc)
	buildContextFile, err := os.Open(buildContextFilePath)
	if err != nil {
		log.Println("ERR Couldn't open buildcontext at '"+buildContextFilePath+"':", err)
	}

	// Build container image
	respBuild, err = cli.ImageBuild(
		ctx,
		buildContextFile,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Tags:       []string{imageName},
			// NoCache:    true,
			Remove: true,
			// BuildArgs: make(map[string]*string),
		})
	if err != nil {
		log.Fatalln("ERR Couldn't create container image:", err)
	}
	buildContextFile.Close()
	os.Remove(buildContextFilePath)
	defer respBuild.Body.Close()

	// Output is json stream; Print it in readable way
	scanner = bufio.NewScanner(respBuild.Body)
	for scanner.Scan() {
		line := scanner.Text()
		data := map[string]interface{}{}
		_ = json.Unmarshal([]byte(line), &data) // Don't care about errors during unmarshal (trusting docker)
		if stream, ok := data["stream"]; ok {
			if parsedLine, ok := stream.(string); ok {
				log.Print(parsedLine) // No Println here, since the unmarshalled json already contains newlines
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalln("ERR There was an error while creating the container image:", err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("ERR Couldn't retrieve working directory:", err)
	}

	// Create Container
	// Expose ports, set privileged, enable NET_ADMIN, delete after stop, mount ./isos to /http/isos in container
	respCreate, err = cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		// Cmd:          strslice.StrSlice{"sleep", "60"},
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
		// AttachStdin:  true,
		// OpenStdin:    true,
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("53/udp"): []nat.PortBinding{{HostIP: bindIp, HostPort: "53"}},
			nat.Port("69/udp"): []nat.PortBinding{{HostIP: bindIp, HostPort: "69"}},
			nat.Port("69/tcp"): []nat.PortBinding{{HostIP: bindIp, HostPort: "69"}},
			nat.Port("80/tcp"): []nat.PortBinding{{HostIP: bindIp, HostPort: "80"}},
		},
		Privileged: true,
		CapAdd: strslice.StrSlice{
			"NET_ADMIN",
		},
		NetworkMode: "host",
		AutoRemove:  true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: pwd + "/dnsmasq/isos",
				Target: "/http/isos",
			},
		},
	}, nil, nil, "")
	if err != nil {
		log.Fatalln("ERR Couldn't create docker container:", err)
	}

	// Store container id
	containerID = respCreate.ID

	// Start previously created container
	if err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		log.Fatalln(err)
	}

	log.Println("INF Container started. ID:", containerID)

	// log.Println("INF Waiting for container to finish.")
	// // Both cases (err and status) mean the container is finished.
	// statusCh, errCh := cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	// select {
	// case err := <-errCh:
	// 	if err != nil {
	// 		log.Println("ERR Waiting for container to finish failed:", err)
	// 	}
	// case status := <-statusCh:
	// 	log.Println("INF statuscode:", status)
	// }
}

// Stop container
func Stop() {
	var (
		err        error
		containers []types.Container
	)

	log.Println("INF Stopping dhcp and tftp server...")

	// Create docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalln("ERR Couldn't create docker client:", err)
	}

	containers, err = cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatalln("ERR Couldn't retrieve list of containers:", err)
	}

	for _, container := range containers {
		if container.ID == containerID {
			// Stop Container
			// Due to autoremove set at ContainerCreate, removal is done automatically
			err = cli.ContainerStop(ctx, containerID, &timeout)
			if err != nil {
				log.Fatalln("ERR Couldn't stop docker container:", err)
			}
			break
		}
	}
	if debug {
		log.Println("INF Container already stopped")
	}
}
