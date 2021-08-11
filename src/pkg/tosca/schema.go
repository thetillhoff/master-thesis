package tosca

type SchemaDefinition struct {

	// grammar
	// <schema_definition>:
	//   type: <schema_type>
	//   description: <schema_description>
	//   constraints:
	// 	   - <schema_constraints>
	//   key_schema: <key_schema_definition>
	//   entry_schema: <entry_schema_definition>

	// The mandatory data type for the key or entry.
	// If this schema definition is for a map key, then the referred type must be derived originally from string.
	DataType string `yaml:"type,omitempty" json:"type,omitempty"`

	// The optional description for the schema.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// The optional list of sequenced constraint clauses for the property.
	Constraints []constraintClause `yaml:"constraints,omitempty" json:"constraints,omitempty"`

	// When the schema itself is of type map, the optional schema definition that is used to specify the type of the keys of that map’s entries (if key_schema is not defined it is assumed to be “string” by default). For other schema types, the key_schema must not be defined.
	KeySchema schemaDefinition `yaml:"key_schema,omitempty" json:"key_schema,omitempty"`

	// When the schema itself is of type map or list, the schema definition is mandatory and is used to specify the type of the entries in that map or list. For other schema types, the entry_schema must not be defined.
	EntrySchema schemaDefinition `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`
}
