package dhcp_and_tftp

import "github.com/thetillhoff/eat/pkg/docker"

var (
	debug bool = false // Output more verbose when true
)

func Debug() {
	debug = true

	docker.Debug()
}
