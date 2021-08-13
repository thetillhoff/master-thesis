package interfaces

import "github.com/thetillhoff/eat/pkg/tosca"

// This is the default (root) TOSCA Interface Type definition that all other TOSCA Interface Types derive from.
type InterfacesRoot struct {
}

// This lifecycle interface defines the essential, normative operations that TOSCA nodes may support.
type InterfacesNodeLifecycleStandard struct {
	InterfacesRoot `yaml:",inline,omitempty" json:",inline,omitempty"`

	Operations map[InterfacesNodeLifecycleStandardOperation]tosca.OperationDefinition `yaml:",inline,omitempty" json:",inline,omitempty"`
}

type InterfacesNodeLifecycleStandardOperation string

const (
	OperationCreate    = "create"
	OperationConfigure = "configure"
	OperationStart     = "start"
	OperationStop      = "stop"
	OperationDelete    = "delete"
)

// The lifecycle interfaces define the essential, normative operations that each TOSCA Relationship Types may support.
//
// Note: When the TOSCA Orchestrator connects a source and target node together using a relationship that supports the Configure interface it will “interleave” the operations invocations of the Configure interface with those of the node’s own Standard lifecycle interface.
type InterfacesRelationshipConfigure struct {
	InterfacesRoot `yaml:",inline,omitempty" json:",inline,omitempty"`

	// Operation to pre-configure the source endpoint.
	PreConfigureSource tosca.OperationDefinition `yaml:"pre_configure_source,omitempty" json:"pre_configure_source,omitempty"`

	// Operation to pre-configure the target endpoint.
	PreConfigureTarget tosca.OperationDefinition `yaml:"pre_configure_target,omitempty" json:"pre_configure_target,omitempty"`

	// Operation to post-configure the source endpoint.
	PostConfigureSource tosca.OperationDefinition `yaml:"post_configure_source,omitempty" json:"post_configure_source,omitempty"`

	// Operation to post-configure the target endpoint.
	PostConfigureTarget tosca.OperationDefinition `yaml:"post_configure_target,omitempty" json:"post_configure_target,omitempty"`

	// Operation to notify the source node of a target node being added via a relationship.
	AddTarget tosca.OperationDefinition `yaml:"add_target,omitempty" json:"add_target,omitempty"`

	// Operation to notify the target node of a source node which is now available via a relationship.
	AddSource tosca.OperationDefinition `yaml:"add_source,omitempty" json:"add_source,omitempty"`

	// Operation to notify source some property or attribute of the target changed.
	TargetChanged tosca.OperationDefinition `yaml:"target_changed,omitempty" json:"target_changed,omitempty"`

	// Operation to remove a target node.
	RemoveTarget tosca.OperationDefinition `yaml:"remove_target,omitempty" json:"remove_target,omitempty"`
}
