package tosca

// The actual state of the entity, at any point in its lifecycle once instantiated, is reflected by an attribute.
//
// TOSCA orchestrators automatically create an attribute for every declared property (with the same symbolic name) to allow introspection of both the desired state (property) and actual state (attribute). If an attribute is reflected from a property, its initial value is the value of the reflected property.
type AttributeDefinition struct {

	// grammar
	// attributes:
	//   <attribute_name>:
	// 	   type: <attribute_type>
	// 	   description: <attribute_description>
	// 	   default: <default_value>
	// 	   status: <status_value>
	// 	   constraints:
	// 		   - <attribute_constraints>
	// 	   key_schema: <key_schema_definition>
	// 	   entry_schema: <entry_schema_definition>
	// 	   metadata:
	// 		   <metadata_map>

	// [mandatory] The mandatory data type for the attribute.
	DataType string `yaml:"type" json:"type"`

	// The optional description for the attribute.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// An optional key that may provide a value to be used as a default if not provided by another means.
	//
	// This value SHALL be type compatible with the type declared by the attribute definition’s type keyname.
	DefaultValue interface{} `yaml:"default,omitempty" json:"default,omitempty"`

	// The optional status of the attribute relative to the specification or implementation. See supported status values. Defaults to 'supported'.
	//
	// TODO copy from somewhere else
	Status string `yaml:"status,omitempty" json:"status,omitempty"`

	// The optional list of sequenced constraint clauses for the attribute.
	Constraints []map[Operator]interface{} `yaml:"constraints,omitempty" json:"constraints,omitempty"`

	// [conditional] The schema definition for the keys used to identify entries in attributes of type TOSCA map (or types that derive from map). If not specified, the key_schema defaults to string. For attributes of type other than map, the key_schema is not allowed.
	KeySchema SchemaDefinition `yaml:"key_schema,omitempty" json:"key_schema,omitempty"`

	// [conditional] The schema definition for the entries in attributes of TOSCA collection types such as list, map, or types that derive from list or map) If the attribute type is a collection type, the entry schema is mandatory. For other types, the entry_schema is not allowed.
	EntrySchema SchemaDefinition `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`

	// Defines a section used to declare additional metadata information.
	Metadata map[string]string `yaml:"metadata,omitempty" json:"metadata,omitempty"`
}

type AttributeAssignment struct {
	// no keywords

	// grammar
	// <attribute_name>: <attribute_value> | { <attribute_value_expression> }
	//
	// attribute_name: represents the name of an attribute that will be used to select an attribute definition with the same name within on a TOSCA entity (e.g., Node Template, Relationship Template, etc.) which is declared (or reflected from a Property definition) in its declared type (e.g., a Node Type, Node Template, Capability Type, etc.).
	// attribute_value, attribute_value_expresssion: represent the type-compatible value to assign to the attribute.  Attribute values may be provided as the result from the evaluation of an expression or a function.
}

// The attribute_selection_format is a list of the following format:
// [ <SELF | SOURCE | TARGET >, <optional_capability_name>, <attribute_name>, <nested_attribute_name_or_index_1>, ...,  <nested_attribute_name_or_index_or_key_n> ]
type AttributeSelectionFormat struct {

	// A create operation could include an output value that sets an attribute to an initial value, and the subsequence configure operation could then update that same attribute to a new value.
	//
	// It is also possible that a node template assigns a value to an attribute that has an operation output mapped to it (including a value that is the result of calling an intrinsic function).  Orchestrators could use the assigned value for the attribute as its initial value. After the operation runs that maps an output value onto that attribute, the orchestrator must then use the updated value, and the value specified in the node template will no longer be used.

	// [mandatory] For operation outputs in interfaces on node templates, the only allowed keyname is SELF: output values must always be stored into attributes that belong to the node template that has the interface for which the output values are returned.
	//
	// For operation outputs in interfaces on relationship templates, allowable keynames are SELF, SOURCE, or TARGET.
	Target string `yaml:"target" json:"target"`

	// The optional name of the capability within the specified node template that contains the attribute into which the output value must be stored.
	CapabilityName string `yaml:"capability_name,omitempty" json:"capability_name,omitempty"`

	// The name of the attribute into which the output value must be stored.
	AttributeName string `yaml:"attribute_name,omitempty" json:"attribute_name,omitempty"`

	// Some TOSCA attributes are complex (i.e., composed as nested structures).  These parameters are used to dereference into the names of these nested structures when needed.
	//
	// Some attributes represent list or map types. In these cases, an index or key may be provided to reference a specific entry in the list or map (identified by the previous parameter).
	NestedAttribute interface{} `yaml:"nested_attribute,omitempty" json:"nested_attribute,omitempty"` // TODO: only string and int are allowed
}
