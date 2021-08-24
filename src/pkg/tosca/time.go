package tosca

import "time"

type Time struct {
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	Value time.Duration `yaml:",inline,omitempty" json:",inline,omitempty"`

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

func (value Time) Equal(arg Time) bool {
	return value.Value == arg.Value
}
func (value Time) GreaterThan(arg Time) bool {
	return value.Value > arg.Value
}
func (value Time) LengthEquals(arg Time) bool {
	return value.Value == arg.Value
}
func (value Time) MinLength(arg Time) bool { // inclusive minimum
	return value.Value >= arg.Value
}
func (value Time) MaxLength(arg Time) bool { // inclusive maximum
	return value.Value <= arg.Value
}
