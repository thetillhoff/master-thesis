package csar

import "github.com/thetillhoff/eat/pkg/tosca"

var (
	debug bool = false // Output more verbose when true
)

func Debug() {
	debug = true
	tosca.Debug()
}
