# `eat`

...

## usage
`*` represents a file/folder

`eat *.zip` -> read zipfile contents, expect CSAR format
`eat *`, `eat */` -> read folder contents, expect contents of CSAR

`eat validate *` -> read input, parse and validate contents

`eat install *` -> read input, parse and validate contents, resolve dependecies in topologyTemplate (detect order of implementations), run implementation artifacts

`eat uninstall *`

? `eat * --dry-run`

## architecture
The package `tosca` contains *only* the structs and general methods (including validations). Thus, this package contains a "TOSCA processor". See https://docs.oasis-open.org/tosca/TOSCA/v2.0/csd03/TOSCA-v2.0-csd03.html#_Toc56506776 for more information about TOSCA processors.

The package `csar` contains *only* the "parser" to read from files and parse to tosca instances. It depends on the `tosca` package. Only service templates conforming to TOSCA YAML service templates should be recognized as valid inputs. See https://docs.oasis-open.org/tosca/TOSCA/v2.0/csd03/TOSCA-v2.0-csd03.html#_Toc56506775 for more information about TOSCA YAML service templates.

The package `tosca_orchestrator` contains the workflow processor.

The package `tosca_hw_extension` contains all newly added content. Depends on the `tosca` package.


## development
Run `go install github.com/spf13/cobra/cobra` to install the cobra cmd-tool.
To add a new command run `cobra add install` where `install` is the name of the new command.
