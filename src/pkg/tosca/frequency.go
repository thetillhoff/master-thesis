package tosca

import (
	"errors"
)

type Frequency struct {
	DataType

	Value *uint64 `yaml:",inline,omitempty" json:",inline,omitempty"` // Unit is Hz.

	// units: (don't forget toLower)
	// hz
	// khz = 1000 hz
	// mhz = 1000 khz
	// ghz = 1000 mhz
}

func (value Frequency) Equal(arg Frequency) bool {
	return *value.Value == *arg.Value
}
func (value Frequency) GreaterThan(arg Frequency) bool {
	return *value.Value > *arg.Value
}

func ParseFrequency(arg interface{}) (uint64, error) {
	var (
		value uint64
	)
	if typedArg, ok := arg.(uint64); ok {
		// TODO: retrieve and remove of unit, then parse as uint, then use unit for multiplication
		return typedArg, nil
	}
	return value, errors.New("cannot parse to frequency")
}
