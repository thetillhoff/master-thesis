package tosca

import "code.cloudfoundry.org/bytefmt"

type Size struct {
	EquallableTypeRoot

	Value uint64 `yaml:",inline,omitempty" json:",inline,omitempty"` // Unit is byte

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

func ParseSize(input string) (Size, error) {
	var (
		bytes uint64
		err   error
	)

	bytes, err = bytefmt.ToBytes(input)
	return Size{Value: bytes}, err

}

func (value Size) Equal(arg Size) bool {
	return value.Value == arg.Value
}
func (value Size) GreaterThan(arg Size) bool {
	return value.Value > arg.Value
}
