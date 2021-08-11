package tosca

// A Policy Type defines a type of a policy that affects or governs an application or service’s topology at some stage of its lifecycle, but is not explicitly part of the topology itself (i.e., it does not prevent the application or service from being deployed or run if it did not exist).
type PolicyType struct {
	abstractType

	// grammar
	// <policy_type_name>:
	//   derived_from: <parent_policy_type_name>
	//   version: <version_number>
	//   metadata:
	//     <map of string>
	//   description: <policy_description>
	//   properties:
	//     <property_definitions>
	//   targets: [ <list_of_valid_target_types> ]
	//   triggers:
	//     <trigger_definitions>
	//
	// policy_type_name: represents the mandatory symbolic name of the Policy Type being declared as a string.
	// parent_policy_type_name: represents the name (string) of the Policy Type this Policy Type definition derives from (i.e., its “parent” type).
	// version_number: represents the optional TOSCA version number for the Policy Type.
	// policy_description: represents the optional description string for the corresponding policy_type_name.
	// property_definitions: represents the optional map of property definitions for the Policy Type.
	// list_of_valid_target_types: represents the optional list of TOSCA types (i.e. Group or Node Types) that are valid targets for this Policy Type; if the targets keyname is not defined then there are no restrictions to the targets’ types.
	// trigger_definitions: represents the optional map of trigger definitions for the policy.

	// An optional map of property definitions for the Policy Type.
	Properties map[string]PropertyDefinition `yaml:"properties,omitempty" json:"properties,omitempty"`

	// An optional list of valid Node Types or Group Types the Policy Type can be applied to.
	Targets []string `yaml:"targets,omitempty" json:"targets,omitempty"`

	// An optional map of policy triggers for the Policy Type.
	Triggers map[string]TriggerDefinition `yaml:"triggers,omitempty" json:"triggers,omitempty"`
}

// A policy definition defines a policy that can be associated with a TOSCA topology or top-level entity definition (e.g., group definition, node template, etc.).
type PolicyDefinition struct {

	// grammar
	// <policy_name>:
	//   type: <policy_type_name>
	//   description: <policy_description>
	//   metadata:
	//     <map of string>
	//   properties:
	//     <property_assignments>
	//   targets: [<list_of_policy_targets>]
	//   triggers:
	//     <trigger_definitions>
	//
	// policy_name: represents the mandatory symbolic name of the policy as a string.
	// policy_type_name: represents the name of the policy the definition is based upon.
	// policy_description: contains an optional description of the policy.
	// property_assignments: represents the optional map of property assignments for the policy definition  that provide values for properties defined in its declared Policy Type.
	// list_of_policy_targets: represents the optional list of names of node templates or groups that the policy is to applied to.
	//   if the targets keyname was defined (by specifying a list_of_valid_target_types) in the policy type of this policy then the targets listed here must be compatible (i.e. be of that type or of type that is derived from) with the types (of nodes or groups) in the list_of_valid_target_types.
	// trigger_definitions: represents the optional map of trigger definitions for the policy; these triggers apply in addition to the triggers defined in the policy type.

	// [mandatory] The name of the policy type the policy definition is based upon.
	PolicyType string `yaml:"type,omitempty" json:"type,omitempty"`

	// The optional description for the policy definition.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`

	// An optional map of property value assignments for the policy definition.
	Properties map[string]PropertyAssignment `yaml:"properties,omitempty" json:"properties,omitempty"`

	// An optional list of valid Node Templates or Groups the Policy can be applied to.
	Targets []string `yaml:"targets,omitempty" json:"targets,omitempty"`

	// An optional map of trigger definitions to invoke when the policy is applied by an orchestrator against the associated TOSCA entity. These triggers apply in addition to the triggers defined in the policy type.
	Triggers map[string]TriggerDefinition `yaml:"triggers,omitempty" json:"triggers,omitempty"`
}
