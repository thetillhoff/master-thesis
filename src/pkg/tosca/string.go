package tosca

import "regexp"

type String struct {
	DataTypeRoot
	Value string
}

func (value String) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(String); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value String) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(String); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value String) LengthEquals(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return len(value.Value) == typedArg.Value
	}
	return false
}
func (value String) MinLength(arg Comparable) bool { // inclusive minimum
	if typedArg, ok := arg.(Integer); ok {
		return len(value.Value) >= typedArg.Value
	}
	return false
}
func (value String) MaxLength(arg Comparable) bool { // inclusive maximum
	if typedArg, ok := arg.(Integer); ok {
		return len(value.Value) <= typedArg.Value
	}
	return false
}
func (value String) Pattern(arg string) bool { // regex as argument
	matched, err := regexp.MatchString(arg, value.Value)
	if err != nil {
		return false // invalid patterns can't be matched ;)
	}
	return matched
}
