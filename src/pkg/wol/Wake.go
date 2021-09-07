package wol

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"runtime"
	"strconv"
	"strings"

	"github.com/thetillhoff/eat/pkg/script_runner"
)

func Wake(macAddress string) {
	// TODO This should check whether the host is already reachable (f.e. with ping)
	if isVM(macAddress) {
		wakeVM(macAddress)
	} else {
		wakePhyical(macAddress)
	}
}

// Creates magic packet for provided mac address and sends it.
// Original taken from https://sabhiram.com/development/2015/02/16/sending_wol_packets_with_golang.html, but edited slightly
func wakePhyical(macAddress string) {
	var (
		magicPacket MagicPacket
		buffer      bytes.Buffer
	)

	// Create magic packet for provided mac address
	magicPacket = createMagicPacket(macAddress)

	// Fill our byte buffer with the bytes in our MagicPacket
	binary.Write(&buffer, binary.BigEndian, magicPacket)

	// Get a UDPAddr to send the broadcast to
	udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		log.Fatalln("ERR Unable to get a UDP address for 255.255.255.255:9;", err)
	}

	// Open a UDP connection, and defer its cleanup
	connection, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalln("ERR Unable to dial UDP address for 255.255.255.255:9;", err)
	}
	defer connection.Close()

	// Write the bytes of the MagicPacket to the connection
	bytesWritten, err := connection.Write(buffer.Bytes())
	if err != nil {
		log.Fatalln("ERR Unable to write packet to connection:", err)
	} else if bytesWritten != 102 {
		log.Println("WRN Wrote " + strconv.Itoa(bytesWritten) + " bytes, instead of the expected 102.")
	}

	if debug {
		log.Println("INF Waked up physical machine with MAC '" + macAddress + "'.")
	}
}

// Start a vm with provided mac address
func wakeVM(macAddress string) {
	var (
		output string
	)

	// Since this highly depends on the hypervisor, This needs to be split for windows (hyperV) and linux (kvm).
	switch runtime.GOOS {
	case "windows":
		var command string = "get-vm | foreach-object {if ($_.NetworkAdapters.MacAddress -eq " + strings.ReplaceAll(macAddress, ":", "") + ") {Write-Host $_.Name}}"
		output = script_runner.RunWindowsCommand(command)
		if debug && output != "" {
			log.Println("INF Output from wol-vm-command:")
			log.Print(output)
		}
	case "linux":
		output = script_runner.RunLinuxCommand("echo waking up linux-vm with mac '" + macAddress + "'.")
		if debug && output != "" {
			log.Println("INF Output from wol-vm-command:")
			log.Print(output)
		}
	default:
		log.Fatalln("ERR Waking VMs not implemented for this OS.")
	}

	if debug {
		log.Println("INF Waked up virtual machine with MAC '" + macAddress + "'.")
	}
}

// Checks local hypervisor whether there is a vm with provided mac-address
func isVM(macAddress string) bool {
	// TODO: This should check whether the local hypervisor has a vm with provided mac and return true if yes, else otherwise.
	//       Since the poc is meant to use vms, this is not an explicit requirement.
	return true
}
