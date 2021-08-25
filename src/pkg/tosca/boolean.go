package tosca

import (
	"errors"
)

type Boolean struct {
	DataType `yaml:",inline,omitempty" json:",inline,omitempty"`
	Value    bool `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value Boolean) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Boolean); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}

func ParseBoolean(arg interface{}) (bool, error) {
	var (
		value bool
	)
	if typedArg, ok := arg.(bool); ok {
		return typedArg, nil
	}
	return value, errors.New("cannot parse to bool")
}
