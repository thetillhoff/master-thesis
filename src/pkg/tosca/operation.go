package tosca

type OperationDefinition struct {

	// The optional description string for the associated operation.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// The optional definition of the operation implementation. May not be used in an interface type definition (i.e. where an operation is initially defined), but only during refinements.
	Implementation ImplementationDefinition `yaml:"implementation,omitempty" json:"implementation,omitempty"`

	// The optional map of parameter definitions for operation input values.
	Inputs map[string]ParameterDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`

	// The optional map of parameter definitions for operation output values.
	//
	// Only as part of node and relationship type definitions, the output definitions may include mappings onto attributes of the node or relationship type that contains the definition.
	Outputs map[string]ParameterDefinition `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}

type OperationAssignment struct {

	// short notation:
	// <operation_name>: <operation_implementation_definition>

	// extended notation:
	// <operation_name>:
	// implementation: <operation_implementation_definition>
	// inputs:
	//   <parameter_value_assignments>
	// outputs:
	//   <parameter_mapping_assignments>

	// The optional definition of the operation implementation. Overrides implementation provided at operation definition.
	Implementation ImplementationDefinition `yaml:"implementation,omitempty" json:"implementation,omitempty"`

	// The optional map of parameter value assignments for assigning values to operation inputs.
	Inputs map[string]ParameterDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`

	// The optional map of parameter mapping assignments that specify how operation outputs are mapped onto attributes of the node or relationship that contains the operation definition.
	Outputs map[string]ParameterMappingAssignment `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}
