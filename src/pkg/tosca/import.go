package tosca

import (
	"bytes"
	"log"

	"gopkg.in/yaml.v3"
)

type ImportDefinition struct {
	// imports can relate to profiles OR service template

	// single-line grammar assumes url keyword is used
	// multi-line grammer requires named parameters

	// [conditional] The url that references a service template to be imported. An import statement must include either a url or a profile, but not both. If the value doesn't start with 'file:' or 'https', a relative path is assumed. If the url has a leading slash, the path name starts at the root of the repository, else a relative path to importing documents location.
	Url string `yaml:"url,omitempty" json:"url,omitempty"`

	// [conditional] The profile name that references a named type profile to be imported. An import statement must include either a url or a profile, but not both.
	Profile string `yaml:"profile,omitempty" json:"profile,omitempty"`

	// [conditional] The optional symbolic name of the repository definition where the imported file can be found as a string. The repository name can only be used when a url is specified. If used, url refers to path starting at root of named repository.
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`

	// The optional namespace into which to import the type definitions from the imported template or profile. Nested namespaces possible.
	//
	// The seperator between namespaces and "normal" names is the colon ':'.
	Namespace string `yaml:"namespace,omitempty" json:"namespace,omitempty"`
}

func (importDefinition ImportDefinition) ToString() string {
	var (
		err         error
		buffer      bytes.Buffer
		yamlEncoder *yaml.Encoder
	)

	yamlEncoder = yaml.NewEncoder(&buffer)
	yamlEncoder.SetIndent(2) // Default is 4 spaces
	err = yamlEncoder.Encode(&importDefinition)
	if err != nil {
		log.Fatalln(err)
	}
	defer yamlEncoder.Close()

	return buffer.String()
}

// Custom unmarshaller, since both single-line and multi-line grammar have to be supported
func (importDefinition *ImportDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var (
		url string
		err error

		multilineImportDefinition struct { // Basically the same as ImportDefinition, but without a custom unmarshaller.
			Url        string `yaml:"url,omitempty" json:"url,omitempty"`
			Profile    string `yaml:"profile,omitempty" json:"profile,omitempty"`
			Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`
			Namespace  string `yaml:"namespace,omitempty" json:"namespace,omitempty"`
		}
	)

	// Try single-line grammar
	err = unmarshal(&url)
	if err == nil {
		importDefinition.Url = url
		return nil
	}

	// Try multi-line grammar
	err = unmarshal(&multilineImportDefinition)
	if err == nil {
		importDefinition.Url = multilineImportDefinition.Url
		importDefinition.Profile = multilineImportDefinition.Profile
		importDefinition.Repository = multilineImportDefinition.Repository
		importDefinition.Namespace = multilineImportDefinition.Namespace
		return nil
	}

	return err
}
