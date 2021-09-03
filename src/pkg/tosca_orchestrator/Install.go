package tosca_orchestrator

import (
	"fmt"
	"log"
	"strconv"

	"github.com/thetillhoff/eat/pkg/arp"
	"github.com/thetillhoff/eat/pkg/csar"
	"github.com/thetillhoff/eat/pkg/dhcp_and_tftp"
	"github.com/thetillhoff/eat/pkg/ssh"
	"github.com/thetillhoff/eat/pkg/tosca"
	"github.com/thetillhoff/eat/pkg/wol"
)

func Install(archive csar.CSAR, inputs []string) {
	if debug {
		// Set debug for imports
		dhcp_and_tftp.Debug()
	}

	// Add derivation from NodeType to NodeTemplate, so all properties etc are available in the base type.
	// TODO: This could be merged with InstantiateTpoplogyTemplate - not sure if that makes more sense though.
	archive.ServiceTemplate = DeriveServiceTemplate(archive) // TODO copy from test

	// Fill in the provided inputs
	archive.ServiceTemplate.TopologyTemplate = AddInputsToTopologyTemplate(archive.ServiceTemplate.TopologyTemplate, inputs)          // TODO copy from test
	archive.ServiceTemplate.TopologyTemplate.Inputs["MacAddresses"] = tosca.ParameterDefinition{Value: []string{"AA:11:22:33:44:55"}} // TODO remove demo setup

	// Verify input contains macAddresses and path to ssh-key
	// TODO Verify inputs (like validity of macaddresses and existance of file at provided path to ssh-key)

	// Merge types and templates, fill with default values where needed, run normativeFunctions, verify constraints, ...
	archive.ServiceTemplate.TopologyTemplate = InstantiateTopologyTemplate(archive.ServiceTemplate)

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

	// Parse MacAddresses from Input
	values, err := tosca.ParseList(archive.ServiceTemplate.TopologyTemplate.Inputs["MacAddresses"].Value)
	if err != nil {
		log.Fatalln("ERR Couldn't parse input 'MacAddresses' to slice", err)
	}
	macAddresses := []string{}
	for _, value := range values {
		if macAddress, ok := value.(string); ok {
			macAddresses = append(macAddresses, macAddress)
		} else {
			log.Fatalln("ERR Couldn't parse macAddress to string:", value)
		}
	}

	if debug {
		log.Println("INF " + strconv.Itoa(len(macAddresses)) + " MAC-addresses provided.")
	}

	// Start dhcp-, tftp- and http-servers
	dhcp_and_tftp.Start()

	// Retrieve information about machines
	machines := []Machine{}
	for _, macAddress := range macAddresses {
		machine := Machine{}

		// Wait until mac is listed in arp table, then return mapped ip-address
		wol.Wake(macAddress) // TODO This should check whether the host is already reachable (f.e. with ping)
		log.Println("INF Waked machine with mac '" + macAddress + "'.")

		// Get ip for mac
		machine.IpAddress = arp.GetIpForMac(macAddress)
		log.Println("INF IP-Address retrieved:", machine.IpAddress)

		// Retrieve hardware information
		machine.Ram = GetRam(machine)
		machine.Cores = GetCores(machine)

		machines = append(machines, machine)
	}
	if debug {
		log.Println("INF " + strconv.Itoa(len(machines)) + " machines are available.")
	}

	// Select some servers out of available ones
	// TODO add proper selection process
	selectedMachine := machines[0]

	// For all machines in parallel:
	//   - Either shutdown if not a selected node
	//   - Or configure if a selected node
	// Optionally do something with the not-selected servers (f.e. shut them down)

	// Run installation and configuration scripts on selected hosts (via ssh)
	// TODO Thoughts
	//   Assume a webserver and a database are to be installed - how are the servers selected?
	//   -> One after another, this means first the whole process runs for the webserver, then for the database
	//      Make sure the selected server isn't always the first in the list that fits, but a random one.
	ssh.RunCommandOnHost(selectedMachine.IpAddress, "hostname") // TODO Replace "hostname" with actual command(s) to run

	// TODO
	// After installation succeeded, add a SRV entry in the DNS. This can then later be used for state detection.
	// State detection == is the desired state already reached?
	// When no DNS is active, use dnsmasq as DNSserver. Dnsmasq uses the local /etc/hosts, and SRV entries can be added there.

	// Collect values for outputs specified in topologyTemplate
	archive.ServiceTemplate.TopologyTemplate = AddOutputsOfTopologyTemplate(archive.ServiceTemplate.TopologyTemplate)

	// Print desired outputs of topologyTemplate
	for key, output := range archive.ServiceTemplate.TopologyTemplate.Outputs {
		fmt.Println(key, output)
	}
}
