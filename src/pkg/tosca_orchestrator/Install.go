package tosca_orchestrator

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/thetillhoff/eat/pkg/csar"
	"github.com/thetillhoff/eat/pkg/hardware_inspection"
)

// TODO bindIp is only used for init command, not for install command. This can be changed when the dnsmasq container allows to set the bindip at runtime via environment variables.
func Install(archive csar.CSAR, inputs []string, bindIp string) {
	if debug {
		// Set debug for imports
		hardware_inspection.Debug()
	}

	// Add derivation from NodeType to NodeTemplate, so all properties etc are available in the base type.
	// TODO: This could be merged with InstantiateTpoplogyTemplate - not sure if that makes more sense though.
	*archive.ServiceTemplate = DeriveServiceTemplate(archive) // TODO copy from test

	// Fill in the provided inputs
	*archive.ServiceTemplate.TopologyTemplate = AddInputsToTopologyTemplate(*archive.ServiceTemplate.TopologyTemplate, inputs) // TODO copy from test

	// Verify input contains macAddresses and path to ssh-key
	// TODO Verify inputs (like validity of macaddresses and existance of file at provided path to ssh-key)

	// Merge types and templates, fill with default values where needed, run normativeFunctions, verify constraints, ...
	*archive.ServiceTemplate.TopologyTemplate = InstantiateTopologyTemplate(*archive.ServiceTemplate)

	// Build <live-os>.iso - Can be ommited for now and a fixed path of a precreated iso can be selected.
	// var liveIsoPath string = "debian-live.iso"

	// When a new machine attempts to boot over pxe, it ...
	// 1. gets a random ip-address from the preexisting dhcp
	// 2. loads a fitting (bios or efi) variant of ipxe from the tftp-server
	// 3. requests the ipxe configuration from the http-server at /default - This contains the location of the <live-os>.iso
	// 4. loads the <live-os>.iso from the http-server at /isos/<live-os>.iso and boot it
	// 5. runs the <live-os> with a embedded webserver, which
	//    - provides hardware-information at /cpuinfo.txt and /meminfo.txt
	//    - runs a ssh-server with a preconfigured ssh-key

	var machines map[string]hardware_inspection.Machine = make(map[string]hardware_inspection.Machine)

	// Parse macAddresses from Input and retrieve hardware information
	if value, ok := archive.ServiceTemplate.TopologyTemplate.Inputs["macAddresses"].Value.(string); !ok {
		log.Fatalln("ERR Couldn't parse input 'macAddresses' to slice")
	} else {
		macAddresses := strings.Split(value, ",")

		if debug {
			log.Println("INF " + strconv.Itoa(len(macAddresses)) + " MAC-addresses provided.")
		}

		// Start dhcp-, tftp- and http-servers
		hardware_inspection.BindIp = bindIp
		hardware_inspection.Start()

		// Retrieve hardware information
		machines = hardware_inspection.InspectMacs(macAddresses)

		// Stop dhcp-, tftp- and http-servers
		hardware_inspection.Stop()
	}

	if debug {
		log.Println("INF " + strconv.Itoa(len(machines)) + " machines are available.")
	}

	// Select some servers out of available ones
	var selectedMachine hardware_inspection.Machine
	// TODO add proper selection process
	for _, machine := range machines {
		selectedMachine = machine
		break
	}

	fmt.Println("selected machine:", selectedMachine)

	// For all machines in parallel:
	//   - Either shutdown if not a selected node
	//   - Or configure if a selected node
	// Optionally do something with the not-selected servers (f.e. shut them down)

	// Run installation and configuration scripts on selected hosts (via ssh)
	// TODO Thoughts
	//   Assume a webserver and a database are to be installed - how are the servers selected?
	//   -> One after another, this means first the whole process runs for the webserver, then for the database
	//      Make sure the selected server isn't always the first in the list that fits, but a random one.
	// ssh.RunCommandOnHost(selectedMachine.IpAddress, "hostname") // TODO Replace "hostname" with actual command(s) to run

	// TODO
	// After installation succeeded, add a SRV entry in the DNS. This can then later be used for state detection.
	// State detection == is the desired state already reached?
	// When no DNS is active, use dnsmasq as DNSserver. Dnsmasq uses the local /etc/hosts, and SRV entries can be added there.

	// Collect values for outputs specified in topologyTemplate
	*archive.ServiceTemplate.TopologyTemplate = AddOutputsOfTopologyTemplate(*archive.ServiceTemplate.TopologyTemplate)

	// Print desired outputs of topologyTemplate
	for key, output := range archive.ServiceTemplate.TopologyTemplate.Outputs {
		fmt.Println(key, output)
	}
}
