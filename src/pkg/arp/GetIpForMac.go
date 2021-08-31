package arp

import (
	"log"
	"runtime"
	"strconv"
	"strings"

	"github.com/thetillhoff/eat/pkg/script_runner"
)

var (
	// Contains a map of MAC to IP, where the MAC address acts as key
	// Separator is colon!
	arpTable map[string]string = make(map[string]string)
)

func GetIpForMac(macAddress string) string {

	log.Println("INF Waiting for mac-address '" + macAddress + "' to come online...")

	for !isMacInArp(macAddress) {
		switch runtime.GOOS {
		case "windows":
			updateArpWindows()
		case "linux":
			updateArpLinux()
		default:
			log.Fatalln("ERR Cannot wait for MAC on this OS.")
		}
	}
	return arpTable[macAddress]
}

func isMacInArp(macAddress string) bool {
	for key := range arpTable {
		if key == macAddress {
			return true
		}
	}
	return false
}

func updateArpLinux() {
	var (
		output string
	)

	// Clear previous entries
	for k := range arpTable {
		delete(arpTable, k)
	}

	// Get output of arp table
	output = script_runner.RunLinuxCommand("arp -a")

	// Parse output of arp command
	for _, line := range strings.Split(output, "\n") {
		// lines look like:
		//   <NAME> (<IP>) at <MAC with colon> [ether] on eth0

		var words []string = strings.Split(line, " ")

		// Validate line
		if len(words) == 1 {
			// If line is empty (last line), skip it - unsure why this has length of 1
			continue
		} else if len(words) != 7 {
			log.Fatalln("ERR Couldn't parse arp table, invalid length of words. Expected 7, got '"+strconv.Itoa(len(words))+"' :", words)
		}

		// Get mac-address
		var macAddress string = words[3]

		// Get ip-address and remove leading and trailing brackets
		var ipAddress string = strings.Trim(words[1], "()")
		arpTable[macAddress] = ipAddress
	}
}

func updateArpWindows() {
	var (
		output string
	)

	// Clear previous entries
	for k := range arpTable {
		delete(arpTable, k)
	}

	output = script_runner.RunWindowsCommand("arp -a")

	// Parse output of arp command
	for _, line := range strings.Split(output, "\n") {
		// lines look like:
		//   Internet Address      Physical Address                         Type
		//   <IP>                  <MAC with dash>                          dynamic
		//   <broadcast IP>        <MAC with dash or ff-ff-ff-ff-ff-ff>     static

		// Remove leading space from lines
		line = strings.TrimSpace(line)

		// Since the arp command on windows contains title lines and other things, only try to parse lines with proper mappings
		if strings.HasSuffix(line, "dynamic") || strings.HasSuffix(line, "static") {
			var words []string = strings.Split(line, " ")

			// Validate line
			if len(words) != 3 {
				log.Fatalln("ERR Couldn't parse arp table, invalid length of words. Expected 3, got '"+strconv.Itoa(len(words))+"' :", words)
			}

			// Get mac-address and replace dashes with colons
			var macAddress string = strings.ReplaceAll(words[1], "-", ":")

			// Get ip-address
			var ipAddress string = words[0]

			// Add only if not a broadcast address
			if macAddress != "ff-ff-ff-ff-ff-ff" {
				arpTable[macAddress] = ipAddress
			}
		}
	}
}
