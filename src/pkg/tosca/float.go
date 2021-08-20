package tosca

type Float struct {
	DataType

	Value float32 `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value Float) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Float) GreaterThan(arg Float) bool {
	return value.Value > arg.Value
}
