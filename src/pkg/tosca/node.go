package tosca

import (
	"bytes"
	"log"

	"gopkg.in/yaml.v3"
)

// A Node Type is a reusable entity that defines the type of one or more Node Templates. As such, a Node Type defines the structure of observable properties and attributes, the capabilities and requirements of the node as well as its supported interfaces and the artifacts it uses.
type NodeType struct {
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	AbstractType `yaml:",inline,omitempty" json:",inline,omitempty"`

	Properties   map[string]PropertyDefinition   `yaml:"properties,omitempty" json:"properties,omitempty"`
	Attributes   map[string]AttributeDefinition  `yaml:"attributes,omitempty" json:"attributes,omitempty"`
	Capabilities map[string]CapabilityDefinition `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`
	Requirements []RequirementDefinition        `yaml:"requirements,omitempty" json:"requirements,omitempty"` // TODO should be resolved in sequence
	Interfaces   map[string]InterfaceDefinition  `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`
	Artifacts    map[string]ArtifactDefinition   `yaml:"artifacts,omitempty" json:"artifacts,omitempty"`

	// allowed states
	// - initial (not created, only defined in template)
	// - created
	// - configured
	// - started
	// - error
	// state transitions
	// - creating: initial -> created
	// - configuring -> configured
	// - starting: configured -> stopped
	// - stopping: * -> configured
	// - deleted: * -> deleted & no longer tracked

	// special "states"
	// - substitute == abstract, orchestrator must substitute with appropriate template (placeholder?)
	// - select == mark node as abstract, orchestrator must select a node of this type from its inventory (based on constraints of "node_filter")

	// childs (of instance) [5.6.1 of http://docs.oasis-open.org/tosca/TOSCA-Instance-Model/v1.0/csd01/TOSCA-Instance-Model-v1.0-csd01.html#_Toc500843787]
	// - template (mutually exclusive with type) ; allows navigation to the template used to create the instance
	// - type (mutually exclusive with template) ; used when node instance was not created with a template
	// - properties ; final value used by orchestrator
	// - attributes ; value at the time the attribute was accessed - reflect state of underlying node - may change at any time
	// - capabilities ; final value used by the orchestrator
	//   - name
	//   - properties
	//   - attributes
	// - requirements ; maps source to target - doesn't need to be 1:1
	//   - name
	//   - targets (0-N)
}

func NewNodeType() NodeType {
	return NodeType{
		Properties:   map[string]PropertyDefinition{},
		Attributes:   map[string]AttributeDefinition{},
		Capabilities: map[string]CapabilityDefinition{},
		Interfaces:   map[string]InterfaceDefinition{},
		Artifacts:    map[string]ArtifactDefinition{},
	}
}

func (nodeType NodeType) ToString() string {
	var (
		err         error
		buffer      bytes.Buffer
		yamlEncoder *yaml.Encoder
	)

	yamlEncoder = yaml.NewEncoder(&buffer)
	yamlEncoder.SetIndent(2) // Default is 4 spaces
	err = yamlEncoder.Encode(&nodeType)
	if err != nil {
		log.Fatalln(err)
	}
	defer yamlEncoder.Close()

	return buffer.String()
}

type NodeTemplate struct {
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	// [mandatory] The name of the Node Type the Node Template is based upon.
	NodeType string `yaml:"type" json:"type"`

	// An optional description for the Node Template.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`

	// An optional list of directive values to provide processing instructions to orchestrators and tooling.
	//
	// allowed values are:
	// - "substitute": Marks a node template as abstract and instructs the TOSCA Orchestrator to substitute this node template with an appropriate substituting template.
	// - "selectable": Marks a node template as abstract and instructs the TOSCA Orchestrator to select a node of this type from its inventory (based on constraints specified in the optional node_filter in the node template)
	Directives []string `yaml:"directives,omitempty" json:"directives,omitempty"`

	// An optional map of property value assignments for the Node Template.
	Properties map[string]interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`

	// An optional map of attribute value assignments for the Node Template.
	Attributes map[string]interface{} `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// An optional list of requirement assignments for the Node Template.
	Requirements []RequirementAssignment `yaml:"requirements,omitempty" json:"requirements,omitempty"`

	// An optional map of capability assignments for the Node Template.
	Capabilities map[string]CapabilityAssignment `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`

	// An optional map of interface assignments for the Node Template.
	Interfaces map[string]InterfaceAssignment `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`

	// An optional map of artifact definitions for the Node Template.
	Artifacts map[string]ArtifactDefinition `yaml:"artifacts,omitempty" json:"artifacts,omitempty"`

	// The optional filter definition that TOSCA orchestrators will use to select the correct target node.
	NodeFilter NodeFilter `yaml:"node_filter,omitempty" json:"node_filter,omitempty"`

	// The optional (symbolic) name of another node template to copy from (all keynames and values) and use as a basis for this node template. The source node template provided MUST NOT itself use the copy keyname.
	Copy string `yaml:"copy,omitempty" json:"copy,omitempty"`
}

func NewNodeTemplate() NodeTemplate {
	return NodeTemplate{
		Metadata:     make(map[string]string),
		Properties:   make(map[string]interface{}),
		Attributes:   make(map[string]interface{}),
		Capabilities: make(map[string]CapabilityAssignment),
		Interfaces:   make(map[string]InterfaceAssignment),
		Artifacts:    make(map[string]ArtifactDefinition),
	}
}

type NodeFilter struct {
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	// An optional list of property filters that will be used to select (filter) matching TOSCA entities (e.g., Node Template, Node Type, Capability Types, etc.) based upon their property definitions’ values.
	Properties []PropertyFilterDefinition `yaml:"properties,omitempty" json:"properties,omitempty"`

	// An optional list of capability names or types that will be used to select (filter) matching TOSCA entities based upon their existence.
	Capabilities []CapabilityType `yaml:"capabilities,omitempty" json:"capabilities,omitempty"` // TODO list of capabilityTypes OR capabilityTypeNames

	// Capabilities used as filters often have their own sets of properties which also can be used to construct a filter.

	// An optional list of property filters that will be used to select (filter) matching TOSCA entities (e.g., Node Template, Node Type, Capability Types, etc.) based upon their capabilities’ property definitions’ values.
	CapabilityProperties []PropertyFilterDefinition `yaml:"capability_properties,omitempty" json:"capability_properties,omitempty"` // TODO [4.3.5.7.2] "within a capability name or type name"
}

func (src NodeFilter) Equal(dest NodeFilter) bool {
	for key := range src.Properties {
		if !src.Properties[key].Equal(dest.Properties[key]) {
			return false
		}
	}
	for key := range src.Capabilities {
		if !src.Capabilities[key].Equal(dest.Capabilities[key]) {
			return false
		}
	}
	for key := range src.CapabilityProperties {
		if !src.CapabilityProperties[key].Equal(dest.CapabilityProperties[key]) {
			return false
		}
	}
	return true
}
