package tosca

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
