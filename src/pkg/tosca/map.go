package tosca

type Map struct {
	DataTypeRoot
	Value map[string]Equallable
}

func (value Map) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Map); ok {
		if len(value.Value) != len(typedArg.Value) {
			return false
		}
		for key := range value.Value {
			if !value.Value[key].Equal(typedArg.Value[key]) {
				return false
			}
		}
		return true
	} // if they are not the same type, they can't be equal ;)
	return false
}
func (value Map) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedElement, ok := element.(Map); ok {
			if value.Equal(typedElement) {
				return true
			}
		} // if they are note the same type, they can't be equal ;)
	}
	return false
}
