package tosca

// A workflow condition clause definition is used to specify a condition that can be used within a workflow precondition or workflow filter.
//
// Keynames are mutually exclusive, i.e. a filter definition can define only one of the and, or, or not keynames.
type ConditionClauseDefinition struct {

	// grammar
	// and: <list_of_condition_clause_definition>
	//
	// or: <list_of_condition_clause_definition>
	//
	// not: <list_of_condition_clause_definition>
	//
	// direct assertion definition grammar
	// <attribute_name>: <list_of_constraint_clauses>
	//
	// attribute_name: represents the name of an attribute defined on the assertion context entity (node instance, relationship instance, group instance) and from which value will be evaluated against the defined constraint clauses.
	// list_of_constraint_clauses: represents the list of constraint clauses that will be used to validate the attribute assertion.

	// [conditional] An and clause allows to define sub-filter clause definitions that must all be evaluated truly so the and clause is considered as true.
	And []ConditionClauseDefinition `yaml:"and,omitempty" json:"and,omitempty"`

	// [conditional] An or clause allows to define sub-filter clause definitions where one of them must all be evaluated truly so the or clause is considered as true.
	Or []ConditionClauseDefinition `yaml:"or,omitempty" json:"or,omitempty"`

	// [conditional] A not clause allows to define sub-filter clause definitions where one or more of them must be evaluated as false.
	Not []ConditionClauseDefinition `yaml:"not,omitempty" json:"not,omitempty"`
}
