package tosca

type TopologyTemplate struct {
	// The optional description for the Topology Template.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// An optional map of input parameters (i.e., as parameter definitions) for the Topology Template.
	Inputs map[string]ParameterDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`

	// [mandatory] An mandatory map of node template definitions for the Topology Template.
	NodeTemplates map[string]NodeTemplate `yaml:"node_templates,omitempty" json:"node_templates,omitempty"`

	// An optional map of relationship templates for the Topology Template.
	RelationshipTemplates map[string]RelationshipTemplate `yaml:"relationship_templates,omitempty" json:"relationship_templates,omitempty"`

	// An optional map of Group definitions whose members are node templates defined within this same Topology Template.
	Groups map[string]GroupDefinition `yaml:"groups,omitempty" json:"groups,omitempty"`

	// An optional list of Policy definitions for the Topology Template.
	Policies []PolicyDefinition `yaml:"policies,omitempty" json:"policies,omitempty"`

	// An optional map of output parameters (i.e., as parameter definitions) for the Topology Template.
	Outputs map[string]ParameterDefinition `yaml:"outputs,omitempty" json:"outputs,omitempty"`

	// An optional declaration that exports the topology template as an implementation of a Node type. This also includes the mappings between the external Node Types capabilities and requirements to existing implementations of those capabilities and requirements on Node templates declared within the topology template.
	SubstitutionMappings SubstitutionMapping `yaml:"substitution_mappings,omitempty" json:"substitution_mappings,omitempty"`

	// An optional map of imperative workflow definition for the Topology Template.
	Workflows map[string]WorkflowDefinition `yaml:"workflows,omitempty" json:"workflows,omitempty"`
}
