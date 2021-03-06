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
	if serviceTemplate.ArtifactTypes == nil {
		serviceTemplate.ArtifactTypes = make(map[string]tosca.ArtifactType)
	}
	if serviceTemplate.DataTypes == nil {
		serviceTemplate.DataTypes = make(map[string]tosca.DataType)
	}
	if serviceTemplate.CapabilityTypes == nil {
		serviceTemplate.CapabilityTypes = make(map[string]tosca.CapabilityType)
	}
	if serviceTemplate.InterfaceTypes == nil {
		serviceTemplate.InterfaceTypes = make(map[string]tosca.InterfaceType)
	}
	if serviceTemplate.RelationshipTypes == nil {
		serviceTemplate.RelationshipTypes = make(map[string]tosca.RelationshipType)
	}
	if serviceTemplate.NodeTypes == nil {
		serviceTemplate.NodeTypes = make(map[string]tosca.NodeType)
	}
	if serviceTemplate.GroupTypes == nil {
		serviceTemplate.GroupTypes = make(map[string]tosca.GroupType)
	}
	if serviceTemplate.PolicyTypes == nil {
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
