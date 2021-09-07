package tosca

// An event filter definition defines criteria for selection of an attribute, for the purpose of monitoring it, within a TOSCA entity, or one its capabilities.
type EventFilterDefinition struct {

	// grammar
	// node: <node_type_name> | <node_template_name>
	// requirement: <requirement_name>
	// capability: <capability_name>
	//
	// node_type_name: represents the mandatory name of the node type that will be used to select (filter) the node that contains the attribute to monitor or contains the requirement that references another node that contains the attribute to monitor.
	// node_template_name: represents the mandatory name of the node template that will be used to select (filter) the node that contains the attribute to monitor or contains the requirement that references another node that contains the attribute to monitor.
	// requirement_name: represents the optional name of the requirement that will be used to select (filter) a referenced node that contains the attribute to monitor.
	// capability_name: represents the optional name of a capability that will be used to select (filter) the attribute to monitor. If a requirement_name is specified, then the capability_name refers to a capability of the node that is targeted by the requirement.

	// [mandatory] The mandatory name of the node type or template that contains either the attribute to be monitored or contains the requirement that references the node that contains the attribute to be monitored.
	NodeType *string `yaml:"node" json:"node"`

	// The optional name of the requirement within the filter’s node that can be used to locate a referenced node that contains an attribute to monitor.
	Requirement *string `yaml:"requirement,omitempty" json:"requirement,omitempty"`

	// The optional name of a capability within the filter’s node or within the node referenced by its requirement that contains the attribute to monitor.
	Capability *string `yaml:"capability,omitempty" json:"capability,omitempty"`
}
