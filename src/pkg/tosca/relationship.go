package tosca

type RelationshipType struct {
	abstractType

	// predefined states (own can be added)
	// - initial (not created, only defined in template)
	// (- established)

	Properties       map[string]PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`                 // An optional map of property definitions for the Relationship Type.
	Attributes       map[string]AttributeDefinition `yaml:"attributes,omitempty" json:"attributes,omitempty"`                 // An optional map of attribute definitions for the Relationship Type.
	Interfaces       map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`                 // An optional map of interface definitions supported by the Relationship Type.
	ValidTargetTypes []string                       `yaml:"valid_target_types,omitempty" json:"valid_target_types,omitempty"` // An optional list of one or more names of Capability Types that are valid targets for this relationship. If undefined, all Capability Types are valid target targets.

}

// The following keywords MAY be used in place of a node or relationship template name:
// SELF: A TOSCA orchestrator will interpret this keyword as the Node or Relationship Template instance that contains the function at the time the function is evaluated.
// SOURCE: A TOSCA orchestrator will interpret this keyword as the Node Template instance that is at the source end of the relationship that contains the referencing function.
// TARGET: A TOSCA orchestrator will interpret this keyword as the Node Template instance that is at the target end of the relationship that contains the referencing function.
//
// TOSCA orchestrators utilize certain reserved keywords in the execution environments that implementation artifacts for Node or Relationship Templates operations are executed in. They are used to provide information to these implementation artifacts such as the results of TOSCA function evaluation or information about the instance model of the TOSCA application.
// The following keywords are reserved environment variable names in any TOSCA supported execution environment:
//
// TARGETS: For an implementation artifact that is executed in the context of a relationship, this keyword, if present, is used to supply a list of Node Template instances in a TOSCA application’s instance model that are currently target of the context relationship. The value of this environment variable will be a comma-separated list of identifiers of the single target node instances (i.e., the tosca_id attribute of the node).
//
// TARGET: For an implementation artifact that is executed in the context of a relationship, this keyword, if present, identifies a Node Template instance in a TOSCA application’s instance model that is a target of the context relationship, and which is being acted upon in the current operation. The value of this environment variable will be the identifier of the single target node instance (i.e., the tosca_id attribute of the node).
//
// SOURCES: For an implementation artifact that is executed in the context of a relationship, this keyword, if present, is used to supply a list of Node Template instances in a TOSCA application’s instance model that are currently source of the context relationship. The value of this environment variable will be a comma-separated list of identifiers of the single source node instances (i.e., the tosca_id attribute of the node).
//
// SOURCE: For an implementation artifact that is executed in the context of a relationship, this keyword, if present, identifies a Node Template instance in a TOSCA application’s instance model that is a source of the context relationship, and which is being acted upon in the current operation. The value of this environment variable will be the identifier of the single source node instance (i.e., the tosca_id attribute of the node).
type RelationshipTemplate struct {
	RelationshipType string                         `yaml:"type,omitempty" json:"type,omitempty"`               // [mandatory] The name of the Relationship Type the Relationship Template is based upon.
	Description      string                         `yaml:"description,omitempty" json:"description,omitempty"` // An optional description for the Relationship Template.
	Metadata         map[string]string              `yaml:"metadata,omitempty" json:"metadata,omitempty"`       // Defines a section used to declare additional metadata information.
	Properties       map[string]PropertyAssignment  `yaml:"properties,omitempty" json:"properties,omitempty"`   // An optional map of property assignments for the Relationship Template.
	Attributes       map[string]AttributeAssignment `yaml:"attributes,omitempty" json:"attributes,omitempty"`   // An optional map of attribute assignments for the Relationship Template.
	Interfaces       map[string]string              `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`   // An optional map of interface assignments for the relationship template.
	Copy             string                         `yaml:"copy,omitempty" json:"copy,omitempty"`               // The optional (symbolic) name of another relationship template to copy from (all keynames and values) and use as a basis for this relationship template. The source relationship template provided MUST NOT itself use the copy keyname.
}
