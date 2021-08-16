package tosca

type Float struct {
	DataTypeRoot
	Value float32
}

func (value Float) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Float) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Float); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Float) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Value > typedArg.Value
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Float); ok {
		if typedUpperBound, ok := upperBound.(Float); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}
