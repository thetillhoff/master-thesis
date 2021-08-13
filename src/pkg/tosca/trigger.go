package tosca

// A trigger definition defines the event, condition and action that is used to “trigger” a policy it is associated with.
type TriggerDefinition struct {

	// short-notation grammar
	// <trigger_name>:
	//   description: <trigger_description>
	//   event: <event_name>
	//   target_filter:
	//     <event_filter_definition>
	//   condition: <list_of_condition_clause_definitions>
	//   action:
	//     - <list_of_activity_definition>
	//
	// extended-notation grammar
	// <trigger_name>:
	//   description: <trigger_description>
	//   event: <event_name>
	//   target_filter:
	//     <event_filter_definition>
	//   condition:
	//     constraint: <list_of_condition_clause_definitions>
	//     period: <scalar-unit.time> # e.g., 60 sec
	//     evaluations: <integer> # e.g., 1
	//     method: <string> # e.g., average
	//   action:
	//     - <list_of_activity_definition>
	//
	// trigger_name: represents the mandatory symbolic name of the trigger as a string.
	// trigger_description: represents the optional description string for the corresponding trigger_name.
	// event_name: represents the mandatory name of an event associated with an interface notification on the identified resource (node).
	// event_filter_definition: represents the optional filter to use to locate the resource (node) or capability attribute to monitor.
	// list_of_condition_clause_definitions: represents one or multiple condition clause definitions containing one or multiple attribute constraints that can be evaluated;
	//   For the condition to be fulfilled all the condition clause definitions must evaluate to true (i.e. a logical and).
	// list_of_activity_definition: represents the list of activities that are performed if the event and the (optional) condition are met. The activity definitions are the same as the ones used in a workflow step. One could regard these activities as an anonymous workflow that is invoked by this trigger and is applied to the target(s) of this trigger’s policy.

	// The optional description string for the trigger.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// [mandatory] The mandatory name of the event that activates the trigger’s action. A deprecated form of this keyname is “event_type”.
	Event string `yaml:"event,omitempty" json:"event,omitempty"`

	// The optional filter used to locate the attribute to monitor for the trigger’s defined condition. This filter helps locate the TOSCA entity (i.e., node or relationship) or further a specific capability of that entity that contains the attribute to monitor.
	TargetFilter EventFilterDefinition `yaml:"target_filter,omitempty" json:"target_filter,omitempty"`

	// The optional condition which contains a list of condition clause definitions containing one or multiple attribute constraints that can be evaluated. For the condition to be fulfilled all the condition clause definitions must evaluate to true (i.e. a logical and). Note: this is optional since sometimes the event occurrence itself is enough to trigger the action.
	Condition []map[Operator]interface{} `yaml:"condition,omitempty" json:"condition,omitempty"`

	// The list of sequential activities to be performed when the event is triggered, and the condition is met (i.e. evaluates to true).
	Action []ActivityDefinition `yaml:"action,omitempty" json:"action,omitempty"`

	// Additional keynames for the extended condition notation

	// The optional condition which contains a condition clause definition specifying one or multiple attribute constraint that can be monitored.  Note: this is optional since sometimes the event occurrence itself is enough to trigger the action.
	Constraint ConditionClauseDefinition `yaml:"constraint,omitempty" json:"constraint,omitempty"`

	// The optional period to use to evaluate for the condition.
	Period ToscaTime `yaml:"period,omitempty" json:"period,omitempty"`

	// The optional number of evaluations that must be performed over the period to assert the condition exists.
	Evaluations int `yaml:"evaluations,omitempty" json:"evaluations,omitempty"`

	// The optional statistical method name to use to perform the evaluation of the condition.
	Method string `yaml:"method,omitempty" json:"method,omitempty"`
}
