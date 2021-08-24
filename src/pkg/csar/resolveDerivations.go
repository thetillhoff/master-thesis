package csar

// import (
// 	"fmt"
// 	"log"

// 	"github.com/thetillhoff/eat/pkg/tosca"
// )

// // Writes a complete serviceTemplate into provided CSAR package and returns the whole package.
// //
// // The complete serviceTemplate includes fully derived types and all substitutions defined in otherDefinitions.
// //
// // The archive.ServiceTemplate must already contain a ServiceTemplate.
// // The archive.Imports must already be "converted" to archive.imports.
// func (archive CSAR) resolveDerivations() CSAR {
// 	var (
// 		mainServiceTemplate tosca.ServiceTemplate
// 	)

// 	mainServiceTemplate = tosca.ServiceTemplate{
// 		// ToscaDefinitionsVersion not needed
// 		// Profile                 not needed
// 		// Metadata                not needed
// 		// Description             not needed
// 		// Repositories are already resolved in archive.imports
// 		// Imports are already resolved in archive.imports
// 		ArtifactTypes:     make(map[string]tosca.ArtifactType),
// 		DataTypes:         make(map[string]tosca.DataType),
// 		CapabilityTypes:   make(map[string]tosca.CapabilityType),
// 		InterfaceTypes:    make(map[string]tosca.InterfaceType),
// 		RelationshipTypes: make(map[string]tosca.RelationshipType),
// 		NodeTypes:         make(map[string]tosca.NodeType),
// 		GroupTypes:        make(map[string]tosca.GroupType),
// 		PolicyTypes:       make(map[string]tosca.PolicyType),
// 		// TopologyTemplate   not needed
// 	}

// 	for namespace, importedServiceTemplate := range archive.imports {

// 		if namespace != "" { // If namespace is not unnamed
// 			namespace = namespace + ":" // Add seperator
// 		}

// 		// ArtifactType
// 		for name, importedType := range importedServiceTemplate.ArtifactTypes {
// 			if _, ok := mainServiceTemplate.ArtifactTypes[namespace+name]; ok {
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				if importedType.DerivedFrom != "" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.ArtifactTypes[namespace+name] = importedType
// 			}
// 		}

// 		// DataType
// 		for name, importedType := range importedServiceTemplate.DataTypes {
// 			if _, ok := mainServiceTemplate.DataTypes[namespace+name]; ok {
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				// TODO: add other normative types as well
// 				if importedType.DerivedFrom != "" &&
// 					importedType.DerivedFrom != "string" &&
// 					importedType.DerivedFrom != "integer" &&
// 					importedType.DerivedFrom != "float" &&
// 					importedType.DerivedFrom != "boolean" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.DataTypes[namespace+name] = importedType
// 			}
// 		}

// 		// CapabilityType
// 		for name, importedType := range importedServiceTemplate.CapabilityTypes {
// 			if _, ok := mainServiceTemplate.CapabilityTypes[namespace+name]; ok {
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				if importedType.DerivedFrom != "" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.CapabilityTypes[namespace+name] = importedType
// 			}
// 		}

// 		// InterfaceType
// 		for name, importedType := range importedServiceTemplate.InterfaceTypes {
// 			if _, ok := mainServiceTemplate.InterfaceTypes[namespace+name]; ok {
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				if importedType.DerivedFrom != "" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.InterfaceTypes[namespace+name] = importedType
// 			}
// 		}

// 		// RelationshipType
// 		for name, importedType := range importedServiceTemplate.RelationshipTypes {
// 			if _, ok := mainServiceTemplate.RelationshipTypes[namespace+name]; ok {
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				if importedType.DerivedFrom != "" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.RelationshipTypes[namespace+name] = importedType
// 			}
// 		}

// 		// NodeType
// 		for name, importedType := range importedServiceTemplate.NodeTypes {
// 			if _, ok := mainServiceTemplate.NodeTypes[namespace+name]; ok { // If nodeType with same name and namespace already exists
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				if importedType.DerivedFrom != "" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.NodeTypes[namespace+name] = importedType
// 			}
// 		}

// 		// GroupType
// 		for name, importedType := range importedServiceTemplate.GroupTypes {
// 			if _, ok := mainServiceTemplate.GroupTypes[namespace+name]; ok {
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				if importedType.DerivedFrom != "" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.GroupTypes[namespace+name] = importedType
// 			}
// 		}

// 		// PolicyType
// 		for name, importedType := range importedServiceTemplate.PolicyTypes {
// 			if _, ok := mainServiceTemplate.PolicyTypes[namespace+name]; ok {
// 				log.Fatalln("ERR duplicate key after import:", namespace+name)
// 			} else {
// 				// If imported Type is derived, add namespace as prefix
// 				if importedType.DerivedFrom != "" {
// 					importedType.DerivedFrom = namespace + importedType.DerivedFrom
// 				}
// 				mainServiceTemplate.PolicyTypes[namespace+name] = importedType
// 			}
// 		}
// 	}

// 	// After all types are loaded, run derivation
// 	mainServiceTemplate = mainServiceTemplate.ResolveDerivations()

// 	// ArtifactTypes
// 	for name := range archive.ServiceTemplate.ArtifactTypes {
// 		archive.ServiceTemplate.ArtifactTypes[name] = mainServiceTemplate.ArtifactTypes[name]
// 	}

// 	// DataTypes
// 	for name := range archive.ServiceTemplate.DataTypes {
// 		archive.ServiceTemplate.DataTypes[name] = mainServiceTemplate.DataTypes[name]
// 	}

// 	// CapabilityTypes
// 	for name := range archive.ServiceTemplate.CapabilityTypes {
// 		archive.ServiceTemplate.CapabilityTypes[name] = mainServiceTemplate.CapabilityTypes[name]
// 	}

// 	// InterfaceTypes
// 	for name := range archive.ServiceTemplate.InterfaceTypes {
// 		archive.ServiceTemplate.InterfaceTypes[name] = mainServiceTemplate.InterfaceTypes[name]
// 	}

// 	// RelationshipTypes
// 	for name := range archive.ServiceTemplate.RelationshipTypes {
// 		archive.ServiceTemplate.RelationshipTypes[name] = mainServiceTemplate.RelationshipTypes[name]
// 	}

// 	// NodeTypes
// 	for name := range archive.ServiceTemplate.NodeTypes {
// 		fmt.Println(name)
// 		archive.ServiceTemplate.NodeTypes[name] = mainServiceTemplate.NodeTypes[name]
// 	}

// 	// GroupTypes
// 	for name := range archive.ServiceTemplate.GroupTypes {
// 		archive.ServiceTemplate.GroupTypes[name] = mainServiceTemplate.GroupTypes[name]
// 	}

// 	// PolicyTypes
// 	for name := range archive.ServiceTemplate.PolicyTypes {
// 		archive.ServiceTemplate.PolicyTypes[name] = mainServiceTemplate.PolicyTypes[name]
// 	}

// 	return archive
// }
