package ssh

import (
	"log"
	"strings"

	"github.com/melbahja/goph"
)

// ssh -i ./id_rsa 192.168.122.168 -l user -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null

func RunCommandOnHost(ipAddress string, command string) string {

	// Start new ssh connection with private key.
	auth, err := goph.Key(privateKeyPath, "")
	if err != nil {
		log.Fatalln("ERR Can't read ssh private-key:", err)
	}

	client, err := goph.NewUnknown("user", ipAddress, auth) // Please note that using NewUnkown skips checking for known hosts and therefore enables man-in-the-middle attacks.
	if err != nil {
		log.Fatalln("ERR Can't create connection to '"+ipAddress+"';", err)
	}

	// Defer closing the network connection.
	defer client.Close()

	// Execute command
	out, err := client.Run(command)

	if err != nil {
		log.Fatalln("ERR Problem while running command on '"+ipAddress+"';", err, string(out))
	}

	// Return your output
	return strings.TrimSpace(string(out))
}
