package tosca

type ArtifactType struct {
	abstractType

	// grammar
	// <artifact_type_name>:
	//   derived_from: <parent_artifact_type_name>
	//   version: <version_number>
	//   metadata:
	// 	   <map of string>
	//   description: <artifact_description>
	//   mime_type: <mime_type_string>
	//   file_ext: [ <file_extensions> ]
	//   properties:
	// 	   <property_definitions>

	// The optional mime type property for the Artifact Type.
	MimeType string `yaml:"mime_type,omitempty" json:"mime_type,omitempty"`

	// The optional file extension property for the Artifact Type.
	FileExt []string `yaml:"file_ext,omitempty" json:"file_ext,omitempty"`

	// An optional map of property definitions for the Artifact Type.
	Properties map[string]PropertyDefinition `yaml:"properties,omitempty" json:"properties,omitempty"`
}

type ArtifactDefinition struct {

	// short notation:
	// <artifact_name>: <artifact_file_URI>

	// extended notation:
	// <artifact_name>:
	// 	 description: <artifact_description>
	// 	 type: <artifact_type_name>
	// 	 file: <artifact_file_URI>
	// 	 repository: <artifact_repository_name>
	// 	 deploy_path: <file_deployment_path>
	// 	 version: <artifact _version>
	// 	 checksum: <artifact_checksum>
	// 	 checksum_algorithm: <artifact_checksum_algorithm>
	// 	 properties: <property assignments>

	// The mandatory artifact type for the artifact definition.
	ArtifactType string `yaml:"artifact_type,omitempty" json:"artifact_type,omitempty"`

	// The mandatory URI string (relative or absolute) which can be used to locate the artifact’s file.
	File string `yaml:"file,omitempty" json:"file,omitempty"`

	// The optional name of the repository definition which contains the location of the external repository that contains the artifact.  The artifact is expected to be referenceable by its file URI within the repository.
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`

	// The optional description for the artifact definition.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// The file path the associated file will be deployed on within the target node’s container.
	DeployPath string `yaml:"deploy_path,omitempty" json:"deploy_path,omitempty"`

	// The version of this artifact. One use of this artifact_version is to declare the particular version of this artifact type, in addition to its mime_type (that is declared in the artifact type definition). Together with the mime_type it may be used to select a particular artifact processor for this artifact. For example, a python interpreter that can interpret python version 2.7.0.
	ArtifactVersion string `yaml:"artifact_version,omitempty" json:"artifact_version,omitempty"`

	// The checksum used to validate the integrity of the artifact.
	Checksum string `yaml:"checksum,omitempty" json:"checksum,omitempty"`

	// Algorithm used to calculate the artifact checksum (e.g. MD5, SHA). Shall be specified if checksum is specified for an artifact.
	ChecksumAlgorithm string `yaml:"checksum_algorithm,omitempty" json:"checksum_algorithm,omitempty"`

	// The optional map of property assignments associated with the artifact.
	Properties map[string]interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`
}

// Custom unmarshaller, since both single-line and multi-line grammar have to be supported
func (artifactDefinition *ArtifactDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var (
		file string
		err  error

		multilineArtifactDefinition struct { // Basically the same as ArtifactDefinition, but without a custom unmarshaller.
			ArtifactType      string                 `yaml:"artifact_type,omitempty" json:"artifact_type,omitempty"`
			File              string                 `yaml:"file,omitempty" json:"file,omitempty"`
			Repository        string                 `yaml:"repository,omitempty" json:"repository,omitempty"`
			Description       string                 `yaml:"description,omitempty" json:"description,omitempty"`
			DeployPath        string                 `yaml:"deploy_path,omitempty" json:"deploy_path,omitempty"`
			ArtifactVersion   string                 `yaml:"artifact_version,omitempty" json:"artifact_version,omitempty"`
			Checksum          string                 `yaml:"checksum,omitempty" json:"checksum,omitempty"`
			ChecksumAlgorithm string                 `yaml:"checksum_algorithm,omitempty" json:"checksum_algorithm,omitempty"`
			Properties        map[string]interface{} `yaml:"properties,omitempty" json:"properties,omitempty"`
		}
	)

	// Try single-line grammar
	err = unmarshal(&file)
	if err == nil {
		artifactDefinition.File = file
		return nil
	}

	// Try multi-line grammar
	err = unmarshal(&multilineArtifactDefinition)
	if err == nil {
		artifactDefinition.ArtifactType = multilineArtifactDefinition.ArtifactType
		artifactDefinition.File = multilineArtifactDefinition.File
		artifactDefinition.Repository = multilineArtifactDefinition.Repository
		artifactDefinition.Description = multilineArtifactDefinition.Description
		artifactDefinition.DeployPath = multilineArtifactDefinition.DeployPath
		artifactDefinition.ArtifactVersion = multilineArtifactDefinition.ArtifactVersion
		artifactDefinition.Checksum = multilineArtifactDefinition.Checksum
		artifactDefinition.ChecksumAlgorithm = multilineArtifactDefinition.ChecksumAlgorithm
		artifactDefinition.Properties = multilineArtifactDefinition.Properties
		return nil
	}

	return err
}
