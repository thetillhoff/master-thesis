package hardware_inspection

import (
	"log"
	"time"

	"github.com/thetillhoff/eat/pkg/wol"
)

// Requires Start() to be executed before
// wol every macAddress
//   TODO not for vms, instead use cli-commands
//   TODO check if machine is already up; if yes, skip it
// for each mac, while ip, ram, cores is empty
//   refresh leases
//   get ip from newest lease
//   get ram via ip
//   get cores via ip
//   wait
// Requires Stop() to be executed afterwards
func InspectMacs(macAddresses []string) map[string]Machine {
	var (
		machines map[string]Machine = make(map[string]Machine)
	)

	for _, macAddress := range macAddresses {
		machines[macAddress] = Machine{MacAddress: macAddress}
	}

	// Waking every machine
	for macAddress := range machines {
		wol.Wake(macAddress)
	}

	for macAddress := range machines {
		// While cores or ram are yet missing
		for machines[macAddress].Cores == 0 || machines[macAddress].Ram == 0 {
			log.Println("INF Waiting for information about machine with macAddress", macAddress)

			// Update leases
			updateLeaseTable()

			// If machine has a lease
			if foundLease, ok := leaseTable[macAddress]; ok {
				machine := machines[macAddress]

				// Update ipAddress
				machine.IpAddress = foundLease.ipAddress

				// Try to update ram if necessary
				if machine.Ram == 0 {
					machine.Ram = getMemInfo(machine.IpAddress)
				}

				// Try to update cores if necessary
				if machine.Cores == 0 {
					machine.Cores = getCpuInfo(machine.IpAddress)
				}

				machines[macAddress] = machine
			}

			// wait
			time.Sleep(2 * time.Second) // No need to refresh every ms or even less
		}
	}

	return machines
}
