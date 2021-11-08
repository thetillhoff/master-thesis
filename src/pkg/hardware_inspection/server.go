package hardware_inspection

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/go-connections/nat"
	"github.com/thetillhoff/eat/pkg/docker"
)

var (
	buildSrc    string = "dnsmasq/"
	imageName   string = "dnsmasq" // TODO: This could include the dockerRegistry-user-id so it can be pushed & pulled -> reduce build time.
	containerID string
	BindIp      string        = "0.0.0.0"
	timeout     time.Duration = 1 * time.Second
	leasePath   string        = "leases"
	liveIsoPath string        = "debian-live-11.1.0-custom.iso"
)

// Start dnsmasq in container
// Later on, include own dhcp server with proxyDHCP capability.
// Additionally, add another package for a dedicated tftp and http server for PXE-boot.
func Start() {
	var (
		err error
		pwd string
	)

	log.Println("INF Starting dhcp and tftp server...")

	pwd, err = os.Getwd()
	if err != nil {
		log.Fatalln("ERR Can't retrieve working directory:", err)
	}

	// Setting correct path for leaseFile
	leasePath = path.Join(pwd, buildSrc, leasePath)

	if err := os.Truncate(leasePath, 0); err != nil {
		log.Fatalln("ERR Can't remove contents of leaseFile;", err)
	}

	// docker.BuildImage(buildSrc, imageName, true) // already done during init

	if debug {
		log.Println("INF live-os-iso-path:", liveIsoPath)
	}

	containerID = docker.StartWithAutoStop(imageName, &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("53/udp"): []nat.PortBinding{{HostIP: BindIp, HostPort: "53"}},
			nat.Port("69/udp"): []nat.PortBinding{{HostIP: BindIp, HostPort: "69"}},
			nat.Port("69/tcp"): []nat.PortBinding{{HostIP: BindIp, HostPort: "69"}},
			nat.Port("80/tcp"): []nat.PortBinding{{HostIP: BindIp, HostPort: "80"}},
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
			{
				Type:   mount.TypeBind,
				Source: pwd + "/dnsmasq/leases",
				Target: "/var/lib/misc/dnsmasq.leases",
			},
		},
	},
		"DHCP=on",
		"BINDIP="+BindIp,
		"ISOSRC="+liveIsoPath,
	)

	// TODO:
	// Only create build context and build image, when image doesn't yet exist.
	// Add parameter for force build

	// TODO:
	// When a container with said image is already running, stop and remove it first.

}

// Stop dhcp
func Stop() {
	log.Println("INF Stopping dhcp and tftp server...")

	docker.Stop(containerID)
}
