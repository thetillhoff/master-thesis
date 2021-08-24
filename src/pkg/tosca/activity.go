package tosca

// Empty struct for embedding
type ActivityDefinition struct {
}

// An activity defines an operation to be performed in a TOSCA workflow step or in an action body of a policy trigger. Activity definitions can be of the following types:
// - Delegate workflow activity definition:
//   Defines the name of the delegate workflow and optional input assignments. This activity requires the target to be provided by the orchestrator (no-op node or relationship).
// - Set state activity definition:
//   Sets the state of a node.
// - Call operation activity definition:
//   Calls an operation defined on a TOSCA interface of a node, relationship or group. The operation name uses the <interface_name>.<operation_name> notation. Optionally, assignments for the operation inputs can also be provided. If provided, they will override for this operation call the operation inputs assignment in the node template.
// - Inline workflow activity definition:
//   Inlines another workflow defined in the topology (allowing reusability). The definition includes the name of a workflow to be inlined and optional workflow input assignments.
type DelegateWorkflowActivityDefinition struct {
	ActivityDefinition

	// short-notation grammar
	// - delegate: <delegate_workflow_name>
	//
	// extended-notation grammar
	// - delegate:
	//     workflow: <delegate_workflow_name>
	//     inputs:
	//       <parameter_assignments>
	//
	// delegate_workflow_name: represents the name of the workflow of the node provided by the TOSCA orchestrator
	// parameter_assignments: represents the optional map of parameter assignments for passing parameters as inputs to this workflow delegation.

	// [mandatory] Defines the name of the delegate workflow and optional input assignments.
	//
	// This activity requires the target to be provided by the orchestrator (no-op node or relationship).
	Delegate string `yaml:"delegate" json:"delegate"`

	// The name of the delegate workflow. Mandatory in the extended notation.
	Workflow string `yaml:"workflow,omitempty" json:"workflow,omitempty"`

	// The optional map of input parameter assignments for the delegate workflow.
	Inputs map[string]ParameterAssignment `yaml:"inputs,omitempty" json:"inputs,omitempty"`
}

// Sets the state of the target node.
type SetStateActivityDefinition struct {
	ActivityDefinition

	// grammar
	// - set_state: <new_node_state>
	//
	// new_node_state: represents the state that will be affected to the node once the activity is performed.

	// Value of the node state.
	SetState string `yaml:"set_state,omitempty" json:"set_state,omitempty"`
}

// This activity is used to call an operation on the target node. Operation input assignments can be optionally provided.
type CallOperationActivityDefinition struct {
	ActivityDefinition

	// short-notation grammar
	// - call_operation: <operation_name>
	//
	// extended-notation grammar
	// - call_operation:
	//   operation: <operation_name>
	//   inputs:
	//     <parameter_assignments>
	//
	// operation_name: represents the name of the operation that will be called during the workflow execution. The notation used is <interface_sub_name>.<operation_sub_name>, where interface_sub_name is the interface name and the operation_sub_name is the name of the operation whitin this interface.
	// parameter_assignments: represents the optional map of parameter assignments for passing parameters as inputs to this workflow delegation.

	// Defines the opration call. The operation name uses the <interface_name>.<operation_name> notation.
	//
	// Optionally, assignments for the operation inputs can also be provided. If provided, they will override for this operation call the operation inputs assignment in the node template.
	CallOperation string `yaml:"call_operation,omitempty" json:"call_operation,omitempty"`

	// The name of the operation to call, using the <interface_name>.<operation_name> notation.
	//
	// Mandatory in the extended notation.
	Operation string `yaml:"operation,omitempty" json:"operation,omitempty"`

	// The optional map of input parameter assignments for the called operation. Any provided input assignments will override the operation input assignment in the target node template for this operation call.
	Inputs map[string]ParameterAssignment `yaml:"inputs,omitempty" json:"inputs,omitempty"`
}

// This activity is used to inline a workflow in the activities sequence. The definition includes the name of the inlined workflow and optional input assignments.
type InlineWorkflowActivityDefinition struct {
	ActivityDefinition

	// short-notation grammar
	// - inline: <inlined_workflow_name>
	//
	// extended-notation grammar
	// - inline:
	//     workflow: <inlined_workflow_name>
	//     inputs:
	//       <parameter_assignments>
	//
	// inlined_workflow_name: represents the name of the workflow to inline.
	// parameter_assignments: represents the optional map of parameter assignments for passing parameters as inputs to this workflow delegation.

	// [mandatory] The definition includes the name of a workflow to be inlined and optional workflow input assignments.
	Inline string `yaml:"inline" json:"inline"`

	// The name of the inlined workflow. Mandatory in the extended notation.
	Workflow string `yaml:"workflow,omitempty" json:"workflow,omitempty"`

	// The optional map of input parameter assignments for the inlined workflow.
	Inputs map[string]ParameterAssignment `yaml:"inputs,omitempty" json:"inputs,omitempty"`
}
