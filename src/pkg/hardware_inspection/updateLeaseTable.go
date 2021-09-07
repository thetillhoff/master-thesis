package hardware_inspection

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func updateLeaseTable() {

	// Get content of leasefile
	byteOutput, err := ioutil.ReadFile("dnsmasq/leases") // just pass the file name
	if err != nil {
		log.Fatalln("ERR Can't read dnsmasq/leases file:", err)
	}

	// Parse output of arp command
	for lineNumber, line := range strings.Split(string(byteOutput), "\n") {
		// lines look like:
		//   1630961876 <MAC with colon> <IP> * <MAC of DHCP?>
		//  ^timestamp^

		var words []string = strings.Split(line, " ")

		// Validate line
		if len(words) == 1 { // If line is empty, it is fine // TODO No idea why the length is then measured as 1 instead of 0...
			continue
		} else if len(words) != 5 { // If line is invalid
			log.Fatalln("ERR Couldn't parse leases, invalid length of words at line "+strconv.Itoa(lineNumber)+". Expected 5, got '"+strconv.Itoa(len(words))+"' ;", words)
		}

		// Sometimes, there are multiple entries for the same MAC; For example when the ipxe-kernel runs and when the live-os runs.
		// Therefore only the newest entry should be considered.

		// Get timestamp
		timestamp, err := strconv.ParseUint(words[0], 0, 64)
		if err != nil {
			log.Fatalln("ERR Couldn't parse timestamp of leases.", err)
		}
		// Get mac-address
		var macAddress string = words[1]
		// Get ip-address
		var ipAddress string = words[2]

		if machine, ok := leaseTable[macAddress]; ok { // If another lease for machine exists
			if debug {
				log.Println("INF Found existing lease with timestamp", machine.timestamp)
			}

			if machine.timestamp < timestamp { // If this lease is newer
				if debug {
					log.Println("INF Overwriting preexisting lease with timestamp", timestamp)
				}
				machine.timestamp = timestamp
				machine.ipAddress = ipAddress
				leaseTable[macAddress] = machine
			}
		} else { // If no other lease for machine exists
			if debug {
				log.Println("INF Adding lease with timestamp", timestamp)
			}
			leaseTable[macAddress] = lease{
				macAddress: macAddress,
				ipAddress:  ipAddress,
				timestamp:  timestamp,
			}
		}
	}
}
