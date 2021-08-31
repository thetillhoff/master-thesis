package ssh

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

func RunCommandOnHost(ipAddress string, command string) string {

	// Start new ssh connection with private key.
	auth, err := goph.Key(privateKeyPath, "")
	if err != nil {
		log.Fatalln("ERR Couldn't read ssh private-key:", err)
	}

	client, err := goph.NewUnknown("root", ipAddress, auth) // Please note that using NewUnkown skips checking for known hosts and therefore enables man-in-the-middle attacks.
	if err != nil {
		log.Fatalln("ERR Couldn't create connection to '"+ipAddress+"';", err)
	}

	// Defer closing the network connection.
	defer client.Close()

	// Execute command
	out, err := client.Run(command)

	if err != nil {
		log.Fatalln("ERR Problem while running command on '"+ipAddress+"';", err)
	}

	// Return your output
	return string(out)
}

// "Stolen" from https://skarlso.github.io/2019/02/17/go-ssh-with-host-key-verification/
func runCommandOnHost2(ipAddress string, command string) string {

	user := "root"
	command = "uptime" // TODO: remove
	port := "22"

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("ERR Couldn't retrieve users home-dir for ssh-keys:", err)
	}

	key, err := ioutil.ReadFile(path.Join(dirname, ".ssh/id_rsa"))
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Add in password check here for moar security.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", ipAddress+":"+port, config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()
	ss, err := client.NewSession()
	if err != nil {
		log.Fatal("unable to create SSH session: ", err)
	}
	defer ss.Close()
	// Creating the buffer which will hold the remotly executed command's output.
	var stdoutBuf bytes.Buffer
	ss.Stdout = &stdoutBuf
	ss.Run(command)

	return stdoutBuf.String()
}
