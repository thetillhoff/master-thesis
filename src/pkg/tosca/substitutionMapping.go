package tosca

// The substitution mapping MUST provide mapping for every property, capability and requirement defined in the specified <node_type>
type SubstitutionMapping struct {

	// grammar
	// node_type: <node_type_name>
	// substitution_filter : <node_filter>
	// properties:
	// 	<property_mappings>
	// capabilities:
	// 	<capability_mappings>
	// requirements:
	// 	<requirement_mappings>
	// attributes:
	// 	<attribute_mappings>
	// interfaces:
	// 	<interface_mappings>
	//
	// node_type_name: represents the mandatory Node Type name that the Service Template’s topology is offering an implementation for.
	// node_filter: represents the optional node filter that reduces the set of abstract node templates for which this topology template is an implementation by only substituting for those node templates whose properties and capabilities satisfy the constraints specified in the node filter.
	// properties: represents the <optional> map of properties mappings.
	// capability_mappings: represents the <optional> map of capability mappings.
	// requirement_mappings: represents the <optional> map of requirement mappings.
	// attributes: represents the <optional> map of attributes mappings.
	// interfaces: represents the <optional> map of interfaces mappings.

	// [mandatory] The mandatory name of the Node Type the Topology Template is providing an implementation for.
	NodeType *string `yaml:"node_type" json:"node_type"`

	// The optional filter that further constrains the abstract node templates for which this topology template can provide an implementation.
	SubstitutionFilter *NodeFilter `yaml:"substitution_filter,omitempty" json:"substitution_filter,omitempty"`

	// The optional map of properties mapping allowing to map properties of the node_type to inputs of the topology template.
	Properties map[string]PropertyMapping `yaml:"properties,omitempty" json:"properties,omitempty"`

	// The optional map of attribute mappings allowing to map outputs from the topology template to attributes of the node_type.
	Attributes map[string]AttributeMapping `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// The optional map of capabilities mapping.
	Capabilities map[string]CapabilityMapping `yaml:"capabilities,omitempty" json:"capabilities,omitempty"`

	// The optional map of requirements mapping.
	Requirements map[string]RequirementMapping `yaml:"requirements,omitempty" json:"requirements,omitempty"`

	// The optional map of interface mapping allows to map an interface and operations of the node type to implementations that could be either workflows or node template interfaces/operations.
	Interfaces map[string]InterfaceMapping `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`
}

// A property mapping allows to map the property of a substituted node type an input of the topology template.
type PropertyMapping struct {

	// single-line grammar
	// <property_name>: [ <input_name> ]

	// multi-line grammar
	// <property_name>:
	//   mapping: [ < input_name > ]

	// An array with 1 string element that references an input of the topology.
	Mapping []string `yaml:"mapping,omitempty" json:"mapping,omitempty"`
}

// Custom unmarshaller, since both single-line and multi-line grammar have to be supported
func (propertyMapping *PropertyMapping) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var (
		mapping []string
		err     error

		multilinePropertyMapping struct { // Basically the same as PropertyMapping, but without a custom unmarshaller.
			Mapping []string `yaml:"mapping,omitempty" json:"mapping,omitempty"`
		}
	)

	// Try single-line grammar
	err = unmarshal(&mapping)
	if err == nil {
		propertyMapping.Mapping = mapping
		return nil
	}

	// Try multi-line grammar
	err = unmarshal(&multilinePropertyMapping)
	if err == nil {
		propertyMapping.Mapping = multilinePropertyMapping.Mapping
		return nil
	}

	return err
}

// An attribute mapping allows to map the attribute of a substituted node type an output of the topology template.
type AttributeMapping struct {

	// grammar
	// <attribute_name>: [ <output_name> ]

	// An array with 1 string element that references an output of the topology.
	Mapping []string `yaml:"mapping,omitempty" json:"mapping,omitempty"`
}

// A capability mapping allows to map the capability of one of the node of the topology template to the capability of the node type the service template offers an implementation for.
type CapabilityMapping struct {

	// single-line grammar
	// <capability_name>: [ <node_template_name>, <node_template_capability_name> ]
	//
	// multi-line grammar
	// <capability_name>:
	//   mapping: [ <node_template_name>, <node_template_capability_name> ]
	//   properties:
	//     <property_name>: <property_value>
	//   attributes:
	//     <attribute_name>: <attribute_value>
	//
	// capability_name: represents the name of the capability as it appears in the Node Type definition for the Node Type (name) that is declared as the value for on the substitution_mappings’ "node_type" key.
	// node_template_name: represents a valid name of a Node Template definition (within the same topology_template declaration as the substitution_mapping is declared).
	// node_template_capability_name: represents a valid name of a capability definition within the <node_template_name> declared in this mapping.
	// property_name: represents the name of a property of the capability.
	// property_value: represents the value to assign to a property of the capability.
	// attribute_name: represents the name a an attribute of the capability.
	// attribute_value: represents the value to assign to an attribute of the capability.

	// [conditional] A list of strings with 2 members, the first one being the name of a node template, the second the name of a capability of the specified node template.
	Mapping []string `yaml:"mapping,omitempty" json:"mapping,omitempty"`

	// [conditional] This field is mutually exclusive with the mapping keyname and allows to provide a capability assignment for the template and specify it’s related properties.
	Properties map[string]interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`

	// [conditional] This field is mutually exclusive with the mapping keyname and allows to provide a capability assignment for the template and specify it’s related attributes.
	Attributes map[string]interface{} `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

// Custom unmarshaller, since both single-line and multi-line grammar have to be supported
func (capabilityMapping *CapabilityMapping) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var (
		mapping []string
		err     error

		multilineCapabilityMapping struct { // Basically the same as CapabilityMapping, but without a custom unmarshaller.
			Mapping    []string               `yaml:"mapping,omitempty" json:"mapping,omitempty"`
			Properties map[string]interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`
			Attributes map[string]interface{} `yaml:"attributes,omitempty" json:"attributes,omitempty"`
		}
	)

	// Try single-line grammar
	err = unmarshal(&mapping)
	if err == nil {
		capabilityMapping.Mapping = mapping
		return nil
	}

	// Try multi-line grammar
	err = unmarshal(&multilineCapabilityMapping)
	if err == nil {
		capabilityMapping.Mapping = multilineCapabilityMapping.Mapping
		capabilityMapping.Properties = multilineCapabilityMapping.Properties
		capabilityMapping.Attributes = multilineCapabilityMapping.Attributes
		return nil
	}

	return err
}

