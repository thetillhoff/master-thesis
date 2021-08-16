package tosca

type AbstractType struct {

	// None of these attributes (or better: their contained values) should be inherited when tosca's type derivation is used.

	// An optional parent type name from which this type derives.
	DerivedFrom string `yaml:"derived_from,omitempty" json:"derived_from,omitempty"`

	// An optional version for the type definition.
	Version Version `yaml:"version,omitempty" json:"version,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`

	// An optional description for the type.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}
