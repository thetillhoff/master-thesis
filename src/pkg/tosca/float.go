package tosca

import "errors"

type Float struct {
	DataType `yaml:",inline,omitempty" json:",inline,omitempty"`
	Value    *float32 `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value Float) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Float); ok {
		return *value.Value == *typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Float) GreaterThan(arg Float) bool {
	return *value.Value > *arg.Value
}

func ParseFloat(arg interface{}) (float32, error) {
	var (
		value float32
	)
	if typedArg, ok := arg.(float32); ok {
		// TODO: remove 'ps' of unit, then parse as size (bytefmt, see Size)
		return typedArg, nil
	}
	return value, errors.New("cannot parse to float")
}
