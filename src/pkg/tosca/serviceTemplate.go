package tosca

type ServiceTemplate struct { // [tosca spec 2.0, 4.2.1.1]

	// [mandatory] Defines the version of the TOSCA specification the template (grammar) complies with.
	ToscaDefinitionsVersion string `yaml:"tosca_definitions_version" json:"tosca_definitions_version"`

	// The optional profile name that can be used by other TOSCA service templates to import the type definitions in this document.
	Profile string `yaml:"profile,omitempty" json:"profile,omitempty"`

	// Defines a section used to declare additional metadata information. Domain-specific TOSCA profile specifications may define keynames that are mandatory for their implementations. Recognized keynames are "template_name", "template_author" and "template_version". Name and version should be used to identify during lifecycle management.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`

	// Declares a description for this Service Template and its contents.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// DslDefinitions declares optional DSL-specific definitions and conventions.  For example, in YAML, this allows defining reusable YAML macros (i.e., YAML alias anchors) for use throughout the TOSCA Service Template.

	// Declares the map of external repositories which contain artifacts that are referenced in the service template along with their addresses used to connect to them in order to retrieve the artifacts.
	Repositories map[string]RepositoryDefinition `yaml:"repositories,omitempty" json:"repositories,omitempty"`

	// Declares a list import statements pointing to external TOSCA Definitions documents. For example, these may be file location or URIs relative to the service template file within the same TOSCA CSAR file.
	Imports []ImportDefinition `yaml:"imports,omitempty" json:"imports,omitempty"`

	// This section contains an optional map of artifact type definitions for use in the service template.
	ArtifactTypes map[string]ArtifactType `yaml:"artifact_types,omitempty" json:"artifact_types,omitempty"`

	// Declares a map of optional TOSCA Data Type definitions.
	DataTypes map[string]DataType `yaml:"data_types,omitempty" json:"data_types,omitempty"`

	// This section contains an optional map of capability type definitions for use in the service template.
	CapabilityTypes map[string]CapabilityType `yaml:"capability_types,omitempty" json:"capability_types,omitempty"`

	// This section contains an optional map of interface type definitions for use in the service template.
	InterfaceTypes map[string]InterfaceType `yaml:"interface_types,omitempty" json:"interface_types,omitempty"`

	// This section contains a map of relationship type definitions for use in the service template.
	RelationshipTypes map[string]RelationshipType `yaml:"relationship_types,omitempty" json:"relationship_types,omitempty"`

	// This section contains a map of node type definitions for use in the service template.
	NodeTypes map[string]NodeType `yaml:"node_types,omitempty" json:"node_types,omitempty"`

	// This section contains a map of group type definitions for use in the service template.
	GroupTypes map[string]GroupType `yaml:"group_types,omitempty" json:"group_types,omitempty"`

	// This section contains a list of policy type definitions for use in the service template.
	PolicyTypes map[string]PolicyType `yaml:"policy_types,omitempty" json:"policy_types,omitempty"`

	// Defines the topology template of an application or service, consisting of node templates that represent the application’s or service’s components, as well as relationship templates representing relations between the components.
	TopologyTemplate TopologyTemplateDefinition `yaml:"topology_template,omitempty" json:"topology_template,omitempty"`
}

// func (serviceTemplate ServiceTemplate) getParentType(name string) (error, interface{}) {
// 	var (
// 		err     error
// 		anyType interface{}
// 	)
// 	err, anyType = serviceTemplate.getArtifactType(name)
// 	if err == nil {
// 		return nil, anyType
// 	}

// 	return errors.New(""), anyType
// }

// // Return the ArtifactType with name <name> from current serviceTemplate
// func (serviceTemplate ServiceTemplate) getArtifactType(name string) (ArtifactType, error) {
// 	if value, ok := serviceTemplate.ArtifactTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return ArtifactType{}, errors.New("no ArtifactType with provided name exists in serviceTemplate")
// 	}
// }

// // Returns the DataType with name <name> from current serviceTemplate
// func (serviceTemplate ServiceTemplate) getDataType(name string) (DataType, error) {
// 	if value, ok := serviceTemplate.DataTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return DataType{}, errors.New("no DataType with provided name exists in serviceTemplate")
// 	}
// }

// // Returns the CapabilityType with name <name> from the current serviceTemplate
// func (serviceTemplate ServiceTemplate) getCapabilityType(name string) (CapabilityType, error) {
// 	if value, ok := serviceTemplate.CapabilityTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return CapabilityType{}, errors.New("no CapabilityType with provided name exists in serviceTemplate")
// 	}
// }

// // Returns the InterfaceType with name <name> from the current serviceTemplate
// func (serviceTemplate ServiceTemplate) getInterfaceType(name string) (InterfaceType, error) {
// 	if value, ok := serviceTemplate.InterfaceTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return InterfaceType{}, errors.New("no InterfaceType with provided name exists in serviceTemplate")
// 	}
// }

// // Returns the RelationshipType with name <name> from the current serviceTemplate
// func (serviceTemplate ServiceTemplate) getRelationshipType(name string) (RelationshipType, error) {
// 	if value, ok := serviceTemplate.RelationshipTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return RelationshipType{}, errors.New("no RelationshipType with provided name exists in serviceTemplate")
// 	}
// }

// // Return the NodeType with name <name> from current serviceTemplate
// func (serviceTemplate ServiceTemplate) getNodeType(name string) (NodeType, error) {
// 	if value, ok := serviceTemplate.NodeTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return NodeType{}, errors.New("no NodeType with provided name exists in serviceTemplate")
// 	}
// }

// // Returns the GroupType with name <name> from the current serviceTemplate
// func (serviceTemplate ServiceTemplate) getGroupType(name string) (GroupType, error) {
// 	if value, ok := serviceTemplate.GroupTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return GroupType{}, errors.New("no GroupType with provided name exists in serviceTemplate")
// 	}
// }

// // Returns the PolicyType with name <name> from the current serviceTemplate
// func (serviceTemplate ServiceTemplate) getPolicyType(name string) (PolicyType, error) {
// 	if value, ok := serviceTemplate.PolicyTypes[name]; ok {
// 		return value, nil // Fails if not in same servicetemplate -> TODO first search local service template, then imported ones / or conditional with ':'
// 	} else {
// 		return PolicyType{}, errors.New("no PolicyType with provided name exists in serviceTemplate")
// 	}
// }

// Loads all imports into respective namespaces
func (serviceTemplate ServiceTemplate) ResolveImports() { //TODO
	// Check all types whether they reference types from imports -> if one is found, add them to a list
	// Load necessary imports (according to referenced namespaces, if no namespace is set, load by default)
	// Find type, check if it references another type from the import or an import (recursion starts here)
	// add all those types to typelist of <serviceTemplate>
}
