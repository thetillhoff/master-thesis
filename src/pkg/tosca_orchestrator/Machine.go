package tosca_orchestrator

import "github.com/thetillhoff/eat/pkg/tosca"

// This type is used to mirror the actual physical servers.
type Machine struct {
	IpAddress string

	Ram   tosca.Size
	Cores tosca.Integer
}

func GetRam(machine Machine) tosca.Size {
	// TODO make the actual request to the ip of the machine
	return tosca.Size{Value: 4}
}

func GetCores(machine Machine) tosca.Integer {
	// TODO make the actual request to the ip of the machine
	return tosca.Integer{Value: 4}
}
