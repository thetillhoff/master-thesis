package tosca

type Integer struct {
	DataTypeRoot
	Value int
}

func (value Integer) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Integer) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Integer); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Integer) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Value > typedArg.Value
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Integer); ok {
		if typedUpperBound, ok := upperBound.(Integer); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}
