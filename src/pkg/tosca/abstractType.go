package tosca

type AbstractType struct {

	// None of these attributes (or better: their contained values) should be inherited when tosca's type derivation is used.

	// An optional parent type name from which this type derives.
	DerivedFrom string `yaml:"derived_from,omitempty" json:"derived_from,omitempty"`

	// A list which contains the list of all ancestors. Ancestors are retrieved by following the chain of "DerivedFrom" parents.
	//
	// Own addition
	derivedFromAncestors []string

	// An optional version for the type definition.
	//
	// Has to be a string sadly, since go doesn't support chaining types (Version inherits from AbstractType indirectly, so it can't be used here).
	Version string `yaml:"version,omitempty" json:"version,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`

	// An optional description for the type.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

func IsDerivedFrom(thisType AbstractType, parent string) bool {
	return listContainsString(thisType.derivedFromAncestors, parent)
}
