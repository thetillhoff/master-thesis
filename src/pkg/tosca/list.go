package tosca

import "errors"

type List struct {
	DataType
	Value []Equallable `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value List) Equal(arg List) bool {
	if len(value.Value) != len(arg.Value) { // unequal length makes them unequal
		return false
	}
	for index := range value.Value {
		if !value.Value[index].Equals(arg.Value[index]) {
			return false
		}
	}
	return false
}

func (value List) Contains(arg Equallable) bool {
	for _, element := range value.Value {
		if element.Equals(arg) {
			return true
		}
	}
	return false
}

func ParseList(arg interface{}) ([]interface{}, error) {
	var (
		value []interface{}
	)
	if typedArg, ok := arg.([]interface{}); ok {
		// TODO: remove 'ps' of unit, then parse as size (bytefmt, see Size)
		return typedArg, nil
	}
	return value, errors.New("cannot parse to list")
}
