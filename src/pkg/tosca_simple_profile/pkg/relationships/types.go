package relationships

import "github.com/thetillhoff/eat/pkg/tosca_simple_profile/pkg/data"

// This is the default (root) TOSCA Relationship Type definition that all other TOSCA Relationship Types derive from.
type Root struct {
	// A unique identifier of the realized instance of a Relationship Template that derives from any TOSCA normative type.
	ToscaId string `yaml:"tosca_id,omitempty" json:"tosca_id,omitempty"`

	// This attribute reflects the name of the Relationship Template as defined in the TOSCA service template.  This name is not unique to the realized instance model of corresponding deployed application as each template in the model can result in one or more instances (e.g., scaled) when orchestrated to a provider environment.
	ToscaName string `yaml:"tosca_name,omitempty" json:"tosca_name,omitempty"`

	// The state of the relationship instance.
	// TODO: See section “Relationship States” for allowed values.
	State string `yaml:"state,omitempty" json:"state,omitempty"`
}

// This type represents a general dependency relationship between two nodes.
type DependsOn Root

// This type represents a hosting relationship between two nodes.
type HostedOn Root

// This type represents a network connection relationship between two nodes.
type ConnectsTo struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// The security credential to use to present to the target endpoint to for either authentication or authorization purposes.
	Credential data.Credential `yaml:"credential,omitempty" json:"credential,omitempty"`
}

// This type represents an attachment relationship between two nodes.  For example, an AttachesTo relationship type would be used for attaching a storage node to a Compute node.
type AttachesTo struct {
	Root `yaml:",inline,omitempty" json:",inline,omitempty"`

	// [mandatory] The relative location (e.g., path on the file system), which provides the root location to address an attached node.
	//
	// e.g., a mount point / path such as ‘/usr/data’
	//
	// Note: The user must provide it and it cannot be “root”.
	Location string `yaml:"location,omitempty" json:"location,omitempty"`

	// The logical device name which for the attached device (which is represented by the target node in the model).
	// e.g., ‘/dev/hda1’
	//
	// OR
	//
	// The logical name of the device as exposed to the instance.
	//
	// Note: A runtime property that gets set when the model gets instantiated by the orchestrator.
	Device string `yaml:"device,omitempty" json:"device,omitempty"`
}

// This type represents an intentional network routing between two Endpoints in different networks.
type RoutesTo ConnectsTo
