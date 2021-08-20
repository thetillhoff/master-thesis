package tosca

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
