package csar

import (
	"log"
	"regexp"
	"strings"
)

// Input value of tosca_definitions_version, f.e. "tosca_2_0"
//
// Returns the extracted version, f.e. "2.0"
func translateToscaDefinitionsVersion(toscaDefinitionsVersion string) string {

	if len(toscaDefinitionsVersion) == 0 {
		log.Fatalln("ERR Empty 'tosca_definitions_version' cannot be translated to 'CSAR-Version'.")
	}

	// Replace all '_' with '.'
	toscaDefinitionsVersion = strings.ReplaceAll(toscaDefinitionsVersion, "_", ".")

	// Make a Regex to say we only want numbers and dots
	reg, err := regexp.Compile("[^.0-9]+")
	if err != nil {
		log.Fatalln("ERR 'tosca_definitions_version' could not be parsed to CSAR version.", err)
	}

	// Remove all chars except numbers and '.'
	toscaDefinitionsVersion = reg.ReplaceAllString(toscaDefinitionsVersion, "")

	// Remove potential leading '.'
	toscaDefinitionsVersion = strings.TrimPrefix(toscaDefinitionsVersion, ".")

	// Validate version
	if len(toscaDefinitionsVersion) == 0 {
		log.Fatalln("ERR 'tosca_definitions_version' could not be parsed to CSAR version (length==0).")
	}
	// TODO: additional validation for "<MAJOR_VERSION>.<MINOR_VERSION>" conformity

	// Return remaining string
	return toscaDefinitionsVersion
}
