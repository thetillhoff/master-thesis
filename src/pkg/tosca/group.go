package tosca

type GroupType struct {
	AbstractType `yaml:",inline,omitempty" json:",inline,omitempty"`

	// - template ; mutually exclusive with type
	// - type ; mutually exclusive with template
	// - members
	// - properties
	// - attributes
	// - capabilities
	// - requirements

	// grammar
	// <group_type_name>:
	//   derived_from: <parent_group_type_name>
	//   version: <version_number>
	//   metadata:
	//     <map of string>
	//   description: <group_description>
	//   properties:
	//     <property_definitions>
	//   attributes:
	//     <attribute_definitions>
	//   members: [ <list_of_valid_member_types> ]
	//
	// group_type_name: represents the mandatory symbolic name of the Group Type being declared as a string.
	// parent_group_type_name: represents the name (string) of the Group Type this Group Type definition derives from (i.e. its “parent” type).
	// version_number: represents the optional TOSCA version number for the Group Type.
	// group_description: represents the optional description string for the corresponding group_type_name.
	// attribute_definitions: represents the optional map of attribute definitions for the Group Type.
	// property_definitions: represents the optional map of property definitions for the Group Type.
	// list_of_valid_member_types: represents the optional list of TOSCA Node Types that are valid member types for being added to (i.e. members of) the Group Type; if the members keyname is not defined then there are no restrictions to the member types;

	// An optional map of property definitions for the Group Type.
	Properties map[string]PropertyDefinition `yaml:"properties,omitempty" json:"properties,omitempty"`

	// An optional map of attribute definitions for the Group Type.
	Attributes map[string]AttributeDefinition `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// An optional list of one or more names of Node Types that are valid (allowed) as members of the Group Type.
	Members []string `yaml:"members,omitempty" json:"members,omitempty"`
}

func NewGroupType() GroupType {
	return GroupType{
		Properties: make(map[string]PropertyDefinition),
		Attributes: make(map[string]AttributeDefinition),
	}
}

type GroupDefinition struct {

	// grammar
	// <group_name>:
	//   type: <group_type_name>
	//   description: <group_description>
	//   metadata:
	//     <map of string>
	//   properties:
	//     <property_assignments>
	//   attributes:
	//     <attribute_assignments>
	//   members: [ <list_of_node_templates> ]
	//
	// group_name: represents the mandatory symbolic name of the group as a string.
	// group_type_name: represents the name of the Group Type the definition is based upon.
	// group_description: contains an optional description of the group.
	// property_assignments: represents the optional map of property assignments for the group definition that provide values for properties defined in its declared Group Type.
	// attribute_assigments: represents the optional map of attribute assignments for the group definition that provide values for attributes defined in its declared Group Type.
	// list_of_node_templates: contains the mandatory list of one or more node template names or group symbolic names (within the same topology template) that are members of this logical group
	//   If the members keyname was defined (by specifying a list_of_valid_member_types) in the group type of this group then the nodes listed here must be compatible (i.e. be of that type or of type that is derived from) with the node types in the list_of_valid_member_types.

	// The mandatory name of the group type the group definition is based upon.
	GroupType string `yaml:"type,omitempty" json:"type,omitempty"`

	// The optional description for the group definition.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`

	// An optional map of property value assignments for the group definition.
	Properties map[string]interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`

	// An optional map of attribute value assignments for the group definition.
	Attributes map[string]AttributeAssignment `yaml:"attributes,omitempty" json:"attributes,omitempty"`

	// The optional list of one or more node template names that are members of this group definition.
	Members []string `yaml:"members,omitempty" json:"members,omitempty"`
}
