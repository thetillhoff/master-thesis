package tosca

type Byte struct {
	DataTypeRoot
	Value byte
}

func (value Byte) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Byte); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Byte) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Byte); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
