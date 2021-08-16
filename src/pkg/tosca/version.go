package tosca

import (
	"errors"
	"strconv"
	"strings"
)

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

func ParseVersion(input string) (Version, error) {
	var (
		version                Version
		err                    error
		splittedVersion        []string
		splitted               []string
		unverifiedBuildVersion string
	)

	splittedVersion = strings.Split(input, ".")
	if !(len(splittedVersion) >= 2) { // minimal version
		return Version{}, errors.New("invalid version: At least <major>.<minor> required")
	}

	version.MajorVersion, err = strconv.Atoi(splittedVersion[0])
	if err != nil {
		return version, err // major invalid: no int
	}
	version.MinorVersion, err = strconv.Atoi(splittedVersion[1])
	if err != nil {
		return version, err // major invalid: no int
	}

	if len(splittedVersion) >= 3 {
		version.FixVersion, err = strconv.Atoi(splittedVersion[2])
		if err != nil {
			return version, err // fix invalid: no int
		}
	}

	if len(splittedVersion) >= 4 {
		version.Qualifier = splittedVersion[3]

		if strings.Contains(version.Qualifier, "-") { // might contain build version -> remove if exists and parse it
			splitted = strings.Split(version.Qualifier, "-")
			unverifiedBuildVersion, splitted = splitted[len(splitted)-1], splitted[:len(splitted)-1]
			splittedVersion[len(splittedVersion)-1] = strings.Join(splitted, "")
			version.BuildVersion, err = strconv.Atoi(unverifiedBuildVersion)
			if err != nil {
				return version, err // build-version invalid: no int
			}
		}
	}

	if len(splittedVersion) > 4 {
		return version, errors.New("invalid version: too many dots (.)")
	}

	return version, nil
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
