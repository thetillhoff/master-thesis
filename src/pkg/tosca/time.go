package tosca

import "time"

type Time struct {
	Value time.Duration

	// units: (don't forget toLower)
	// d
	// h
	// m
	// s
	// ms
	// us
	// ns
}

func ParseTime(input string) (Time, error) {
	var (
		newValue time.Duration
		err      error
	)

	newValue, err = time.ParseDuration(input)
	return Time{Value: newValue}, err

}

func (value Time) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Time) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Time); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Time) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Value > typedArg.Value
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Time); ok {
		if typedUpperBound, ok := upperBound.(Time); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) LengthEquals(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Value == typedArg.Value
	}
	return false
}
func (value Time) MinLength(arg Comparable) bool { // inclusive minimum
	if typedArg, ok := arg.(Time); ok {
		return value.Value >= typedArg.Value
	}
	return false
}
func (value Time) MaxLength(arg Comparable) bool { // inclusive maximum
	if typedArg, ok := arg.(Time); ok {
		return value.Value <= typedArg.Value
	}
	return false
}
