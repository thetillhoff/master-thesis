package csar

import (
	"log"
	"strings"

	"github.com/thetillhoff/eat/pkg/tosca"
	"gopkg.in/yaml.v3"
)

// Input content of EntryDefinitionsFile and optionally space-seperated list of paths to OtherDefinitions
//
// Returns full TOSCA service template and list of full TOSCA service templates declared in otherDefinitions
//
// NOTE: Actual imports and substitutions happen at runtime
func parseServiceTemplate(relativePath string, otherDefinitionsFilepaths string) tosca.ServiceTemplate {
	var (
		serviceTemplate        tosca.ServiceTemplate
		err                    error
		filepath               string
		serviceTemplateContent string
		otherDefinitionContent string
		otherServiceTemplate   tosca.ServiceTemplate
		otherServiceTemplates  []tosca.ServiceTemplate
	)

	serviceTemplateContent = archiveContents[relativePath]

	if len(serviceTemplateContent) == 0 {
		log.Fatalln("ERR Empty serviceTemplate at '" + relativePath + "' or file not found.")
	}

	// According to 4.2.1.2.1 of the tosca spec, the first line of a service template MUST contain the "tosca_definitions_version"
	firstline := strings.Split(serviceTemplateContent, "\n")[0]
	if !strings.HasPrefix(firstline, "tosca_definitions_version: ") {
		log.Fatalln("ERR Invalid serviceTemplate. First line must define 'tosca_definitions_version'.")
	}

	// Parse EntryDefinitions and add them to serviceTemplate
	err = yaml.Unmarshal([]byte(serviceTemplateContent), &serviceTemplate)
	if err != nil {
		log.Fatalf("ERR Cannot unmarshal data: %v", err)
	}

	// Load contents of otherDefinitions and merge them into serviceTemplate
	// TODO this should resolve the substitions defined in otherDefinitions instead of just merging it
	for _, filepath = range strings.Split(otherDefinitionsFilepaths, " ") {

		otherDefinitionContent = archiveContents[filepath] // load file contents
		err = yaml.Unmarshal([]byte(otherDefinitionContent), &otherServiceTemplate)
		if err != nil {
			log.Fatalf("ERR Cannot unmarshal data: %v", err)
		}

		otherServiceTemplates = append(otherServiceTemplates, otherServiceTemplate)

	}

	// TODO: Implement substitution contained in otherServiceTemplates here. Return only one serviceTemplate later on.
	// otherServiceTemplates contains absolue/relative paths IN the CSAR package.

	// Note that any further TOSCA definitions files required by the definitions specified by Entry-Definitions or Other-Definitions can be found by a TOSCA orchestrator by processing respective imports statements. Note also that artifact files (e.g. scripts, binaries, configuration files) used by the TOSCA definitions and included in the CSAR are fully described and referred via relative path names in artifact definitions in the respective TOSCA definitions files contained in the CSAR.

	return serviceTemplate
}
