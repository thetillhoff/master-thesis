package tosca

type Boolean struct {
	DataTypeRoot
	Value bool
}

func (value Boolean) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Boolean); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Boolean) ValidValues(arg []Equallable) bool { // Doesn't _really_ make sense, but needed to satisfy Equallable interface
	for _, element := range arg {
		if typedArg, ok := element.(Boolean); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
