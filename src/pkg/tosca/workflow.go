package tosca

// A workflow definition defines an imperative workflow that is associated with a TOSCA topology. A workflow definition can either include the steps that make up the workflow, or it can refer to an artifact that expresses the workflow using an external workflow language.
type WorkflowDefinition struct { // grammar is incomplete in docs [4.7.1.2]

	// grammar
	// <workflow_name>:
	//   description: <workflow_description>
	//   metadata:
	//     <map of string>
	//   inputs:
	//     <parameter_definitions>
	//   preconditions:
	//     - <workflow_precondition_definition>
	//   steps:
	//     <workflow_steps>
	//   implementation:
	//     <operation_implementation_definitions>
	//   outputs:
	//     <attribute_mappings>
	//
	// workflow_name:
	// workflow_description:
	// parameter_definitions:
	// workflow_precondition_definition:
	// workflow_steps:
	// operation_implementation_definition: represents a full inline definition of an implementation artifact
	// attribute_mappings: represents the optional map of of attribute_mappings that consists of named output values returned by operation implementations (i.e. artifacts) and associated mappings that specify the attribute into which this output value must be stored.

	// The optional description for the workflow definition.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`

	// The optional map of input parameter definitions.
	Inputs map[string]ParameterDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`

	// List of preconditions to be validated before the workflow can be processed.
	Preconditions []PreconditionDefinition `yaml:"preconditions,omitempty" json:"preconditions,omitempty"`

	// An optional map of valid imperative workflow step definitions.
	Steps map[string]StepDefinition `yaml:"steps,omitempty" json:"steps,omitempty"`

	// The optional definition of an external workflow definition. This keyname is mutually exclusive with the steps keyname above.
	Implementation ImplementationDefinition `yaml:"implementation,omitempty" json:"implementation,omitempty"`

	// The optional map of attribute mappings that specify workflow  output values and their mappings onto attributes of a node or relationship defined in the topology.
	Outputs map[string]AttributeMapping `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}

// A workflow condition can be used as a filter or precondition to check if a workflow can be processed or not based on the state of the instances of a TOSCA topology deployment. When not met, the workflow will not be triggered.
type PreconditionDefinition struct {

	// grammar
	// - target: <target_name>
	//   target_relationship: <target_requirement_name>
	//   condition:
	// 	   <list_of_condition_clause_definition>
	//
	// target_name: represents the name of a node template or group in the topology.
	// target_requirement_name: represents the name of a requirement of the node template (in case target_name refers to a node template.
	// list_of_condition_clause_definition: represents the list of condition clauses to be evaluated. The value of the resulting condition is evaluated as an AND clause between the different elements.

	// [mandatory] The target of the precondition (this can be a node template name, a group name)
	Target string `yaml:"target" json:"target"`

	// The optional name of a requirement of the target in case the precondition has to be processed on a relationship rather than a node or group. Note that this is applicable only if the target is a node.
	TargetRelationship string `yaml:"target_relationship,omitempty" json:"target_relationship,omitempty"`

	// A list of workflow condition clause definitions. Assertion between elements of the condition are evaluated as an AND condition.
	Condition []ConditionClauseDefinition `yaml:"condition,omitempty" json:"condition,omitempty"`
}

// A workflow step allows to define one or multiple sequenced activities in a workflow and how they are connected to other steps in the workflow. They are the building blocks of a declarative workflow.
type StepDefinition struct {

	// grammar
	// steps:
	// <step_name>
	//   target: <target_name>
	//   target_relationship: <target_requirement_name>
	//   operation_host: <operation_host_name>
	//   filter:
	//     - <list_of_condition_clause_definition>
	//   activities:
	//     - <list_of_activity_definition>
	//   on_success:
	//     - <target_step_name>
	//   on_failure:
	//     - <target_step_name>
	//
	// target_name: represents the name of a node template or group in the topology.
	// target_requirement_name: represents the name of a requirement of the node template (in case target_name refers to a node template.
	// operation_host: the node on which the operation should be executed
	// list_of_condition_clause_definition: represents a list of condition clause definition.
	// list_of_activity_definition: represents a list of activity definition
	// target_step_name: represents the name of another step of the workflow.

	// [mandatory] The target of the step (this can be a node template name, a group name)
	Target string `yaml:"target" json:"target"`

	// The optional name of a requirement of the target in case the step refers to a relationship rather than a node or group. Note that this is applicable only if the target is a node.
	TargetRelationship string `yaml:"target_relationship,omitempty" json:"target_relationship,omitempty"`

	// This element is mandatory only for relationships and groups target.
	//
	// If target is a relationships operation_host is mandatory and valid_values are SOURCE or TARGET – referring to the relationship source or target node.
	//
	// If target is a group operation_host is optional.
	// If not specified the operation will be triggered on every node of the group.
	// If specified the valid_value is a node_type or the name of a node template.
	OperationHost string `yaml:"operation_host,omitempty" json:"operation_host,omitempty"`

	// Filter is a map of attribute name, list of constraint clause that allows to provide a filtering logic.
	Filter []map[Operator]interface{} `yaml:"filter,omitempty" json:"filter,omitempty"`

	// [mandatory] The list of sequential activities to be performed in this step.
	Activities []ActivityDefinition `yaml:"activities" json:"activities"`

	// The optional list of step names to be performed after this one has been completed with success (all activities has been correctly processed).
	OnSuccess []string `yaml:"on_success,omitempty" json:"on_success,omitempty"`

	// The optional list of step names to be called after this one in case one of the step activity failed.
	OnFailure []string `yaml:"on_failure,omitempty" json:"on_failure,omitempty"`
}
