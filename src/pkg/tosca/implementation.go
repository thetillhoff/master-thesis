package tosca

// used in operations and notifications, [4.3.6.8]
type ImplementationDefinition struct {

	// short notation:
	// implementation: <primary_artifact_name>

	// short notation for use with multiple artifacts:
	// implementation:
	//   primary: <primary_artifact_name>
	// 	 dependencies:
	//     - <list_of_dependent_artifact_names>
	//   timeout: 60

	// extended notation for use with single artifact:
	// implementation:
	//   primary:
	//     <primary_artifact_definition>
	//   timeout: 100

	// extended notation for use with multiple artifacts:
	// implementation:
	//   primary:
	//     <primary_artifact_definition>
	//   dependencies:
	//     - <list_of_dependent_artifact definitions>
	//   timeout: 120

	// The optional implementation artifact (i.e., the primary script file within a TOSCA CSAR file).
	Primary ArtifactDefinition `yaml:"primary,omitempty" json:"primary,omitempty"`

	// The optional list of one or more dependent or secondary implementation artifacts which are referenced by the primary implementation artifact (e.g., a library the script installs or a secondary script).
	Dependencies []ArtifactDefinition `yaml:"dependencies,omitempty" json:"dependencies,omitempty"`

	// Timeout value in seconds. Has no meaning and should not be used within a notification implementation definition.
	Timeout int `yaml:"timeout,omitempty" json:"timeout,omitempty"`
}
