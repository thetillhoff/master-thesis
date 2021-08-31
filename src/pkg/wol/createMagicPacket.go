package wol

import (
	"log"
	"net"
	"regexp"
)

var (
	// Define globals for MacAddress parsing
	delims    = ":-"
	regex_mac = regexp.MustCompile(`^([0-9a-fA-F]{2}[` + delims + `]){5}([0-9a-fA-F]{2})$`)
)

// Takes MAC-address as string, creates magicpacket for it.
func createMagicPacket(mac string) MagicPacket {

	var (
		magicPacket MagicPacket
		macAddress  MacAddress
	)

	// We only support 6 byte MAC addresses
	if !regex_mac.MatchString(mac) {
		log.Fatalln("ERR Provided mac address is not valid.")
	}

	hardwareAddress, err := net.ParseMAC(mac)
	if err != nil {
		log.Fatalln("ERR Couldn't parse mac '"+mac+"';", err)
	}

	// Copy bytes from the returned HardwareAddr -> a fixed size MACAddress
	for i := range macAddress {
		macAddress[i] = hardwareAddress[i]
	}

	// Setup the header which is 6 repetitions of 0xFF
	for i := range magicPacket.header {
		magicPacket.header[i] = 0xFF
	}

	// Setup the payload which is 16 repetitions of the MAC addr
	for i := range magicPacket.payload {
		magicPacket.payload[i] = macAddress
	}

	return magicPacket
}
