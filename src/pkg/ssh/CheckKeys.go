package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var (
	bitSize int = 4096

	privateKeyPath string = "private.key"
	publicKeyPath  string = "public.key"
)

// Taken from https://gist.github.com/devinodaniel/8f9b8a4f31573f428f29ec0e884e6673

// If keys are empty, try to load them; if that fails, generate new ones
// First argument is private-key path
// Second argument is public-key path
func CheckKeys(args ...string) {
	if len(args) == 0 {
		// Do nothing, since thats fine as well
	} else if len(args) == 2 {
		privateKeyPath = args[0]
		publicKeyPath = args[1]
	} else {
		log.Println("ERR Invalid number of arguments for LoadKeys.")
	}

	// Generate keys when needed, load them otherwise
	if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) { // Private-key doesn't exist
		// Test if public key exists
		if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) { // Public-key doesn't exist
			generateKeys()
		} else { // Public-key does exist
			log.Fatalln("ERR Private key doesn't exist, but public key does.")
		}
	} else { // Private-key does exist
		// Test if public key exists
		if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) { // Public-key doesn't exist
			log.Println("WRN Private key exists, but public key doesn't.")
		} else { // Public key doesn't exist
			// Do nothing, since this is fine
		}
	}
}

func generateKeys() {
	var (
		privateKey rsa.PrivateKey
	)

	privateKey = generatePrivateKey()
	generatePublicKey(privateKey)

	if debug {
		log.Println("INF SSH-keys generated")
	}
}

func generatePrivateKey() rsa.PrivateKey {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		log.Println("ERR Couldn't generate private-key:", err)
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		log.Println("ERR Couldn't validate private-key:", err)
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(privateKeyPath)
	defer privateKeyFile.Close()
	if err != nil {
		log.Println("ERR Couldn't create private-key-file:", err)
	}
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		log.Println("ERR Couldn't save private-key:", err)
	}

	return *privateKey
}

func generatePublicKey(privateKey rsa.PrivateKey) {
	// generate and write public key
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalln("ERR Couldn't generate public-key:", err)
	}
	err = ioutil.WriteFile(publicKeyPath, ssh.MarshalAuthorizedKey(publicKey), 0655)
	if err != nil {
		log.Fatalln("ERR Couldn't save public-key:", err)
	}
}
