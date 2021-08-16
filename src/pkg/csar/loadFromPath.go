package csar

import (
	"log"
	"path"
	"strings"
)

var (
	archiveContents map[string]string
)

// Input is a path either to a CSAR zip-file or an folder containing an extracted CSAR archive.
//
// Automatically detects, whether the provided path is a folder or a zip-file (raises an error if neither).
//
// Returns extracted CSAR object.
func LoadFromPath(csarPath string) CSAR {
	var (
		elementPath    string
		elementContent string
		archive        CSAR
	)

	csarPath = path.Clean(csarPath) // Cleaning for further usage

	// Detect whether folder or zip-file is provided (only by extension for now)
	if path.Ext(csarPath) == ".zip" {
		if Debug {
			log.Println("INF The provided CSAR archive is in zip format.")
		}
		archiveContents = loadZipContents(csarPath)
		// For zip-variant remove the extension from csarPath (required later for "level"-checking).
		csarPath = strings.TrimSuffix(csarPath, path.Ext(csarPath))
	} else {
		if Debug {
			log.Println("INF The provided CSAR archive is in folder format.")
		}
		archiveContents = loadFolderContents(csarPath)
	}

	if archiveContents["TOSCA.meta"] != "" { // If metadata is located at root of CSAR
		archive = unmarshalMetadata(archiveContents["TOSCA.meta"])
		archive.ServiceTemplate, archive.OtherServiceTemplates = ParseDefinitionContents(elementContent, archive.OtherDefinitions)
	} else if archiveContents["TOSCA-Metadata/TOSCA.meta"] != "" { // If metadata is located in dedicated metadata subdirectory
		archive = unmarshalMetadata(archiveContents["TOSCA-Metadata/TOSCA.meta"])
		archive.ServiceTemplate, archive.OtherServiceTemplates = ParseDefinitionContents(elementContent, archive.OtherDefinitions)
	} else { // If only one yaml-file exists at root of CSAR assume metadata is embedded in that file
		archive.EntryDefinition = "" // Initialize

		if len(archiveContents) == 0 {
			log.Fatalln("ERR CSAR doesn't contain files.")
		}

		for elementPath, elementContent = range archiveContents {
			if Debug {
				log.Println("INF Checking file at '" + elementPath + "'.")
			}
			if (path.Ext(elementPath) == ".yaml" || path.Ext(elementPath) == ".yml") && path.Dir(elementPath) == csarPath { // Else if yaml-file exists at root of CSAR (".yaml" OR ".yml")
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
				archive.ServiceTemplate, archive.OtherServiceTemplates = ParseDefinitionContents(elementContent, "") // Parse Entry-Definitions file

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