// A requirement mapping allows to map the requirement of one of the node of the topology template to the requirement of the node type the service template offers an implementation for.
type RequirementMapping struct {
	// single-line grammar
	// <requirement_name>: [ <node_template_name>, <node_template_requirement_name> ]
	//
	// multi-line grammar
	// <requirement_name>:
	//   mapping: [ <node_template_name>, <node_template_capability_name> ]
	//   properties:
	//     <property_name>: <property_value>

	// [conditional] A list of strings with 2 elements, the first one being the name of a node template, the second the name of a requirement of the specified node template.
	Mapping []string `yaml:"mapping,omitempty" json:"mapping,omitempty"`

	// [conditional] This field is mutually exclusive with the mapping keyname and allow to provide a requirement for the template and specify it’s related properties.
	Properties []interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`

	// [conditional] This field is mutually exclusive with the mapping keyname and allow to provide a requirement for the template and specify it’s related attributes.
	Attributes []interface{} `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

// Custom unmarshaller, since both single-line and multi-line grammar have to be supported
func (requirementMapping *RequirementMapping) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var (
		mapping []string
		err     error

		multilineRequirementMapping struct { // Basically the same as RequirementMapping, but without a custom unmarshaller.
			Mapping    []string      `yaml:"mapping,omitempty" json:"mapping,omitempty"`
			Properties []interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`
			Attributes []interface{} `yaml:"attributes,omitempty" json:"attributes,omitempty"`
		}
	)

	// Try single-line grammar
	err = unmarshal(&mapping)
	if err == nil {
		requirementMapping.Mapping = mapping
		return nil
	}

	// Try multi-line grammar
	err = unmarshal(&multilineRequirementMapping)
	if err == nil {
		requirementMapping.Mapping = multilineRequirementMapping.Mapping
		requirementMapping.Properties = multilineRequirementMapping.Properties
		requirementMapping.Attributes = multilineRequirementMapping.Attributes
		return nil
	}

	return err
}

// An interface mapping allows to map a workflow of the topology template to an operation of the node type the service template offers an implementation for.
type InterfaceMapping struct {

	// grammar
	// <interface_name>:
	//   <operation_name>: <workflow_name>
	//
	// interface_name: represents the name of the interface as it appears in the Node Type definition for the Node Type (name) that is declared as the value for on the substitution_mappings’ "node_type" key. Or the name of a new management interface to add to the generated type.
	// operation_name: represents the name of the operation as it appears in the interface type definition.
	// workflow_name: represents the name of a workflow of the template to map to the specified operation.

}
