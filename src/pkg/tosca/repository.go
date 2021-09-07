package tosca

type RepositoryDefinition struct {

	// single-line grammar assumes url keyword is used ('<repository_name>: <repository_url>')
	// multi-line grammer requires named parameters

	Description *string `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the repository.
	Url         *string `yaml:"url,omitempty" json:"url,omitempty"`                 // [mandatory] The mandatory URL or network address used to access the repository.
}
