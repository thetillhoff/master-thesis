package nodes

import (
	"github.com/thetillhoff/eat/pkg/tosca"
	"github.com/thetillhoff/eat/pkg/tosca_simple_profile/pkg/data"
)

// The TOSCA Root Node Type is the default type that all other TOSCA base Node Types derive from.  This allows for all TOSCA nodes to have a consistent set of features for modeling and management (e.g., consistent definitions for requirements, capabilities and lifecycle interfaces).
//
// All Node Type definitions that wish to adhere to the TOSCA Simple Profile SHOULD extend from the TOSCA Root Node Type to be assured of compatibility and portability across implementations.
type Root struct {

	// A unique identifier of the realized instance of a Node Template that derives from any TOSCA normative type.
	ToscaId string `yaml:"tosca_id,omitempty" json:"tosca_id,omitempty"`

	// This attribute reflects the name of the Node Template as defined in the TOSCA service template.  This name is not unique to the realized instance model of corresponding deployed application as each template in the model can result in one or more instances (e.g., scaled) when orchestrated to a provider environment.
	ToscaName string `yaml:"tosca_name,omitempty" json:"tosca_name,omitempty"`

	// The state of the node instance.
	// TODO See section “Node States” for allowed values.
	//
	// default: initial
	State string `yaml:"state,omitempty" json:"state,omitempty"`
}

// The TOSCA Abstract.Compute node represents an abstract compute resource without any requirements on storage or network resources.
type AbstractCompute Root

// The TOSCA Compute node represents one or more real or virtual processors of software applications or services along with other essential local resources.  Collectively, the resources the compute node represents can logically be viewed as a (real or virtual) “server”.
//
// The underlying implementation of the Compute node SHOULD have the ability to instantiate guest operating systems (either actual or virtualized) based upon the OperatingSystem capability properties if they are supplied in the a node template derived from the Compute node type.
type Compute struct {
	AbstractCompute `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The primary private IP address assigned by the cloud provider that applications may use to access the Compute node.
	PrivateAddress string `yaml:"private_address,omitempty" json:"private_address,omitempty"`

	// The primary public IP address assigned by the cloud provider that applications may use to access the Compute node.
	PublicAddress string `yaml:"public_address,omitempty" json:"public_address,omitempty"`

	// The map of logical networks assigned to the compute host instance and information about them.
	Networks map[string]data.NetworkInfo `yaml:"networks,omitempty" json:"networks,omitempty"`

	// The map of logical ports assigned to the compute host instance and information about them.
	Ports map[string]data.NetworkPortInfo `yaml:"ports,omitempty" json:"ports,omitempty"`
}

// The TOSCA SoftwareComponent node represents a generic software component that can be managed and run by a TOSCA Compute Node Type.
//
// Nodes that can directly be managed and run by a TOSCA Compute Node Type SHOULD extend from this type.
type SoftwareComponent struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The optional software component’s version.
	ComponentVersion tosca.Version `yaml:"component_version,omitempty" json:"component_version,omitempty"`

	// The optional credential that can be used to authenticate to the software component.
	AdminCredential data.Credential `yaml:"admin_credential,omitempty" json:"admin_credential,omitempty"`
}

// This TOSA WebServer Node Type represents an abstract software component or service that is capable of hosting and providing management operations for one or more WebApplication nodes.
type WebServer SoftwareComponent

// The TOSCA WebApplication node represents a software application that can be managed and run by a TOSCA WebServer node.  Specific types of web applications such as Java, etc. could be derived from this type.
type WebApplication struct {
	SoftwareComponent `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The web application’s context root which designates the application’s URL path within the web server it is hosted on.
	ContextRoot string `yaml:"context_root,omitempty" json:"context_root,omitempty"`
}

// The TOSCA DBMS node represents a typical relational, SQL Database Management System software component or service.
type DBMS struct {
	SoftwareComponent `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The optional root password for the DBMS server.
	RootPassword string `yaml:"root_password,omitempty" json:"root_password,omitempty"`

	// The DBMS server’s port.
	Port int `yaml:"port,omitempty" json:"port,omitempty"`
}

// The TOSCA Database node represents a logical database that can be managed and hosted by a TOSCA DBMS node.
type Database struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The logical database Name
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// The port the database service will use to listen for incoming data and requests.
	Port int `yaml:"port,omitempty" json:"port,omitempty"`

	// The special user account used for database administration.
	User string `yaml:"user,omitempty" json:"user,omitempty"`

	// The password associated with the user account provided in the ‘user’ property.
	Password string `yaml:"password,omitempty" json:"password,omitempty"`
}

// The TOSCA Abstract.Storage node represents an abstract storage resource without any requirements on compute or network resources.
type AbstractStorage struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// [mandatory] The logical name (or ID) of the storage resource.
	Name string `yaml:"name" json:"name"`

	// The requested initial storage size (default unit is in Gigabytes).
	//
	// greater_or_equal: 0 MB
	Size tosca.Size `yaml:"size,omitempty" json:"size,omitempty"`
}

// The TOSCA ObjectStorage node represents storage that provides the ability to store data as objects (or BLOBs of data) without consideration for the underlying filesystem or devices.
//
// Subclasses of the tosca.nodes.Storage.ObjectStorage node type may impose further constraints on properties.  For example, a subclass may constrain the (minimum or maximum) length of the ‘name’ property or include a regular expression to constrain allowed characters used in the ‘name’ property.
type ObjectStorage struct {
	AbstractStorage `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The requested maximum storage size (default unit is in Gigabytes).
	//
	// greater_or_equal: 1 GB
	Maxsize tosca.Size `yaml:"maxsize,omitempty" json:"maxsize,omitempty"`
}

// The TOSCA BlockStorage node currently represents a server-local block storage device (i.e., not shared) offering evenly sized blocks of data from which raw storage volumes can be created.
//
// Note:
//
// - In this draft of the TOSCA Simple Profile, distributed or Network Attached Storage (NAS) are not yet considered (nor are clustered file systems), but the TC plans to do so in future drafts.
//
// - The size property is required when an existing volume (i.e., volume_id) is not available. However, if the property volume_id is provided, the size property is ignored.
//
// - Resize is of existing volumes is not considered at this time.
//
// - It is assumed that the volume contains a single filesystem that the operating system (that is hosting an associate application) can recognize and mount without additional information (i.e., it is operating system independent).
//
// - Currently, this version of the Simple Profile does not consider regions (or availability zones) when modeling storage.
type BlockStorage struct {
	AbstractStorage `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The requested storage size (default unit is MB).
	//
	// greater_or_equal: 1 MB
	Size tosca.Size `yaml:"size,omitempty" json:"size,omitempty"`

	// ID of an existing volume (that is in the accessible scope of the requesting application).
	VolumeId string `yaml:"volume_id,omitempty" json:"volume_id,omitempty"`

	// Some identifier that represents an existing snapshot that should be used when creating the block storage (volume).
	SnapshotId string `yaml:"snapshot_id,omitempty" json:"snapshot_id,omitempty"`
}

// The TOSCA Container Runtime node represents operating system-level virtualization technology used to run multiple application services on a single Compute host.
type ContainerRuntime SoftwareComponent

// The TOSCA Container Application node represents an application that requires Container-level virtualization technology.
type ContainerApplication Root

// The TOSCA Load Balancer node represents logical function that be used in conjunction with a Floating Address to distribute an application’s traffic (load) across a number of instances of the application (e.g., for a clustered or scaled application).
//
// Note:
//
// A LoadBalancer node can still be instantiated and managed independently of any applications it would serve; therefore, the load balancer’s application requirement allows for zero occurrences.
type LoadBalancer Root
