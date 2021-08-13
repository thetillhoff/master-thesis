package tosca

type AbstractType struct {

	// None of these attributes (or better: their contained values) should be inherited when tosca's type derivation is used.

	DerivedFrom string            `yaml:"derived_from,omitempty" json:"derived_from,omitempty"` // An optional parent type name from which this type derives.
	Version     ToscaVersion      `yaml:"version,omitempty" json:"version,omitempty"`           // An optional version for the type definition.
	Metadata    map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`         // Defines a section used to declare additional metadata information.
	Description string            `yaml:"description,omitempty" json:"description,omitempty"`   // An optional description for the type.

}
