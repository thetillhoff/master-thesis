package tosca_orchestrator

import (
	"github.com/thetillhoff/eat/pkg/hardware_inspection"
	"github.com/thetillhoff/eat/pkg/ssh"
)

var (
	debug bool = false // Output more verbose when true
)

func Debug() {
	debug = true

	hardware_inspection.Debug()
	ssh.Debug()
}
