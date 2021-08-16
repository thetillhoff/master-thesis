package csar

import (
	"log"
	"strings"

	"github.com/thetillhoff/eat/pkg/tosca"
	"gopkg.in/yaml.v3"
)

// Input content of EntryDefinitionsFile and space-seperated list of paths to OtherDefinitions
//
// Returns full TOSCA service template and list of full TOSCA service templates declared in otherDefinitions
//
// NOTE: Actual imports and substitutions happen at runtime
func ParseDefinitionContents(entryDefinitionsContent string, otherDefinitionsFilepaths string) (tosca.ServiceTemplate, []tosca.ServiceTemplate) {
	var (
		serviceTemplate        tosca.ServiceTemplate
		err                    error
		filepath               string
		otherDefinitionContent string
		otherServiceTemplate   tosca.ServiceTemplate
		otherServiceTemplates  []tosca.ServiceTemplate
	)

	if len(entryDefinitionsContent) == 0 {
		log.Fatalln("ERR Empty entryDefinition")
	}

	// Parse EntryDefinitions and add them to serviceTemplate
	err = yaml.Unmarshal([]byte(entryDefinitionsContent), &serviceTemplate)
	if err != nil {
		log.Fatalf("ERR Cannot unmarshal data: %v", err)
	}

	// Load contents of otherDefinitions and merge them into serviceTemplate
	for _, filepath = range strings.Split(otherDefinitionsFilepaths, " ") {

		otherDefinitionContent = archiveContents[filepath] // load file contents
		err = yaml.Unmarshal([]byte(otherDefinitionContent), &otherServiceTemplate)
		if err != nil {
			log.Fatalf("ERR Cannot unmarshal data: %v", err)
		}

		otherServiceTemplates = append(otherServiceTemplates, otherServiceTemplate)

	}

	// Note that any further TOSCA definitions files required by the definitions specified by Entry-Definitions or Other-Definitions can be found by a TOSCA orchestrator by processing respective imports statements. Note also that artifact files (e.g. scripts, binaries, configuration files) used by the TOSCA definitions and included in the CSAR are fully described and referred via relative path names in artifact definitions in the respective TOSCA definitions files contained in the CSAR.

	return serviceTemplate, otherServiceTemplates
}
