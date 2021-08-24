package tosca

import (
	"errors"
	"strconv"
	"strings"
)

type Version struct {
	// <major_version>.<minor_version>[.<fix_version>[.<qualifier>[-<build_version] ] ]

	DataType

	MajorVersion int    `yaml:"major_version" json:"major_version"` // mandatory
	MinorVersion int    `yaml:"minor_version" json:"minor_version"` // mandatory
	FixVersion   int    `yaml:"fix_version,omitempty" json:"fix_version,omitempty"`
	Qualifier    string `yaml:"qualifier,omitempty" json:"qualifier,omitempty"`
	BuildVersion int    `yaml:"build_version,omitempty" json:"build_version,omitempty"`
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

func (value Version) Equal(arg Version) bool {
	return value.MajorVersion == arg.MajorVersion &&
		value.MinorVersion == arg.MinorVersion &&
		value.FixVersion == arg.FixVersion &&
		value.Qualifier == arg.Qualifier &&
		value.BuildVersion == arg.BuildVersion
}
func (value Version) GreaterThan(arg Version) bool {
	if value.MajorVersion > arg.MajorVersion { // MajorVersion larger
		return true
	} else if value.MinorVersion > arg.MinorVersion { // MinorVersion larger
		return true
	} else if value.FixVersion > arg.FixVersion { // FixVersion larger
		return true
	} else if value.MajorVersion == arg.MajorVersion &&
		value.MinorVersion == arg.MinorVersion &&
		value.FixVersion == arg.FixVersion &&
		value.Qualifier == "" && arg.Qualifier != "" { // Versions that include the optional Qualifier are considered older than those without
		return true
	} else if value.MajorVersion == arg.MajorVersion &&
		value.MinorVersion == arg.MinorVersion &&
		value.FixVersion == arg.FixVersion &&
		value.Qualifier == arg.Qualifier &&
		arg.BuildVersion > value.BuildVersion { // Versions with same major, minor and fix versions and same Qualifier string are compared based on build version
		return true
	}
	return false
}
