package tosca

type ImportDefinition struct {

	// imports can relate to profiles OR service template

	// single-line grammar assumes url keyword is used
	// multi-line grammer requires named parameters

	Url        string `yaml:"url,omitempty" json:"url,omitempty"`               // [conditional] The url that references a service template to be imported. An import statement must include either a url or a profile, but not both. If the value doesn't start with 'file:' or 'https', a relative path is assumed. If the url has a leading slash, the path name starts at the root of the repository, else a relative path to importing documents location.
	Profile    string `yaml:"profile,omitempty" json:"profile,omitempty"`       // [conditional] The profile name that references a named type profile to be imported. An import statement must include either a url or a profile, but not both.
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty"` // [conditional] The optional symbolic name of the repository definition where the imported file can be found as a string. The repository name can only be used when a url is specified. If used, url refers to path starting at root of named repository.
	Namespace  string `yaml:"namespace,omitempty" json:"namespace,omitempty"`   // The optional namespace into which to import the type definitions from the imported template or profile. Nested namespaces possible.
}
