package tosca_orchestrator

import (
	"fmt"

	"github.com/thetillhoff/eat/pkg/arp"
	"github.com/thetillhoff/eat/pkg/dhcp_and_tftp"
	"github.com/thetillhoff/eat/pkg/tosca"
)

func Install(serviceTemplate tosca.ServiceTemplate, inputs []string) {
	if debug {
		dhcp_and_tftp.Debug()
	}

	// TODO: Add derivation from NodeType to NodeTemplate, so all properties etc are available.

	serviceTemplate.TopologyTemplate = AddInputsToTopologyTemplate(serviceTemplate.TopologyTemplate, inputs)

	// TODO: retrieve MAC-addresses as inputs
	var macAddresses []string = []string{"AA:11:22:33:44:55"}

	// Build <live-os>.iso - Can be ommited for now and a fixed path of a precreated iso can be selected.
	// var liveIsoPath string = "debian-live.iso"

	// Start dhcp-, tftp- and http-servers
	//   The dhcp-server is in proxy mode, which means an active dhcp is required on the network
	dhcp_and_tftp.Start()

	// When a new machine attempts to boot over pxe, it ...
	// 1. gets a random ip-address from the preexisting dhcp
	// 2. loads a fitting (bios or efi) variant of ipxe from the tftp-server
	// 3. requests the ipxe configuration from the http-server at /default - This contains the location of the <live-os>.iso
	// 4. loads the <live-os>.iso from the http-server at /isos/<live-os>.iso and boot it
	// 5. runs the <live-os> with a embedded webserver, which
	//    - provides hardware-information at /cpuinfo.txt and /meminfo.txt
	//    - runs a ssh-server with a preconfigured ssh-key

	// Information about the machines can be retrieved now, requirement/capability matches can be made and fitting nodes are selected
	for _, macAddress := range macAddresses {
		arp.GetIpForMac(macAddress) // Wait until mac is listed in arp table, then return mapped ip-address
		// For all machines in parallel:
		//   - Either shutdown if not a selected node
		//   - Or configure if a selected node
	}

	// TODO: Add installation here

	serviceTemplate.TopologyTemplate = AddOutputsOfTopologyTemplate(serviceTemplate.TopologyTemplate)

	fmt.Println(serviceTemplate.ToString())

}
