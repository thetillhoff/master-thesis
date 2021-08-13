package tosca

// All entries in a map or list for one property or parameter must be of the same type. Similarly, all keys for map entries for one property or parameter must be of the same type as well. A TOSCA schema definition specifies the type (for simple entries) or schema (for complex entries) for keys and entries in TOSCA set types such as the TOSCA list or map.
//
// If the schema definition specifies a map key, the type of the key schema must be derived originally from the string type (which basically ensures that the schema type is a string with additional constraints). As there is little need for complex keys this caters to more straight-forward and clear specifications. If the key schema is not defined it is assumed to be string by default.
//
// Schema definitions appear in data type definitions when derived_from a map or list type or in parameter, property, or attribute definitions of a map or list type.
//
// To prevent a declaration cycle, KeySchema and EntrySchema are of type interface{}
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
	Constraints []map[Operator]interface{} `yaml:"constraints,omitempty" json:"constraints,omitempty"`

	// When the schema itself is of type map, the optional schema definition that is used to specify the type of the keys of that map’s entries (if key_schema is not defined it is assumed to be “string” by default). For other schema types, the key_schema must not be defined.
	KeySchema interface{} `yaml:"key_schema,omitempty" json:"key_schema,omitempty"`

	// When the schema itself is of type map or list, the schema definition is mandatory and is used to specify the type of the entries in that map or list. For other schema types, the entry_schema must not be defined.
	EntrySchema interface{} `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`
}
