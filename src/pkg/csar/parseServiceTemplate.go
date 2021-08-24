package csar

import (
	"log"
	"strings"

	"github.com/thetillhoff/eat/pkg/tosca"
	"gopkg.in/yaml.v3"
)

// Input path to serviceTemplate and optionally space-seperated list of paths to OtherDefinitions (==serviceTemplates with substitutions).
//
// Returns full TOSCA service template, with completed substitutions.
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

	// Parse serviceTemplateContents
	err = yaml.Unmarshal([]byte(serviceTemplateContent), &serviceTemplate)
	if err != nil {
		log.Fatalf("ERR Cannot unmarshal data: %v", err)
	}

	// Initialize maps even if empty
	if len(serviceTemplate.ArtifactTypes) == 0 {
		serviceTemplate.ArtifactTypes = make(map[string]tosca.ArtifactType)
	}
	if len(serviceTemplate.DataTypes) == 0 {
		serviceTemplate.DataTypes = make(map[string]tosca.DataType)
	}
	if len(serviceTemplate.CapabilityTypes) == 0 {
		serviceTemplate.CapabilityTypes = make(map[string]tosca.CapabilityType)
	}
	if len(serviceTemplate.InterfaceTypes) == 0 {
		serviceTemplate.InterfaceTypes = make(map[string]tosca.InterfaceType)
	}
	if len(serviceTemplate.RelationshipTypes) == 0 {
		serviceTemplate.RelationshipTypes = make(map[string]tosca.RelationshipType)
	}
	if len(serviceTemplate.NodeTypes) == 0 {
		serviceTemplate.NodeTypes = make(map[string]tosca.NodeType)
	}
	if len(serviceTemplate.GroupTypes) == 0 {
		serviceTemplate.GroupTypes = make(map[string]tosca.GroupType)
	}
	if len(serviceTemplate.PolicyTypes) == 0 {
		serviceTemplate.PolicyTypes = make(map[string]tosca.PolicyType)
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

	if debug {
		var totalTypes int = len(serviceTemplate.ArtifactTypes) +
			len(serviceTemplate.DataTypes) +
			len(serviceTemplate.CapabilityTypes) +
			len(serviceTemplate.InterfaceTypes) +
			len(serviceTemplate.RelationshipTypes) +
			len(serviceTemplate.NodeTypes) +
			len(serviceTemplate.GroupTypes) +
			len(serviceTemplate.PolicyTypes)
		log.Println("INF Parsed ServiceTemplate from '"+relativePath+"' contained", totalTypes, "own Types.")
	}

	return serviceTemplate
}
