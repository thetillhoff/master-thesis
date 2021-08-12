package csar

import (
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/thetillhoff/eat/pkg/tosca"
	"gopkg.in/yaml.v3"
)

var (
	folderContents map[string]string
)

// Reads the CSAR at the provided location
//
// Returns full CSAR object.
func LoadFromFile(filePath string) CSAR {
	log.Fatalln("ERR // TODO: Implement LoadFromFile")
	return CSAR{}
}

// Reads the folder-contents at the provided location and treats them according to CSAR specification.
//
// Returns full CSAR object.
func LoadFromFolder(folderPath string) CSAR {
	var (
		elementPath    string
		elementContent string
		archive        CSAR
	)

	folderPath = path.Clean(folderPath) // Cleaning for further usage
	folderContents = loadFolderContents(folderPath)

	if folderContents["TOSCA.meta"] != "" { // If metadata is located at root of CSAR
		archive = UnmarshalMetadata(folderContents["TOSCA.meta"])
		archive.ServiceTemplate, archive.otherServiceTemplates = ParseDefinitionContents(elementContent, archive.OtherDefinitions)
	} else if folderContents["TOSCA-Metadata/TOSCA.meta"] != "" { // If metadata is located in dedicated metadata subdirectory
		archive = UnmarshalMetadata(folderContents["TOSCA-Metadata/TOSCA.meta"])
		archive.ServiceTemplate, archive.otherServiceTemplates = ParseDefinitionContents(elementContent, archive.OtherDefinitions)
	} else { // If only one yaml-file exists at root of CSAR assume metadata is embedded in that file
		archive.EntryDefinition = "" // Initialize

		if len(folderContents) == 0 {
			log.Fatalln("ERR Folder contains no files.")
		}

		for elementPath, elementContent = range folderContents {
			if Debug {
				log.Println("INF Checking file at '" + elementPath + "'.")
			}
			if (path.Ext(elementPath) == ".yaml" || path.Ext(elementPath) == ".yml") && path.Dir(elementPath) == folderPath { // Else if yaml-file exists at root of CSAR (".yaml" OR ".yml")
				if Debug {
					log.Println("INF Entry-file detected at '" + elementPath + "'.")
				}
				if archive.EntryDefinition != "" { // If another EntryDefinition was already detected == If another yaml-file exists at root of CSAR
					log.Println("ERR Invalid CSAR file. No dedicated metadata and ambiguous entry-files detected.")
					if Debug {
						log.Println(archive.EntryDefinition)
						log.Println(elementPath)
						log.Fatal()
					} else {
						log.Fatal()
					}
				}

				// Recognize file as the CSAR Entry-Definitions file
				archive.EntryDefinition = elementPath

				archive.OtherDefinitions = ""                                                                        // OtherDefinitions: Stays empty; "Note that in a CSAR without TOSCA-metadata it is not possible to unambiguously include definitions for substitution templates as we can have only one topology template defined in a yaml file."
				archive.ServiceTemplate, archive.otherServiceTemplates = ParseDefinitionContents(elementContent, "") // Parse Entry-Definitions file

				// Try to parse metadata out of entry-file.
				archive.CsarVersion = translateToscaDefinitionsVersion(archive.ServiceTemplate.ToscaDefinitionsVersion)
				archive.CreatedBy = archive.ServiceTemplate.Metadata["template_author"] // Not specified in docs, but seems intuitive
			}
		}

		if archive.EntryDefinition == "" { // if no yaml-file was detected at root of CSAR the CSAR is invalid
			log.Fatalln("ERR Invalid CSAR file. No dedicated metadata and no entry-file detected.")
		}
	}

	return archive
}

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
	if Debug {
		log.Println("EntryDefinition:", serviceTemplate)
	}

	// Load contents of otherDefinitions and merge them into serviceTemplate
	for _, filepath = range strings.Split(otherDefinitionsFilepaths, " ") {

		otherDefinitionContent = folderContents[filepath] // load file contents
		err = yaml.Unmarshal([]byte(otherDefinitionContent), &otherServiceTemplate)
		if err != nil {
			log.Fatalf("ERR Cannot unmarshal data: %v", err)
		}

		otherServiceTemplates = append(otherServiceTemplates, otherServiceTemplate)

	}

	// Note that any further TOSCA definitions files required by the definitions specified by Entry-Definitions or Other-Definitions can be found by a TOSCA orchestrator by processing respective imports statements. Note also that artifact files (e.g. scripts, binaries, configuration files) used by the TOSCA definitions and included in the CSAR are fully described and referred via relative path names in artifact definitions in the respective TOSCA definitions files contained in the CSAR.

	return serviceTemplate, otherServiceTemplates
}

// Input value of tosca_definitions_version, f.e. "tosca_2_0"
//
// Returns the extracted version, f.e. "2.0"
func translateToscaDefinitionsVersion(toscaDefinitionsVersion string) string {

	if Debug {
		log.Println("tosca_definitions_version:", toscaDefinitionsVersion)
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
