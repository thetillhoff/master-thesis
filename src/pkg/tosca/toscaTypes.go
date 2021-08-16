package tosca

import (
	"errors"
	"reflect"
	"regexp"
	"time"
)

type Equallable interface {
	Equal(Equallable) bool
	ValidValues([]Equallable) bool
}

type Comparable interface {
	Equallable
	GreaterThan(Comparable) bool
	GreaterOrEqual(Comparable) bool
	LessThan(Comparable) bool
	LessOrEqual(Comparable) bool
	InRange(Comparable, Comparable) bool
}

type Indexable interface {
	Equallable
	LengthEquals(arg Comparable) bool
	MinLength(arg Comparable) bool
	MaxLength(arg Comparable) bool
}

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

type Integer struct {
	DataTypeRoot
	Value int
}

func (value Integer) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Integer) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Integer); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Integer) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Value > typedArg.Value
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Integer); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Integer) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Integer); ok {
		if typedUpperBound, ok := upperBound.(Integer); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}

type Float struct {
	DataTypeRoot
	Value float32
}

func (value Float) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Float) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Float); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Float) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Value > typedArg.Value
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Float); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Float) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Float); ok {
		if typedUpperBound, ok := upperBound.(Float); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}

type Boolean struct {
	DataTypeRoot
	Value bool
}

func (value Boolean) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Boolean); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Boolean) ValidValues(arg []Equallable) bool { // Doesn't _really_ make sense, but needed to satisfy Equallable interface
	for _, element := range arg {
		if typedArg, ok := element.(Boolean); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}

type Byte struct {
	DataTypeRoot
	Value byte
}

func (value Byte) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Byte); ok {
		return value.Value == typedArg.Value
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Byte) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Byte); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}

type Nil struct {
	DataTypeRoot
}

type List struct {
	DataTypeRoot
	Value []Equallable
}

func (value List) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(List); ok {
		if len(value.Value) != len(typedArg.Value) { // unequal length makes them unequal
			return false
		}
		for index := range value.Value {
			if !value.Value[index].Equal(typedArg.Value[index]) {
				return false
			}
		}
		return true
	} // if they are not the same type, they can't be equal ;)
	return false
}
func (value List) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(List); ok {
			for _, element := range typedArg.Value {
				if value.Equal(element) {
					return true
				}
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}

type Map struct {
	DataTypeRoot
	Value map[string]Equallable
}

func (value Map) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Map); ok {
		if len(value.Value) != len(typedArg.Value) {
			return false
		}
		for key := range value.Value {
			if !value.Value[key].Equal(typedArg.Value[key]) {
				return false
			}
		}
		return true
	} // if they are not the same type, they can't be equal ;)
	return false
}
func (value Map) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedElement, ok := element.(Map); ok {
			if value.Equal(typedElement) {
				return true
			}
		} // if they are note the same type, they can't be equal ;)
	}
	return false
}

type Version struct {
	// <major_version>.<minor_version>[.<fix_version>[.<qualifier>[-<build_version] ] ]
	//
	MajorVersion int    `yaml:"major_version" json:"major_version"` // mandatory
	MinorVersion int    `yaml:"minor_version" json:"minor_version"` // mandatory
	FixVersion   int    `yaml:"fix_version,omitempty" json:"fix_version,omitempty"`
	Qualifier    string `yaml:"qualifier,omitempty" json:"qualifier,omitempty"`
	BuildVersion int    `yaml:"build_version,omitempty" json:"build_version,omitempty"`

	// version string // TODO: only allow specific version types (f.e. SEMVER)
}

