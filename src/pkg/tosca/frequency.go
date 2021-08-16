package tosca

import (
	"errors"
)

type Frequency struct {
	Value int64 // Unit is Hz.

	// units: (don't forget toLower)
	// hz
	// khz = 1000 hz
	// mhz = 1000 khz
	// ghz = 1000 mhz
}

func ParseFrequency(input string) (Frequency, error) {
	var (
		newValue int64
		err      error
	)

	// remove whitespace (between value and unit)
	// input = strings.ReplaceAll(input, " ", "")

	err = errors.New("frequency cannot be parsed. Not implemented")

	return Frequency{Value: newValue}, err

}

func (value Frequency) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Frequency) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Frequency); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Frequency) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Value > typedArg.Value
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Frequency); ok {
		if typedUpperBound, ok := upperBound.(Frequency); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}
