package tosca

import "time"

type Timestamp struct { // example: 2021-08-11T11:09:32
	Value time.Time
}

func ParseTimestamp(input string) (Timestamp, error) {
	var (
		newValue time.Time
		err      error
	)

	newValue, err = time.Parse("2006-01-02T15:04:05Z07:00", input) // according to RFC3339, as defined in tosca spec
	return Timestamp{Value: newValue}, err

}

func (value Timestamp) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Timestamp); ok {
		return value.Value.Equal(typedArg.Value)
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Timestamp) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Timestamp); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Timestamp) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Timestamp); ok {
		return value.Value.After(typedArg.Value)
	}
	return false // if they are not the same type, they can't be compared
}
func (value Timestamp) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Timestamp); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Timestamp) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Timestamp); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Timestamp) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Timestamp); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Timestamp) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Timestamp); ok {
		if typedUpperBound, ok := upperBound.(Timestamp); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}
