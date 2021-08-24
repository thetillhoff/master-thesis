package tosca

import "time"

type Timestamp struct { // example: 2021-08-11T11:09:32
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	Value time.Time `yaml:",inline,omitempty" json:",inline,omitempty"`
}

func ParseTimestamp(input string) (Timestamp, error) {
	var (
		newValue time.Time
		err      error
	)

	newValue, err = time.Parse("2006-01-02T15:04:05Z07:00", input) // according to RFC3339, as defined in tosca spec
	return Timestamp{Value: newValue}, err

}

func (value Timestamp) Equal(arg Timestamp) bool {
	return value.Value == arg.Value
}
func (value Timestamp) GreaterThan(arg Timestamp) bool {
	return value.Value.After(arg.Value)
}
