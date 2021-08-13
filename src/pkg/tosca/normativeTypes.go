package tosca

// The recognized node states for TOSCA that will be set by the orchestrator to describe a node instance’s state.
type NodeState int

const (
	NodeStateInitial     NodeState = iota
	NodeStateCreating    NodeState = iota
	NodeStateCreated     NodeState = iota
	NodeStateConfiguring NodeState = iota
	NodeStateConfigured  NodeState = iota
	NodeStateStarting    NodeState = iota
	NodeStateStarted     NodeState = iota
	NodeStateStopping    NodeState = iota
	NodeStateDeleting    NodeState = iota
	NodeStateError       NodeState = iota
)

// The recognized relationship states for TOSCA that will be set by the orchestrator to describe a node instance’s state.
//
// Note: Additional states may be defined in future versions of the TOSCA specification.
type RelationshipState int

const (
	RelationshipStateInitial RelationshipState = iota
)

// The directive values defined for this version of TOSCA.
type Directive int

const (
	DirectiveSubstitute Directive = iota
	DirectiveSelect     Directive = iota
)

// The recognized values that may be used as aliases to reference types of networks within an application model without knowing their actual name (or identifier) which may be assigned by the underlying Cloud platform at runtime.
type Network string

const (
	NetworkPrivate Network = "PRIVATE"
	NetworkPublic  Network = "PUBLIC"
)
