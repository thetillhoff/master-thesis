package tosca

import "code.cloudfoundry.org/bytefmt"

type Size struct {
	Value uint64 // Unit is byte

	// units: (don't forget toLower)
	// b
	// kb = 1000 b
	// kib = 1024 b
	// mb = 1000 kb
	// mib = 1024 kib
	// gb = 1000 mb
	// gib = 1024 mib
	// tb = 1000 gb
	// tib = 1024 gib
	// ... <- tb and tib are the last ones documented in tosca spec
	// ~18 eib are the maximum of uint64

}

func ParseSize(input string) (Size, error) {
	var (
		bytes uint64
		err   error
	)

	bytes, err = bytefmt.ToBytes(input)
	return Size{Value: bytes}, err

}

func (value Size) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Size) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Size); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Size) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.Value > typedArg.Value
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Size); ok {
		if typedUpperBound, ok := upperBound.(Size); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}
