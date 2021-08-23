package tosca

import "regexp"

type String struct {
	DataType `yaml:",inline,omitempty" json:",inline,omitempty"`
	Value    string `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func (value String) Equal(arg String) bool {
	return value.Value == arg.Value
}
func (value String) LengthEquals(arg String) bool {
	return len(value.Value) == len(arg.Value)
}
func (value String) MinLength(arg String) bool { // inclusive minimum
	return len(value.Value) >= len(arg.Value)
}
func (value String) MaxLength(arg String) bool { // inclusive maximum
	return len(value.Value) <= len(arg.Value)
}
func (value String) Pattern(arg string) bool { // regex as argument
	matched, err := regexp.MatchString(arg, value.Value)
	if err != nil {
		return false // invalid patterns can't be matched ;)
	}
	return matched
}
