package tosca

type NotificationDefinition struct { // examples in docs TBD

	// short notation:
	// <notification_name>: <notification_implementation_definition>

	// extended notation:
	// <notification_name>:
	//   description: <notification_description>
	//   implementation: <notification_implementation_definition>
	//   outputs:
	//     <parameter_definitions>

	// The optional description string for the associated notification.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// The optional definition of the notification implementation.
	Implementation NotificationImplementationDefinition `yaml:"implementation,omitempty" json:"implementation,omitempty"`

	// The optional map of parameter definitions that specify notification output values.
	// Only as part of node and relationship type definitions, the output definitions may include their mappings onto attributes of the node type or relationship type that contains the definition.
	Outputs map[string]ParameterDefinition `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}

type NotificationAssignment struct { // examples in docs TBD

	// short notation:
	// <notification_name>: <notification_implementation_definition>

	// extended notation:
	// <notification_name>:
	//   implementation: <notification_implementation_definition>
	//   outputs:
	//     <parameter_mapping_assignments>

	// The optional definition of the notification implementation. Overrides implementation provided at notification definition.
	Implementation NotificationImplementationDefinition `yaml:"implementation,omitempty" json:"implementation,omitempty"`

	// The optional map of parameter mapping assignments that specify how notification outputs values are mapped onto attributes of the node or relationship type that contains the notification definition.
	Outputs map[string]ParameterMappingAssignment `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}
