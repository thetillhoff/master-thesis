package tosca

// A workflow assertion is used to specify a single condition on a workflow filter definition. The assertion allows to assert the value of an attribute based on TOSCA constraints.
type AssertionDefinition struct {

	// no keynames

	// grammar
	// <attribute_name>: <list_of_constraint_clauses>
	//
	// attribute_name: represents the name of an attribute defined on the assertion context entity (node instance, relationship instance, group instance) and from which value will be evaluated against the defined constraint clauses.
	// list_of_constraint_clauses: represents the list of constraint clauses that will be used to validate the attribute assertion.

}
