package policies

// This is the default (root) TOSCA Policy Type definition that all other TOSCA base Policy Types derive from.
type Root struct {
}

// This is the default (root) TOSCA Policy Type definition that is used to govern placement of TOSCA nodes or groups of nodes.
type Placement Root

// This is the default (root) TOSCA Policy Type definition that is used to govern scaling of TOSCA nodes or groups of nodes.
type Scaling Root

// This is the default (root) TOSCA Policy Type definition that is used to govern update of TOSCA nodes or groups of nodes.
type Update Root

// This is the default (root) TOSCA Policy Type definition that is used to declare performance requirements for TOSCA nodes or groups of nodes.
type Performance Root
