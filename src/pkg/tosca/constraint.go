package tosca

type ConstraintClauseDefinition struct {

	// allowed keynames [4.4.6.1]:
	// - equal
	// - greater_than
	// - greater_or_equal
	// - less_than
	// - less_or_equal
	// - in_range
	// - valid_values
	// - length
	// - min_length
	// - max_length
	// - pattern
	// - schema
	Operator string `yaml:"operator,omitempty" json:"operator,omitempty"`

	// TODO: implement via functions or similar
	// idea: use yaml-schema / json-schema for this. nicky?

}
