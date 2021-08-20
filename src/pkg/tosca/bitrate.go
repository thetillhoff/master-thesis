package tosca

import (
	"errors"
)

type Bitrate struct {
	DataType
	Value uint64 `yaml:",inline,omitempty" json:",inline,omitempty"` // Unit is bps

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

func (value Bitrate) Equal(arg Bitrate) bool {
	return value.Value == arg.Value
}
func (value Bitrate) GreaterThan(arg Bitrate) bool {
	return value.Value > arg.Value
}
