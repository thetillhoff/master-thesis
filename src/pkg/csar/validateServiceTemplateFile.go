package csar

import (
	"log"
	"strings"
)

// According to 4.2.1.2.1 of the tosca spec, the first line of a service template MUST contain the "tosca_definitions_version"
func ValidateServiceTemplateFile(fileContents string) {

	firstline := strings.Split(fileContents, "\n")[0]

	if !strings.HasPrefix(firstline, "tosca_definitions_version: ") {
		log.Fatalln("ERR Invalid service template. First line must define 'tosca_definitions_version'.")
	}

}
