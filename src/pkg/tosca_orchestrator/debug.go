package tosca_orchestrator

import (
	"github.com/thetillhoff/eat/pkg/hardware_inspection"
)

var (
	debug bool = false // Output more verbose when true
)

func Debug() {
	debug = true

	hardware_inspection.Debug()
}
