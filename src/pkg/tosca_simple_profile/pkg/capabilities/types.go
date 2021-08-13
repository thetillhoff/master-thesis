package capabilities

import (
	"github.com/thetillhoff/eat/pkg/tosca"
	"github.com/thetillhoff/eat/pkg/tosca_simple_profile/pkg/data"
)

// This is the default (root) TOSCA Capability Type definition that all other TOSCA Capability Types derive from.
type Root struct {
}

// The Node capability indicates the base capabilities of a TOSCA Node Type.
type Node Root

// The Compute capability, when included on a Node Type or Template definition, indicates that the node can provide hosting on a named compute resource.
type Compute struct {
	Container `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The otional name (or identifier) of a specific compute resource for hosting.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// Number of (actual or virtual) CPUs associated with the Compute node.
	//
	// greater_or_equal: 1
	NumCpus tosca.Integer `yaml:"num_cpus,omitempty" json:"num_cpus,omitempty"`

	// Specifies the operating frequency of CPU's core.  This property expresses the expected frequency of one (1) CPU as provided by the property “num_cpus”.
	//
	// greater_or_equal: 0.1 GHz
	CpuFrequency tosca.Frequency `yaml:"cpu_frequency,omitempty" json:"cpu_frequency,omitempty"`

	// Size of the local disk available to applications running on the Compute node (default unit is MB).
	//
	// greater_or_equal: 0 MB
	DiskSize tosca.Size `yaml:"disk_size,omitempty" json:"disk_size,omitempty"`

	// Size of memory available to applications running on the Compute node (default unit is MB).
	//
	// greater_or_equal: 0 MB
	MemSize tosca.Size `yaml:"mem_size,omitempty" json:"mem_size,omitempty"`
}

// The Storage capability, when included on a Node Type or Template definition, indicates that the node can provide addressiblity for the resource a named network with the specified ports.
type Network struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The otional name (or identifier) of a specific network resource.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
}

// The Storage capability, when included on a Node Type or Template definition, indicates that the node can provide a named storage location with specified size range.
type Storage struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The otional name (or identifier) of a specific storage resource.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
}

// The Container capability, when included on a Node Type or Template definition, indicates that the node can act as a container for (or a host for) one or more other declared Node Types.
type Container Root

// This is the default TOSCA type that should be used or extended to define a network endpoint capability. This includes the information to express a basic endpoint with a single port or a complex endpoint with multiple ports.  By default the Endpoint is assumed to represent an address on a private network unless otherwise specified.
//
// Although both the port and ports properties are not required, one of port or ports must be provided in a valid Endpoint.
type Endpoint struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The name of the protocol (i.e., the protocol prefix) that the endpoint accepts (any OSI Layer 4-7 protocols)
	//
	//Examples: http, https, ftp, tcp, udp, etc.
	//
	// default: tcp
	Protocol string `yaml:"protocol,omitempty" json:"protocol,omitempty"`

	// The optional port of the endpoint.
	//
	// greater_or_equal: 1
	//
	// less_or_equal: 65535
	Port data.NetworkPortDef `yaml:"port,omitempty" json:"port,omitempty"`

	// Requests for the endpoint to be secure and use credentials supplied on the ConnectsTo relationship.
	//
	// default: false
	Secure bool `yaml:"secure,omitempty" json:"secure,omitempty"`

	// The optional URL path of the endpoint’s address if applicable for the protocol.
	UrlPath string `yaml:"url_path,omitempty" json:"url_path,omitempty"`

	// The optional name (or ID) of the network this endpoint should be bound to.
	//
	// network_name: PRIVATE | PUBLIC |<network_name> | <network_id>
	PortName string `yaml:"port_name,omitempty" json:"port_name,omitempty"`

	// default: PRIVATE
	NetworkName string `yaml:"network_name,omitempty" json:"network_name,omitempty"`

	// The optional indicator of the direction of the connection.
	//
	// one of:
	//
	// - source
	//
	// - target
	//
	// - peer
	//
	// default: source
	Inititator string `yaml:"initiator,omitempty" json:"initiator,omitempty"`

	// The optional map of ports the Endpoint supports (if more than one)
	Ports map[string]data.NetworkPortSpec `yaml:"ports,omitempty" json:"ports,omitempty"`

	// Note: This is the IP address as propagated up by the associated node’s host (Compute) container.
	IpAddress string `yaml:"ip_address,omitempty" json:"ip_address,omitempty"`
}

// This capability represents a public endpoint which is accessible to the general internet (and its public IP address ranges).
//
// This public endpoint capability also can be used to create a floating (IP) address that the underlying network assigns from a pool allocated from the application’s underlying public network.  This floating address is managed by the underlying network such that can be routed an application’s private address and remains reliable to internet clients.
//
// If the network_name is set to the reserved value PRIVATE or if the value is set to the name of network (or subnetwork) that is not public (i.e., has non-public IP address ranges assigned to it) then TOSCA Orchestrators SHALL treat this as an error.
//
// If a dns_name is set, TOSCA Orchestrators SHALL attempt to register the name in the (local) DNS registry for the Cloud provider.
type EndPointPublic struct {
	Endpoint `yaml:",inline,omitempty" json:",inline,omitempty"`
}

// This is the default TOSCA type that should be used or extended to define a specialized administrator endpoint capability.
//
// TOSCA Orchestrator implementations of Endpoint.Admin (and connections to it) SHALL assure that network-level security is enforced if possible.
type EndPointAdmin struct {
	Endpoint `yaml:",inline,omitempty" json:",inline,omitempty"`
}

// This is the default TOSCA type that should be used or extended to define a specialized database endpoint capability.
type EndPointDatabase struct {
	Endpoint `yaml:",inline,omitempty" json:",inline,omitempty"`
}

// This is the default TOSCA type that should be used or extended to define an attachment capability of a (logical) infrastructure device node (e.g., BlockStorage node).
type Attachment struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`
}