func (value Version) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Version); ok {
		return value.MajorVersion == typedArg.MajorVersion &&
			value.MinorVersion == typedArg.MinorVersion &&
			value.FixVersion == typedArg.FixVersion &&
			value.Qualifier == typedArg.Qualifier &&
			value.BuildVersion == typedArg.BuildVersion
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Version) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Version); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Version) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Version); ok {
		if value.MajorVersion > typedArg.MajorVersion { // MajorVersion larger
			return true
		} else if value.MinorVersion > typedArg.MinorVersion { // MinorVersion larger
			return true
		} else if value.FixVersion > typedArg.FixVersion { // FixVersion larger
			return true
		} else if value.MajorVersion == typedArg.MajorVersion &&
			value.MinorVersion == typedArg.MinorVersion &&
			value.FixVersion == typedArg.FixVersion &&
			value.Qualifier == "" && typedArg.Qualifier != "" { // Versions that include the optional Qualifier are considered older than those without
			return true
		} else if value.MajorVersion == typedArg.MajorVersion &&
			value.MinorVersion == typedArg.MinorVersion &&
			value.FixVersion == typedArg.FixVersion &&
			value.Qualifier == typedArg.Qualifier &&
			typedArg.BuildVersion > value.BuildVersion { // Versions with same major, minor and fix versions and same Qualifier string are compared based on build version
			return true
		}
		return false
	}
	return false // if they are not the same type, they can't be compared
}
func (value Version) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Version); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Version) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Version); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Version) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Version); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Version) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Version); ok {
		if typedUpperBound, ok := upperBound.(Version); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}

// Keyword "UNBOUND" is mapped to nil.
type Range struct { // example: [ 1, 4 ]
	LowerBound Comparable `yaml:"lower_bound,omitempty" json:"lower_bound,omitempty"`
	UpperBound Comparable `yaml:"upper_bound,omitempty" json:"upper_bound,omitempty"`
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

type Timestamp struct { // example: 2021-08-11T11:09:32
	Value time.Time
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

type Size struct {
	ByteCount uint64

	// units: (don't forget toLower)
	// b
	// kb = 1000 b
	// kib = 1024 b
	// mb = 1000 kb
	// mib = 1024 kib
	// gb = 1000 mb
	// gib = 1024 mib
	// tb = 1000 gb
	// tib = 1024 gib
	// ... <- tb and tib are the last ones documented in tosca spec
	// ~18 eib are the maximum of uint64

}

func (value Size) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.ByteCount == typedArg.ByteCount
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Size) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Size); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Size) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.ByteCount > typedArg.ByteCount
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Size); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Size) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Size); ok {
		if typedUpperBound, ok := upperBound.(Size); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}

type Time struct {
	Time time.Duration

	// units: (don't forget toLower)
	// d
	// h
	// m
	// s
	// ms
	// us
	// ns
}

//TODO Equal should consider constraints etc as well!

func (value Time) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Time == typedArg.Time
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Time) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Time); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Time) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Time > typedArg.Time
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Time); ok {
		if typedUpperBound, ok := upperBound.(Time); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}
func (value Time) LengthEquals(arg Comparable) bool {
	if typedArg, ok := arg.(Time); ok {
		return value.Time == typedArg.Time
	}
	return false
}
func (value Time) MinLength(arg Comparable) bool { // inclusive minimum
	if typedArg, ok := arg.(Time); ok {
		return value.Time >= typedArg.Time
	}
	return false
}
func (value Time) MaxLength(arg Comparable) bool { // inclusive maximum
	if typedArg, ok := arg.(Time); ok {
		return value.Time <= typedArg.Time
	}
	return false
}

type Frequency struct {
	Frequency int64

	// units: (don't forget toLower)
	// hz
	// khz = 1000 hz
	// mhz = 1000 khz
	// ghz = 1000 mhz
}

func (value Frequency) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Frequency == typedArg.Frequency
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Frequency) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Frequency); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Frequency) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Frequency > typedArg.Frequency
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Frequency); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Frequency) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Frequency); ok {
		if typedUpperBound, ok := upperBound.(Frequency); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}

type Bitrate struct {
	Bitrate uint64

	// units: (don't forget toLower)
	// bps
	// kbps = 1000 bps
	// kibps = 1024 bps
	// mbps = 1000 kbps
	// mibps = 1024 kibps
	// gbps = 1000 mbps
	// gibps = 1024 mibps
	// tbps = 1000 gbps
	// tibps = 1024 gibps
	// ... <- tbps and tibps are the last ones documented in tosca spec
	// ~18 eibps are the maximum of uint64

}

