package tosca

type RequirementDefinition struct {
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	// The optional description of the Requirement definition.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// [mandatory] The mandatory keyname used to provide either the: 'symbolic name of a Capability definition' within a target Node Type that can fulfill the requirement. 'name of a Capability Type' that the TOSCA orchestrator will use to select a type-compatible target node to fulfill the requirement at runtime.
	Capability string `yaml:"capability,omitempty" json:"capability,omitempty"`

	// [conditional] The optional keyname used to provide the name of a valid Node Type that contains the capability definition that can be used to fulfill the requirement.	If a symbolic name of a Capability definition has been used for the capability keyname, then the node keyname is mandatory.
	Node string `yaml:"node,omitempty" json:"node,omitempty"`

	// The optional keyname used to provide the name of a valid Relationship Type to construct a relationship when fulfilling the requirement.
	Relationship string `yaml:"relationship,omitempty" json:"relationship,omitempty"`

	// The optional filter definition that TOSCA orchestrators will use to select a type-compatible target node that can fulfill the associated abstract requirement at runtime.
	NodeFilter NodeFilter `yaml:"node_filter,omitempty" json:"node_filter,omitempty"`

	// The optional minimum and maximum occurrences for the requirement. If this key is not specified, the implied default of [1,1] will be used.
	//
	// Note: the keyword UNBOUNDED is also supported to represent any positive integer.
	Occurences Range `yaml:"occurences,omitempty" json:"occurences,omitempty"`

	// Sometimes additional parameters need to be passed to the relationship (perhaps for configuration). Therefore, interface refinements can be declared (e.g. changing implementation definition or declaring additional parameter definitions to be used as inputs/outputs).

	// The optional keyname used to provide the name of the Relationship Type as part of the relationship keyname definition.
	RelationshipType string `yaml:"type,omitempty" json:"type,omitempty"`

	// The optional keyname used to reference declared interface definitions on the corresponding Relationship Type for refinement.
	Interfaces map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`
}

// func (src RequirementDefinition) Equal(dest RequirementDefinition) bool {
// 	// Assumption: When only Description is different, they can still be equal!
// 	if src.Capability != dest.Capability ||
// 		src.Node != dest.Node ||
// 		src.Relationship != dest.Relationship ||
// 		src.NodeFilter != dest.NodeFilter ||
// 		src.Occurences != dest.Occurences ||
// 		src.RelationshipType != dest.RelationshipType ||
// 		src.Interfaces != dest.Interfaces {
// 		return false
// 	}
// 	return true
// }

type RequirementAssignment struct {

	// short notation:
	// <requirement_name>: <node_template_name>

	// extended notation:
	// <requirement_name>:
	// 	 capability: <capability_symbolic_name> | <capability_type_name>
	// 	 node: <node_template_name> | <node_type_name>
	//   relationship: <relationship_template_name> | <relationship_type_name>
	//   node_filter: <node_filter_definition>
	//   occurrences: <occurrences_value>

	// extended grammar with property assignments and interface assignments for the relationship
	// <requirement_name>:
	// 	 # Other keynames omitted for brevity
	// 	 relationship:
	//   	 type: <relationship_template_name> | <relationship_type_name>
	//   	 properties: <property_assignments>
	//   	 interfaces: <interface_assignments>

	// The optional keyname used to provide either the:
	//
	// - symbolic name of a Capability definition within a target node that can fulfill the requirement.
	//
	// - name of a Capability Type that the TOSCA orchestrator will use to select a type-compatible target node to fulfill the requirement at runtime.
	Capability string `yaml:"capability,omitempty" json:"capability,omitempty"`

	// The optional keyname used to identify the target node of a relationship; specifically, it is used to provide either the:
	//
	// - name of a Node Template that can fulfill the target node requirement.
	//
	// - name of a Node Type that the TOSCA orchestrator will use to select a type-compatible target node to fulfill the requirement at runtime.
	Node string `yaml:"node,omitempty" json:"node,omitempty"`

	// The optional keyname used to provide either the:
	//
	// - name of a Relationship Template to use to relate this node to the target node when fulfilling the requirement.
	//
	// - name of a Relationship Type that the TOSCA orchestrator will use to create a relationship to relate this node to the target node when fulfilling the requirement.
	Relationship string `yaml:"relationship,omitempty" json:"relationship,omitempty"`

	// The optional filter definition that TOSCA orchestrators will use to select a type-compatible target node that can fulfill the requirement at runtime.
	NodeFilter NodeFilter `yaml:"node_filter,omitempty" json:"node_filter,omitempty"`

	// An optional keyname that sets the occurrences for this requirement. The sum of all occurrences’ values for all Requirement assignments with the same symbolic name must be within the range specified in the corresponding Requirement definition. If not defined, the assumed occurrences for an assignment is one (1).
	Occurences int `yaml:"occurences,omitempty" json:"occurences,omitempty"`

	// assignment's relationship keyname which is used when property assignments or interface assignments (e.g. changing the implementation keyname or declaring additional parameter definitions as inputs/outputs) need to be provided:

	// The optional keyname used to provide the name of the Relationship Type for the Requirement assignment’s relationship.
	RelationshipType string `yaml:"type,omitempty" json:"type,omitempty"`

	// An optional keyname providing property assignments for the relationship.
	Properties map[string]interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`

	// The optional keyname providing Interface assignments for the corresponding Interface definitions in the Relationship Type.
	Interfaces map[string]InterfaceAssignment `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`
}
