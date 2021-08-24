package csar

import (
	"log"
	"path"

	"github.com/thetillhoff/eat/pkg/tosca"
)

// Takes ServiceTemplatePath (relative within CSAR)
//
// Returns ServiceTemplates fully resolved types (==resolved derivations, including the ones from imports).
func loadServiceTemplate(serviceTemplatePath string, otherDefinitions string) tosca.ServiceTemplate {
	var (
		serviceTemplate tosca.ServiceTemplate

		importPath              string
		importedServiceTemplate tosca.ServiceTemplate
	)

	// parse Service Template at path into serviceTemplate variable, so imports and types are filled
	serviceTemplate = parseServiceTemplate(serviceTemplatePath, otherDefinitions)

	// for each importDefinition of serviceTemplate.imports:
	for _, importDefinition := range serviceTemplate.Imports {

		if importDefinition.Url == "" {
			log.Fatalln("ERR Imports can currently only be defined with Url. Other fields are ignored.")
		}

		// create path to should-be-imported serviceTemplate
		// path.Dir is required, since path.Join does something like Join("a/b/c.yaml","../y/z.yaml")=>"a/b/y/z.yaml" -> This means the filename has to be removed before joining.
		importPath = path.Join(path.Dir(serviceTemplatePath), importDefinition.Url)

		// load ServiceTemplate of importDefinition (recursion) into importedServiceTemplate -> might contain types like "someNodeType" but also "chainedImport:someNodeType"
		// otherDefinitions could also be meant to substitute nodes in imports. Therefore it has to be passed in recursion
		importedServiceTemplate = loadServiceTemplate(importPath, otherDefinitions)

		// if namespace of importDefinition is named
		if importDefinition.Namespace != "" {
			// Add colon to named namespaces only
			importDefinition.Namespace = importDefinition.Namespace + ":"
		}

		// ArtifactTypes
		for typeName, typeDefinition := range importedServiceTemplate.ArtifactTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.ArtifactTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.ArtifactTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}

		// DataTypes
		for typeName, typeDefinition := range importedServiceTemplate.DataTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.DataTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.DataTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}

		// CapabilityTypes
		for typeName, typeDefinition := range importedServiceTemplate.CapabilityTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.CapabilityTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.CapabilityTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}

		// InterfaceTypes
		for typeName, typeDefinition := range importedServiceTemplate.InterfaceTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.InterfaceTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.InterfaceTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}

		// RelationshipTypes
		for typeName, typeDefinition := range importedServiceTemplate.RelationshipTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.RelationshipTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.RelationshipTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}

		// NodeTypes
		for typeName, typeDefinition := range importedServiceTemplate.NodeTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.NodeTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.NodeTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}

		// GroupTypes
		for typeName, typeDefinition := range importedServiceTemplate.GroupTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.GroupTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.GroupTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}

		// PolicyTypes
		for typeName, typeDefinition := range importedServiceTemplate.PolicyTypes {
			// if type with namespace+name already exists in serviceTemplate:
			if _, ok := serviceTemplate.PolicyTypes[importDefinition.Namespace+typeName]; ok {
				// fail because of collision
				log.Fatalln("ERR Type collision detected. Duplicate name is '" + importDefinition.Namespace + typeName + "'.")
			} else {
				// add type of importedServiceTemplate as importDefinition.namespace+typeName to serviceTemplate
				serviceTemplate.PolicyTypes[importDefinition.Namespace+typeName] = typeDefinition
			}
		}
	}

	if debug {
		log.Println("INF Imported ServiceTemplate at '"+serviceTemplatePath+"' contains", len(serviceTemplate.NodeTypes), "NodeTypes.")
	}

	// now all imported types exist (with correct namespace in name) in serviceTemplate; next comes the derivation of types (== resolving derivedFrom).
	// loops below and above can't be merged since types might depend on other imports.
	serviceTemplate = serviceTemplate.ResolveTypeDerivations()

	// return serviceTemplate
	return serviceTemplate
}
