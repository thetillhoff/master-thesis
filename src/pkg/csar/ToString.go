package csar

import (
	"bytes"
	"log"

	"gopkg.in/yaml.v3"
)

func (archive CSAR) ToString() string {
	var (
		err         error
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

	return buffer.String()
}
