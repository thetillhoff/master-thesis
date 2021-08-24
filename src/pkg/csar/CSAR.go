package csar

import (
	"log"

	"github.com/thetillhoff/eat/pkg/tosca"
	"gopkg.in/yaml.v3"
)

type CSAR struct {
	// This is the version number of the CSAR specification. It defines the structure of the CSAR and the format of the TOSCA.meta file. The value MUST be "2.0" for this version of the CSAR specification.
	CsarVersion string `yaml:"CSAR-Version,omitempty" json:"CSAR-Version,omitempty"`

	// The person or organization that created the CSAR.
	CreatedBy string `yaml:"Created-By,omitempty" json:"Created-By,omitempty"`

	// This references the TOSCA definitions file that SHOULD be used as entry point for processing the contents of the CSAR (e.g. the main TOSCA service template).
	EntryDefinition string `yaml:"Entry-Definitions,omitempty" json:"Entry-Definitions,omitempty"`

	// This references an unambiguous set of files containing substitution templates that can be used to implement nodes defined in the main template (i.e. the file declared in Entry-Definitions). Thus, all the topology templates defined in files listed under the Other-Definitions key are to be used only as substitution templates, and not as standalone services. If such a topology template cannot act as a substitution template, it will be ignored by the orchestrator. The value of the Other-Definitions key is a string containing a list of filenames (relative to the root of the CSAR archive) delimited by a blank space. If the filenames contain blank spaces, the filename should be enclosed by double quotation marks (")
	OtherDefinitions string `yaml:"Other-Definitions,omitempty" json:"Other-Definitions,omitempty"`

	ServiceTemplate tosca.ServiceTemplate // Holds contents of main serviceTemplate
}

// Used when loading metadata files from CSAR archives
func unmarshalMetadata(metadata string) CSAR {
	var (
		archive CSAR
		err     error
	)
	err = yaml.Unmarshal([]byte(metadata), &archive)
	if err != nil {
		log.Fatalf("ERR Cannot unmarshal data: %v", err)
	}

	return archive
}

// Is this needed somewhere?
// func UnmarshalServiceTemplate(yamlContent string) tosca.ServiceTemplate {
// 	var (
// 		serviceTemplate tosca.ServiceTemplate
// 		err             error
// 	)
// 	err = yaml.Unmarshal([]byte(yamlContent), &serviceTemplate)
// 	if err != nil {
// 		log.Fatalf("ERR Cannot unmarshal data: %v", err)
// 	}

// 	return serviceTemplate
// }
