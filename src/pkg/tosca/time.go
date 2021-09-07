package tosca

import (
	"errors"
	"time"
)

type Time struct {
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	Value *time.Duration `yaml:",inline,omitempty" json:",inline,omitempty"`

	// units: (don't forget toLower)
	// d
	// h
	// m
	// s
	// ms
	// us
	// ns
}

func (value Time) Equal(arg *Time) bool {
	return *value.Value == *arg.Value
}
func (value Time) GreaterThan(arg *Time) bool {
	return *value.Value > *arg.Value
}
func (value Time) LengthEquals(arg *Time) bool {
	return *value.Value == *arg.Value
}
func (value Time) MinLength(arg *Time) bool { // inclusive minimum
	return *value.Value >= *arg.Value
}
func (value Time) MaxLength(arg *Time) bool { // inclusive maximum
	return *value.Value <= *arg.Value
}

func ParseTime(arg *interface{}) (*time.Duration, error) {
	var (
		value time.Duration
		ok    bool
	)

	if value, ok = (*arg).(time.Duration); ok {
		return &value, nil
	}
	return &value, errors.New("cannot parse to time")
}
