package tosca

import "errors"

type Integer struct {
	DataType
	Value *int `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value Integer) Equal(arg Integer) bool {
	return *value.Value == *arg.Value
}

func (value Integer) GreaterThan(arg Integer) bool {
	return *value.Value > *arg.Value
}

func ParseInteger(arg interface{}) (int, error) {
	var (
		value int
	)
	if typedArg, ok := arg.(int); ok {
		return typedArg, nil
	}
	return value, errors.New("cannot parse to integer")
}
