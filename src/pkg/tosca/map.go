package tosca

import "errors"

type Map struct {
	DataType

	Value map[string]Equallable `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value Map) Equal(arg Map) bool {
	if len(value.Value) != len(arg.Value) { // unequal length makes them unequal
		return false
	}
	for key := range value.Value {
		if !value.Value[key].Equals(arg.Value[key]) {
			return false
		}
	}
	return true
}

func ParseMap(arg interface{}) (map[string]interface{}, error) {
	var (
		value map[string]interface{}
	)
	if typedArg, ok := arg.(map[string]interface{}); ok {
		return typedArg, nil
	}
	return value, errors.New("cannot parse to map")
}