// This is the default TOSCA type that should be used to express an Operating System capability for a node.
//
// Please note that the string values for the properties architecture, type and distribution SHALL be normalized to lowercase by processors of the service template for matching purposes.  For example, if a “type” value is set to either “Linux”, “LINUX” or “linux” in a service template, the processor would normalize all three values to “linux” for matching purposes.
type OperationSystem struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The Operating System (OS) architecture.
	//
	// Examples of valid values include:
	//
	// x86_32, x86_64, etc.
	Architecture string `yaml:"architecture,omitempty" json:"architecture,omitempty"`

	// The Operating System (OS) type.
	//
	// Examples of valid values include:
	//
	// linux, aix, mac, windows, etc.
	Type string `yaml:"type,omitempty" json:"type,omitempty"`

	// The Operating System (OS) distribution.
	//
	// Examples of valid values for an “type” of “Linux” would include:
	// debian, fedora, rhel and ubuntu.
	Distribution string `yaml:"distribution,omitempty" json:"distribution,omitempty"`

	// The Operating System version.
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

// This is the default TOSCA type that should be used to express a scalability capability for a node.
//
// The actual number of instances for a node may be governed by a separate scaling policy which conceptually would be associated to either a scaling-capable node or a group of nodes in which it is defined to be a part of.  This is a planned future feature of the TOSCA Simple Profile and not currently described.
type Scalable struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// [mandatory] This property is used to indicate the minimum number of instances that should be created for the associated TOSCA Node Template by a TOSCA orchestrator.
	//
	// default: 1
	MinInstances int `yaml:"min_instances,omitempty" json:"min_instances,omitempty"`

	// [mandatory] This property is used to indicate the maximum number of instances that should be created for the associated TOSCA Node Template by a TOSCA orchestrator.
	//
	// default: 1
	MaxInstances int `yaml:"max_instances,omitempty" json:"max_instances,omitempty"`

	// An optional property that indicates the requested default number of instances that should be the starting number of instances a TOSCA orchestrator should attempt to allocate.
	// Note: The value for this property MUST be in the range between the values set for ‘min_instances’ and ‘max_instances’ properties.
	DefaultInstances int `yaml:"default_instances,omitempty" json:"default_instances,omitempty"`
}

// A node type that includes the Bindable capability indicates that it can be bound to a logical network association via a network port.
type NetworkBindable Node
