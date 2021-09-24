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

	privateKeyPath string = "id_rsa"
	publicKeyPath  string = "id_rsa.pub"
)

// Taken from https://gist.github.com/devinodaniel/8f9b8a4f31573f428f29ec0e884e6673

// If keys are empty, try to load them; if that fails, generate new ones
func CheckKeys() {
	var (
		privateKeyExists bool = true
		publicKeyExists  bool = true
	)

	// Check whether the keys exist
	if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
		privateKeyExists = false
	}
	if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) {
		publicKeyExists = false
	}

	if privateKeyExists && !publicKeyExists {
		log.Fatalln("ERR Public key is missing.")
	} else if !privateKeyExists && publicKeyExists {
		log.Fatalln("ERR Private key is missing.")
	} else if privateKeyExists && publicKeyExists {
		return // If both keys exist, do nothing
	}

	generateKeys()
}

func generateKeys() {
	var (
		privateKey rsa.PrivateKey
	)

	privateKey = generatePrivateKey()
	generatePublicKey(privateKey)

	if debug {
		log.Println("SUC Generated SSH-keys.")
	}
}

func generatePrivateKey() rsa.PrivateKey {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		log.Println("ERR Can't generate private-key:", err)
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		log.Println("ERR Can't validate private-key:", err)
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(privateKeyPath)
	defer privateKeyFile.Close()
	if err != nil {
		log.Println("ERR Can't create private-key-file:", err)
	}
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		log.Println("ERR Can't save private-key:", err)
	}

	return *privateKey
}

func generatePublicKey(privateKey rsa.PrivateKey) {
	// generate and write public key
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalln("ERR Can't generate public-key:", err)
	}
	err = ioutil.WriteFile(publicKeyPath, ssh.MarshalAuthorizedKey(publicKey), 0655)
	if err != nil {
		log.Fatalln("ERR Can't save public-key:", err)
	}
}
