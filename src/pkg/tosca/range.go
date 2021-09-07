package tosca

import (
	"errors"
	"reflect"

	"gopkg.in/yaml.v3"
)

// Keyword "UNBOUND" is mapped to nil.
type Range struct { // example: [ 1, 4 ]
	EquallableTypeRoot `yaml:",omitempty" json:",omitempty"`

	LowerBound *Comparable
	UpperBound *Comparable
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
		r.LowerBound = nil
	} else { // lowerBound not UNBOUND
		if err := value.Content[0].Decode(&r.LowerBound); err != nil {
			return err
		}
	}

	if err := value.Content[1].Decode(&test); err == nil { // if upperBound is UNBOUND
		r.UpperBound = nil
	} else { // upperBound not UNBOUND
		if err := value.Content[0].Decode(&r.UpperBound); err != nil {
			return err
		}
	}

	return nil
}

func (value Range) Equal(arg Range) bool {
	return (*value.LowerBound).Equals(*arg.LowerBound) &&
		(*value.UpperBound).Equals(*arg.UpperBound)
}
func (value Range) ContainedIn(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	// func (value Range) InRange(parentRange Range) bool { // "inclusive"
	if value.LowerBound == nil && lowerBound != nil { // parent Range lowerbound bounded, but own is unbounded
		return false
	} else if value.UpperBound == nil && upperBound != nil { // parent Range upperbound bounded, but own is unbounded
		return false
	} else if reflect.TypeOf(value.LowerBound) == reflect.TypeOf(lowerBound) &&
		(*value.LowerBound).GreaterOrEqual(lowerBound) &&
		reflect.TypeOf(value.UpperBound) == reflect.TypeOf(upperBound) &&
		(*value.UpperBound).LessOrEqual(upperBound) {
		return true
	} else {
		return false
	}
}

// TODO func ParseRange(arg interface{}) (int, error) {
// 	var (
// 		value int
// 	)
// 	if typedArg, ok := arg.(int); ok {
// 		return typedArg, nil
// 	}
// 	return value, errors.New("cannot parse to range")
// }
