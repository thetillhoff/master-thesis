package tosca

type DataType struct {
	abstractType

	// grammar
	// <data_type_name>:
	//   derived_from: <existing_type_name>
	//   version: <version_number>
	//   metadata:
	//     <map of string>
	//   description: <datatype_description>
	//   constraints:
	// 	   - <type_constraints>
	//   properties:
	// 	   <property_definitions>
	//   key_schema: <key_schema_definition>
	//   entry_schema: <entry_schema_definition>

	// The optional list of sequenced constraint clauses for the Data Type.
	Constraints []ConstraintClauses `yaml:"constraints,omitempty" json:"constraints,omitempty"`

	// The optional map property definitions that comprise the schema for a complex Data Type in TOSCA.
	// TODO [4.4.4.4] "A valid datatype definition MUST have either a valid derived_from declaration or at least on valid property definition."
	Properties map[string]PropertyDefinition `yaml:"properties,omitempty" json:"properties,omitempty"`

	// [conditional] For data types that derive from the TOSCA map data type, the optional schema definition for the keys used to identify entries in properties of this data type. If not specified, the key_schema defaults to string.
	// For data types that do not derive from the TOSCA map data type, the key_schema is not allowed.
	KeySchema SchemaDefinition `yaml:"key_schema,omitempty" json:"key_schema,omitempty"`

	// [conditional] For data types that derive from the TOSCA map or list data types, the mandatory schema definition for the entries in properties of this data type.
	// For data types that do not derive from the TOSCA list or map data type, the entry_schema is not allowed.
	Entryschema SchemaDefinition `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`
}
