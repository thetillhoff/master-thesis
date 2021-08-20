package tosca

type Integer struct {
	DataType
	Value int `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value Integer) Equal(arg Integer) bool {
	return value.Value == arg.Value
}

func (value Integer) GreaterThan(arg Integer) bool {
	return value.Value > arg.Value
}
