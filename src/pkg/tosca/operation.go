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

// Custom unmarshaller, since both single-line and multi-line grammar have to be supported
func (operationAssignment *OperationAssignment) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var (
		implementation ImplementationDefinition
		err            error

		multilineOperationAssignment struct { // Basically the same as OperationAssignment, but without a custom unmarshaller.
			Implementation ImplementationDefinition              `yaml:"implementation,omitempty" json:"implementation,omitempty"`
			Inputs         map[string]ParameterDefinition        `yaml:"inputs,omitempty" json:"inputs,omitempty"`
			Outputs        map[string]ParameterMappingAssignment `yaml:"outputs,omitempty" json:"outputs,omitempty"`
		}
	)

	// Try single-line grammar
	err = unmarshal(&implementation)
	if err == nil {
		operationAssignment.Implementation = implementation
		return nil
	}

	// Try multi-line grammar
	err = unmarshal(&multilineOperationAssignment)
	if err == nil {
		operationAssignment.Implementation = multilineOperationAssignment.Implementation
		operationAssignment.Inputs = multilineOperationAssignment.Inputs
		operationAssignment.Outputs = multilineOperationAssignment.Outputs
		return nil
	}

	return err
}
