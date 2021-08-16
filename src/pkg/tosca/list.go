package tosca

type List struct {
	DataTypeRoot
	Value []Equallable
}

func (value List) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(List); ok {
		if len(value.Value) != len(typedArg.Value) { // unequal length makes them unequal
			return false
		}
		for index := range value.Value {
			if !value.Value[index].Equal(typedArg.Value[index]) {
				return false
			}
		}
		return true
	} // if they are not the same type, they can't be equal ;)
	return false
}
func (value List) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(List); ok {
			for _, element := range typedArg.Value {
				if value.Equal(element) {
					return true
				}
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
