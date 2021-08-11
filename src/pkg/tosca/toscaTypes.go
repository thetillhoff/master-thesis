package tosca

import (
	"time"
)

// other types
// - string
// - int
// - float
// - boolean
// - bytes
// - nil
// - list
// - map

type ToscaVersion struct {
	// <major_version>.<minor_version>[.<fix_version>[.<qualifier>[-<build_version] ] ]
	//
	MajorVersion int    `yaml:"major_version,omitempty" json:"major_version,omitempty"` // mandatory
	MinorVersion int    `yaml:"minor_version,omitempty" json:"minor_version,omitempty"` // mandatory
	FixVersion   int    `yaml:"fix_version,omitempty" json:"fix_version,omitempty"`
	Qualifier    string `yaml:"qualifier,omitempty" json:"qualifier,omitempty"`
	BuildVersion int    `yaml:"build_version,omitempty" json:"build_version,omitempty"`

	// version string // TODO: only allow specific version types (f.e. SEMVER)
}

// Keyword "UNBOUND" is mapped to nil.
type ToscaIntRange struct { // example: [ 1, 4 ]
	LowerBound int `yaml:"lower_bound,omitempty" json:"lower_bound,omitempty"`
	UpperBound int `yaml:"upper_bound,omitempty" json:"upper_bound,omitempty"`
}

type ToscaTimestamp struct { // example: 2021-08-11T11:09:32
	Timestamp time.Time
}

type ToscaSize struct {
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

type ToscaTime struct {
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

type ToscaFrequency struct {
	Frequency int64

	// units: (don't forget toLower)
	// hz
	// khz = 1000 hz
	// mhz = 1000 khz
	// ghz = 1000 mhz
}

type ToscaBitrate struct {
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
