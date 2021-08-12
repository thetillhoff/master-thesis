package tosca

type ServiceTemplate struct { // [tosca spec 2.0, 4.2.1.1]
	ToscaDefinitionsVersion string            `yaml:"tosca_definitions_version" json:"tosca_definitions_version"` // [mandatory] Defines the version of the TOSCA specification the template (grammar) complies with.
	Profile                 string            `yaml:"profile,omitempty" json:"profile,omitempty"`                 // The optional profile name that can be used by other TOSCA service templates to import the type definitions in this document.
	Metadata                map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`               // Defines a section used to declare additional metadata information. Domain-specific TOSCA profile specifications may define keynames that are mandatory for their implementations. Recognized keynames are "template_name", "template_author" and "template_version". Name and version should be used to identify during lifecycle management.
	Description             string            `yaml:"description,omitempty" json:"description,omitempty"`         // Declares a description for this Service Template and its contents.
	// dslDefinitions Declares optional DSL-specific definitions and conventions.  For example, in YAML, this allows defining reusable YAML macros (i.e., YAML alias anchors) for use throughout the TOSCA Service Template.
	Repositories      map[string]RepositoryDefinition `yaml:"repositories,omitempty" json:"repositories,omitempty"`             // Declares the map of external repositories which contain artifacts that are referenced in the service template along with their addresses used to connect to them in order to retrieve the artifacts.
	Imports           []ImportDefinition              `yaml:"imports,omitempty" json:"imports,omitempty"`                       // Declares a list import statements pointing to external TOSCA Definitions documents. For example, these may be file location or URIs relative to the service template file within the same TOSCA CSAR file.
	ArtifactTypes     map[string]ArtifactType         `yaml:"artifact_types,omitempty" json:"artifact_types,omitempty"`         // This section contains an optional map of artifact type definitions for use in the service template.
	DataTypes         map[string]DataType             `yaml:"data_types,omitempty" json:"data_types,omitempty"`                 // Declares a map of optional TOSCA Data Type definitions.
	CapabilityTypes   map[string]CapabilityType       `yaml:"capability_types,omitempty" json:"capability_types,omitempty"`     // This section contains an optional map of capability type definitions for use in the service template.
	InterfaceTypes    map[string]InterfaceType        `yaml:"interface_types,omitempty" json:"interface_types,omitempty"`       // This section contains an optional map of interface type definitions for use in the service template.
	RelationshipTypes map[string]RelationshipType     `yaml:"relationship_types,omitempty" json:"relationship_types,omitempty"` // This section contains a map of relationship type definitions for use in the service template.
	NodeTypes         map[string]NodeType             `yaml:"node_types,omitempty" json:"node_types,omitempty"`                 // This section contains a map of node type definitions for use in the service template.
	GroupTypes        map[string]GroupType            `yaml:"group_types,omitempty" json:"group_types,omitempty"`               // This section contains a map of group type definitions for use in the service template.
	PolicyTypes       map[string]PolicyType           `yaml:"policy_types,omitempty" json:"policy_types,omitempty"`             // This section contains a list of policy type definitions for use in the service template.
	TopologyTemplate  TopologyTemplateDefinition      `yaml:"topology_template,omitempty" json:"topology_template,omitempty"`   // Defines the topology template of an application or service, consisting of node templates that represent the application’s or service’s components, as well as relationship templates representing relations between the components.

}
