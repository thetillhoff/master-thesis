package tosca

import (
	"errors"
)

type Frequency struct {
	DataType

	Value int64 `yaml:",inline,omitempty" json:",inline,omitempty"` // Unit is Hz.

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

func (value Frequency) Equal(arg Frequency) bool {
	return value.Value == arg.Value
}
func (value Frequency) GreaterThan(arg Frequency) bool {
	return value.Value > arg.Value
}
