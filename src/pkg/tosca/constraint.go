package tosca

// A constraint clause defines an operation along with one or more compatible values that can be used to define a constraint on a property or parameter’s allowed values when it is defined in a TOSCA Service Template or one of its entities.
//
// The recognized operators (keynames) when defining constraint clauses.
//
// Defaults to "equal"
//
// <list_of_constraint_clause_definitions> == []map[Operator]interface{}
type Operator string

const (
	OperatorEqual          Operator = "equal"
	OperatorGreaterThan    Operator = "greater_than"
	OperatorGreatorOrEqual Operator = "greater_or_equal"
	OperatorLessThan       Operator = "less_than"
	OperatorLessOrEqual    Operator = "less_or_equal"
	OperatorInRange        Operator = "in_range"
	OperatorValidValues    Operator = "valid_values"
	OperatorLength         Operator = "length"
	OperatorMinLength      Operator = "min_length"
	OperatorMaxLength      Operator = "max_length"
	OperatorPattern        Operator = "pattern"
	OperatorSchema         Operator = "schema"
)
