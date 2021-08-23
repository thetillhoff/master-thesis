package csar

import (
	"bytes"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

func (archive CSAR) Print() {
	var (
		err         error
		archiveYaml string
		buffer      bytes.Buffer
		yamlEncoder *yaml.Encoder
	)

	yamlEncoder = yaml.NewEncoder(&buffer)
	yamlEncoder.SetIndent(2) // Default is 4 spaces
	err = yamlEncoder.Encode(&archive)
	if err != nil {
		log.Fatalln(err)
	}
	defer yamlEncoder.Close()

	archiveYaml = buffer.String()

	fmt.Println(archiveYaml)
}
