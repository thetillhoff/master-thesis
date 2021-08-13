package tosca

type PropertyDefinition struct { // indicate desired state
	// allowed types are

	// - string (default)
	// - integer
	// - float
	// - boolean
	// - timestamp
	// - null

	// - range (upper,lower,UNBOUND)
	// - list
	// - map
	// - scalar-units
	//   - .size == 1b, 2TiB
	//   - .time
	//   - .frequency
	//   - .bitrate

	// [mandatory] The data type for the property.
	DataType string `yaml:"type,omitempty" json:"type,omitempty"`

	// The optional description for the property.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// An optional key that declares a property as required (true) or not (false). Defaults to true.
	Required bool `yaml:"required" json:"required"`

	// An optional key that may provide a value to be used as a default if not provided by another means.
	//
	// The default keyname SHALL NOT be defined when property is not required (i.e. the value of the required keyname is false).
	DefaultValue interface{} `yaml:"default,omitempty" json:"default,omitempty"`

	// An optional key that may provide a fixed value to be used. A property that has a fixed value provided (as part of a definition or refinement) cannot be subject to a further refinement or assignment. That is, a fixed value cannot be changed.
	FixedValue interface{} `yaml:"value,omitempty" json:"value,omitempty"`

	// The optional status of the property relative to the specification or implementation. See table below for valid values. Defaults to supported.
	//
	// "supported"    : Indicates the property is supported.  This is the default value for all property definitions.
	// "unsupported"  : Indicates the property is not supported.
	// "experimental" : Indicates the property is experimental and has no official standing.
	// "deprecated"   : Indicates the property has been deprecated by a new specification version.
	Status string `yaml:"status,omitempty" json:"status,omitempty"`

	// The optional list of sequenced constraint clauses for the property.
	Constraints []ConstraintClauseDefinition `yaml:"constraints,omitempty" json:"constraints,omitempty"`

	// The schema definition for the keys used to identify entries in properties of type TOSCA map (or types that derive from map). If not specified, the key_schema defaults to string. For properties of type other than map, the key_schema is not allowed.
	KeySchema SchemaDefinition `yaml:"key_schema,omitempty" json:"key_schema,omitempty"`

	// The schema definition for the entries in properties of TOSCA collection types such as list, map, or types that derive from list or map) If the property type is a collection type, the entry schema is mandatory. For other types, the entry_schema is not allowed.
	EntrySchema SchemaDefinition `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`

	// The optional key that contains a schema definition that TOSCA Orchestrators MAY use for validation when the “type” key’s value indicates an External schema (e.g., “json”)// See section “External schema” below for further explanation and usage.
	ExternalSchema string `yaml:"external-schema,omitempty" json:"external-schema,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`
}

type PropertyFilterDefinition struct { // docs are incomplete [4.3.5.8]

	// short notation:
	// <property_name>: <property_constraint_clause>

	// extended notation:
	// <property_name>:
	//   - <property_constraint_clause_1>
	//   - ...
	//   - <property_constraint_clause_n>

	// peroperty_constraint_clause: represents constraint clause(s) that will be used to filter entities based upon the named property’s value(s).

}