func (value Bitrate) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Bitrate); ok {
		return value.Bitrate == typedArg.Bitrate
	}
	return false // if they are not the same type, they can't be equal ;)
}
func (value Bitrate) ValidValues(arg []Equallable) bool {
	for _, element := range arg {
		if typedArg, ok := element.(Bitrate); ok {
			if value.Equal(typedArg) {
				return true
			}
		} // if they are not the same type, they can't be equal ;)
	}
	return false
}
func (value Bitrate) GreaterThan(arg Comparable) bool {
	if typedArg, ok := arg.(Bitrate); ok {
		return value.Bitrate > typedArg.Bitrate
	}
	return false // if they are not the same type, they can't be compared
}
func (value Bitrate) GreaterOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Bitrate); ok {
		return value.Equal(typedArg) || value.GreaterThan(typedArg) // if equal or greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Bitrate) LessThan(arg Comparable) bool {
	if typedArg, ok := arg.(Bitrate); ok {
		return !value.Equal(typedArg) && !value.GreaterThan(typedArg) // if not equal and not greater
	}
	return false // if they are not the same type, they can't be compared
}
func (value Bitrate) LessOrEqual(arg Comparable) bool {
	if typedArg, ok := arg.(Bitrate); ok {
		return value.Equal(typedArg) || value.LessThan(typedArg) // if equal or less
	}
	return false // if they are not the same type, they can't be compared
}
func (value Bitrate) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	if typedLowerBound, ok := lowerBound.(Bitrate); ok {
		if typedUpperBound, ok := upperBound.(Bitrate); ok {
			return value.GreaterOrEqual(typedLowerBound) && value.LessOrEqual(typedUpperBound)
		}
	}
	return false // if they are not the same type, they can't be compared
}

func (dt DataTypeRoot) ValidateConstraints(value interface{}) error {
	var (
		constraint map[Operator]interface{}
		operator   Operator
		arg        interface{}
	)

	for _, constraint = range dt.Constraints {
		if len(constraint) != 1 {
			return errors.New("only one Operator per Constraint allowed")
		}

		// Even though only one entry exists, this for-loop is the easiest way to retrieve the operator and value
		for operator, arg = range constraint {
			switch operator {
			case OperatorEqual:
				// if len(arg) != 1 {
				// 	errors.New("Invalid number of properties/parameters.")
				// }
				if value != arg {
					return errors.New("constraint not fulfilled")
				}
			case OperatorGreaterThan:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.GreaterThan(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorGreatorOrEqual:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.GreaterOrEqual(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorLessThan:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.LessThan(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorLessOrEqual:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.LessOrEqual(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorInRange:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.([]Comparable); ok {
						if reflect.TypeOf(typedArg[0]) == reflect.TypeOf(typedArg[1]) {
							if reflect.TypeOf(typedValue) == reflect.TypeOf(typedArg[0]) {
								if !(typedValue.InRange(typedArg[0], typedArg[1])) {
									return errors.New("constraint not fulfilled")
								}
							} else {
								return errors.New("bounds not of same type as value")
							}
						} else {
							return errors.New("bounds not of same type")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorValidValues:
				if typedValue, ok := value.(Equallable); ok {
					if typedArg, ok := arg.([]Equallable); ok {
						for _, element := range typedArg {
							if reflect.TypeOf(typedValue) == reflect.TypeOf(element) {
								if !(typedValue.ValidValues(typedArg)) {
									return errors.New("constraint not fulfilled")
								}
							} else {
								return errors.New("validvalue not of same type as value")
							}
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorLength:
				if typedValue, ok := value.(Indexable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.LengthEquals(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorMinLength:
				if typedValue, ok := value.(Indexable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.MinLength(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorMaxLength:
				if typedValue, ok := value.(Indexable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.MaxLength(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorPattern:
				if typedValue, ok := value.(String); ok {
					if typedArg, ok := value.(string); ok {
						typedValue.Pattern(typedArg)
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorSchema:
				return errors.New("operator 'schema' is not implemented yet")
			default:
				return errors.New("invalid operator")
			}
		}
	}

	return nil
}
