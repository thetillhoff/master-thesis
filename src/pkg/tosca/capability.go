package tosca

type CapabilityType struct {
	AbstractType `yaml:",inline,omitempty" json:",inline,omitempty"`

	Properties       map[string]PropertyDefinition  // An optional map of property definitions for the Capability Type.
	Attributes       map[string]AttributeDefinition // An optional map of attribute definitions for the Capability Type.
	ValidSourceTypes []string                       // An optional list of one or more valid names of Node Types that are supported as valid sources of any relationship established to the declared Capability Type. If undefined, all Node Types are valid sources.

	// capability types
	// - endpoint
	//   - network type: PRIVATE/PUBLIC
	//
	// PRIVATE:
	// An alias used to reference the first private network within a property or attribute of a Node or Capability which will be assigned to them by the underlying platform at runtime.
	// A private network contains IP addresses and ports typically used to listen for incoming traffic to an application or service from the Intranet and not accessible to the public internet.
	//
	// PUBLIC:
	// An alias used to reference the first public network within a property or attribute of a Node or Capability which will be assigned to them by the underlying platform at runtime.
	// A public network contains IP addresses and ports typically used to listen for incoming traffic to an application or service from the Internet.
}

type CapabilityDefinition struct {

	// single-line grammar assumes capabilityType keyword is used ('<capability_definition_name>: <capability_type>')
	// multi-line grammer requires named parameters
	// <capability_definition_name>:
	//   type: <capability_type>
	//   description: <capability_description>
	//   properties:
	//     <property_refinements>
	//   attributes:
	//     <attribute_refinements>
	//   valid_source_types: [ <node type_names> ]
	//   occurrences: <range_of_occurrences>

	CapabilityType   string                         `yaml:"type,omitempty" json:"type,omitempty"`                             // [mandatory] The mandatory name of the Capability Type this capability definition is based upon. MUST be derived from parent node type definition OR the same.
	Description      string                         `yaml:"description,omitempty" json:"description,omitempty"`               // The optional description of the Capability definition.
	Properties       map[string]PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`                 // An optional map of property refinements for the Capability definition. The referred properties must have been defined in the Capability Type definition referred by the type keyword. New properties may not be added.
	Attributes       map[string]AttributeDefinition `yaml:"attributes,omitempty" json:"attributes,omitempty"`                 // An optional map of attribute refinements for the Capability definition. The referred attributes must have been defined in the Capability Type definition referred by the type keyword. New attributes may not be added.
	ValidSourceTypes []string                       `yaml:"valid_source_types,omitempty" json:"valid_source_types,omitempty"` // An optional list of one or more valid names of Node Types that are supported as valid sources of any relationship established to the declared Capability Type. If undefined, all node types are valid sources. If valid_source_types is defined in the Capability Type, each element in this list must either be in or derived from an element in the list defined in the type.
	Occurences       ToscaIntRange                  `yaml:"occurences,omitempty" json:"occurences,omitempty"`                 // The optional minimum and maximum of occurrences for the capability. The occurrence represents the maximum number of relationships that are allowed by the Capability. If not defined the implied default is [1,UNBOUNDED] (which means that an exported Capability should allow at least one relationship to be formed with it and maximum a UNBOUNDED number of relationships). MUST be within range of parent node type definition.
}

// Custom unmarshaller, since both single-line and multi-line grammar have to be supported
func (capabilityDefinition *CapabilityDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var (
		capabilityType string
		err            error

		multilineCapabilityDefinition struct { // Basically the same as CapabilityDefinition, but without a custom unmarshaller.
			CapabilityType   string                         `yaml:"type,omitempty" json:"type,omitempty"`
			Description      string                         `yaml:"description,omitempty" json:"description,omitempty"`
			Properties       map[string]PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`
			Attributes       map[string]AttributeDefinition `yaml:"attributes,omitempty" json:"attributes,omitempty"`
			ValidSourceTypes []string                       `yaml:"valid_source_types,omitempty" json:"valid_source_types,omitempty"`
			Occurences       ToscaIntRange                  `yaml:"occurences,omitempty" json:"occurences,omitempty"`
		}
	)

	// Try single-line grammar
	err = unmarshal(&capabilityType)
	if err == nil {
		capabilityDefinition.CapabilityType = capabilityType
		return nil
	}

	// Try multi-line grammar
	err = unmarshal(&multilineCapabilityDefinition)
	if err == nil {
		capabilityDefinition.CapabilityType = multilineCapabilityDefinition.CapabilityType
		capabilityDefinition.Description = multilineCapabilityDefinition.Description
		capabilityDefinition.Properties = multilineCapabilityDefinition.Properties
		capabilityDefinition.Attributes = multilineCapabilityDefinition.Attributes
		capabilityDefinition.ValidSourceTypes = multilineCapabilityDefinition.ValidSourceTypes
		capabilityDefinition.Occurences = multilineCapabilityDefinition.Occurences
		return nil
	}

	return err
}

type CapabilityAssignment struct {
	Properties map[string]interface{}         `yaml:"properties,omitempty" json:"properties,omitempty"` // An optional map of property assignments for the Capability definition.
	Attributes map[string]AttributeAssignment `yaml:"attributes,omitempty" json:"attributes,omitempty"` // An optional map of attribute assignments for the Capability definition.
	Occurences int                            `yaml:"occurences,omitempty" json:"occurences,omitempty"` // An optional integer that sets the number of occurrences. It defines the maximum number of allowed relationships to this capability. Must be within the range specified in the corresponding capability definition. If not defined, the orchestrator uses a suitable value from the range defined in the corresponding capability definition (e.g. the maximum in the range).
}
