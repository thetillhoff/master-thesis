package hardware_inspection

import "github.com/thetillhoff/eat/pkg/docker"

var (
	debug bool = false // Output more verbose when true
)

func Debug() {
	debug = true

	docker.Debug()
}
