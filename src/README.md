# `eat`

...

## usage
`*` represents a file/folder

`eat *.zip` -> read zipfile contents, expect CSAR format
`eat *`, `eat */` -> read folder contents, expect contents of CSAR

`eat validate *` -> read input, parse and validate contents

`eat install *`

`eat uninstall *`

? `eat * --dry-run`

## architecture
The package `tosca` contains *only* the structs and general methods.
The package `csar` contains *only* the "parser" to read from files and parse to tosca instances.
The package `tosca_simple_profile` contains the equivalent of the tosca simple profile spec. Depends on the `tosca` package.
The package `tosca_hw_extension` contains all newly added content. Depends on the `tosca` package.
