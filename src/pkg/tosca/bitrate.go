package tosca

import (
	"errors"
)

type Bitrate struct {
	Value uint64 // Unit is bps

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

func ParseBitrate(input string) (Bitrate, error) {
	var (
		newValue uint64
		err      error
	)

	// remove whitespace (between value and unit)
	//input = strings.ReplaceAll(input, " ", "")

	err = errors.New("bitrate cannot be parsed. Not implemented")

	return Bitrate{Value: newValue}, err

}

func (value Bitrate) Equal(arg Equallable) bool {
	if typedArg, ok := arg.(Bitrate); ok {
		return value.Value == typedArg.Value
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
		return value.Value > typedArg.Value
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
