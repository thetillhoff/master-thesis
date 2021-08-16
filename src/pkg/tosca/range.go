package tosca

import (
	"errors"
	"reflect"

	"gopkg.in/yaml.v3"
)

// Keyword "UNBOUND" is mapped to nil.
type Range struct { // example: [ 1, 4 ]
	LowerBound   Comparable
	NoLowerBound bool
	UpperBound   Comparable
	NoUpperBound bool
}

func (r *Range) UnmarshalYAML(value *yaml.Node) error {
	var (
		test *string
	)

	if value.Kind != yaml.SequenceNode {
		return errors.New("tosca.Range expects a sequence")
	}

	if len(value.Content) != 2 { // both values are mandatory
		return errors.New("tosca.Range expects two bounds")
	}

	if err := value.Content[0].Decode(&test); err == nil { // if lowerBound is UNBOUND
		r.NoLowerBound = true
	} else { // lowerBound not UNBOUND
		if err := value.Content[0].Decode(&r.LowerBound); err != nil {
			return err
		}
	}

	if err := value.Content[1].Decode(&test); err == nil { // if upperBound is UNBOUND
		r.NoUpperBound = true
	} else { // upperBound not UNBOUND
		if err := value.Content[0].Decode(&r.UpperBound); err != nil {
			return err
		}
	}

	return nil
}

func (value Range) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Range); ok {
		return value.LowerBound.Equal(typedArg.LowerBound) &&
			value.UpperBound.Equal(typedArg.UpperBound)
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Range) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Range); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Range) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	// func (value Range) InRange(parentRange Range) bool { // "inclusive"
	if value.LowerBound == nil && lowerBound != nil { // parent Range lowerbound bounded, but own is unbounded
		return false
	} else if value.UpperBound == nil && upperBound != nil { // parent Range upperbound bounded, but own is unbounded
		return false
	} else if reflect.TypeOf(value.LowerBound) == reflect.TypeOf(lowerBound) &&
		value.LowerBound.GreaterOrEqual(lowerBound) &&
		reflect.TypeOf(value.UpperBound) == reflect.TypeOf(upperBound) &&
		value.UpperBound.LessOrEqual(upperBound) {
		return true
	} else {
		return false
	}
}
