package tosca

type Equallable interface {
	Equals(Equallable) bool
	ContainedIn([]Equallable) bool
}

// [3.6.3.1.1, 3.3.6.3] includes integer, float, timestamp, string, version, size, time, frequency, and bitrate
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

func listContainsString(l []string, e string) bool {
	for _, listEntry := range l {
		if listEntry == e {
			return true
		}
	}
	return false
}
