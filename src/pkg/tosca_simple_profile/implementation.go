package tosca_simple_profile

import "github.com/thetillhoff/eat/pkg/tosca"

type ImplementationDefinition struct {
	tosca.ImplementationDefinition `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The node on which operations should be executed (for TOSCA call_operation activities).
	//
	// If the operation is associated with an interface on a node type or a relationship template, valid values are SELF or HOST – referring to the node itself or to the node that is the target of the HostedOn relationship for that node.
	//
	// If the operation is associated with a relationship type or a relationship template, valid values are SOURCE or TARGET – referring to the relationship source or target node.
	//
	// In both cases, the value can also be set to ORCHESTRATOR to indicated that the operation must be executed in the orchestrator environment rather than within the context of the service being orchestrated.
	OperationHost OperationHost `yaml:"operation_host,omitempty" json:"operation_host,omitempty"`
}

type OperationHost string

const (
	OperationHostSelf   OperationHost = "SELF"
	OperationHostHost   OperationHost = "HOST"
	OperationHostSource OperationHost = "SOURCE"
	OperationHostTarget OperationHost = "TARGET"
)
