package tosca

type Byte struct {
	DataType `yaml:",inline,omitempty" json:",inline,omitempty"`
	Value    *byte `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value Byte) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Byte); ok {
		return *value.Value == *typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
