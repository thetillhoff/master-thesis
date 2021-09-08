package wol

import (
	"strings"

	"github.com/thetillhoff/eat/pkg/script_runner"
)

// Returns map[macAddress]vmName
func getVMs() map[string]string {
	var (
		vms map[string]string = make(map[string]string)
	)

	output := script_runner.RunLinuxCommand("sudo virsh list --all --name") // TODO remove sudo as soon as "groups" works without error

	for _, vmName := range strings.Split(output, " ") {
		if vmName == "" { // the output of 'virsh list' contains an empty line
			continue
		}
		macAddress := script_runner.RunLinuxCommand("sudo virsh dumpxml " + vmName + " | grep 'mac address' | cut -d\"'\" -f2") // TODO remove sudo as soon as "groups" works without error
		vms[macAddress] = vmName
	}

	return vms
}
