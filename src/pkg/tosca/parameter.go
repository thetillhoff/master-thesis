package tosca

// A parameter definition defines a named, typed value and related data and may be used to exchange values between the TOSCA orchestrator and the external world. Such values may be
// - inputs and outputs of interface operations and notifications
// - inputs and outputs of workflows
// - inputs and outputs of service templates
//
// From the perspective of the TOSCA orchestrator such parameters are either "incoming" (i.e. transferring a value from the external world to the orchestrator) or "outgoing" (transferring a value from the orchestrator to the external world). Thus:
// outgoing parameters are:
// – template outputs
// – internal workflow outputs
// – external workflow inputs
// – operation inputs
// incoming parameters are:
// – template inputs
// –      internal workflow inputs
// –      external workflow outputs
// –      operation outputs
// –      notification outputs
//
// An "outgoing" parameter definition is essentially the same as a TOSCA property definition, however it may optionally inherit the data type of the value assigned to it rather than have an explicit data type defined.
//
// An "incoming" parameter definition may define an attribute mapping of the parameter value to an attribute of a node. Optionally, it may inherit the data type of the attribute it is mapped to, rather than have an explicit data type defined for it.
type ParameterDefinition struct {

	// grammar
	// <parameter_name>:
	//   type: <parameter_type>
	//   description: <parameter_description>
	//   value: <parameter_value> | { <parameter_value_expression> }
	//   required: <parameter_required>
	//   default: <parameter_default_value>
	//   status: <status_value>
	//   constraints:
	//     - <parameter_constraints>
	//   key_schema: <key_schema_definition>
	//   entry_schema: <entry_schema_definition>
	//   mapping: <attribute_selection_form>

	// single-line grammar is supported when only a fixed value needs to be provided provided to an outgoing parameter:
	// <parameter_name>: <parameter_value> | { <parameter_value_expression> }
	// OR
	// <parameter_name>:
	//   value: <parameter_value> | { <parameter_value_expression> }

	// single-line grammar is supported when only a parameter to attribute mapping needs to be provided to an incoming parameter:
	// <parameter_name>: <attribute_selection_form>
	// OR
	// <parameter_name>:
	//   mapping: <attribute_selection_form>

	// The data type of the parameter.
	//
	// Note: This keyname is mandatory for a TOSCA Property definition but is NOT mandatory for a TOSCA Parameter definition.
	DataType string `yaml:"type,omitempty" json:"type,omitempty"`

	// The type-compatible value to assign to the parameter. Parameter values may be provided as the result from the evaluation of an expression or a function. May only be defined for outgoing parameters. Mutually exclusive with the "mapping" keyname.
	Value interface{} `yaml:"value,omitempty" json:"value,omitempty"`

	// A mapping that specifies the node or relationship attribute into which the returned output value must be stored. May only be defined for incoming parameters. Mutually exclusive with the "value" keyname.
	Mapping AttributeSelectionFormat `yaml:"mapping,omitempty" json:"mapping,omitempty"`
}

// Parameters that have a (fixed) value defined during their definition or during a subsequent refinement may not be assigned (as their value is already set).
//
// If a required parameter has no value defined or assigned, its default value is assigned.
//
// A non-required parameter that has no value assigned it stays undefined, thus the default keyname is irrelevant for a non-required parameter.
type ParameterAssignment struct {
	// no keynames

	// grammar
	// <parameter_name>: <parameter_value> | { <parameter_value_expression> }
	//
	// parameter_name: represents the symbolic name of the parameter to assign; note that in some cases, even parameters that do not have a corresponding definition in the entity type of the entity containing them may be assigned (see e.g. inputs and outputs in interfaces).
	// parameter_value, parameter_value_expression: represent the type-compatible value to assign to the parameter.  Parameter values may be provided as the result from the evaluation of an expression or a function.
}

type ParameterMappingAssignment struct {
	// no keynames

	// grammar
	// <parameter_name>: <attribute_selection_format>
	// parameter_name: represents the symbolic name of the parameter to assign; note that in some cases, even parameters that do not have a corresponding definition in the entity type of the entity containing them may be assigned (see e.g. inputs and outputs in interfaces).
	// attribute_selection_format: represents a format that is used to select an attribute or a nested attribute on which to map the parameter value of the incoming parameter referred by parameter_name.
}
