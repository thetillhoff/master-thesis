# CSAR

> These informations are mostly copied from the TOSCA specification.

A CSAR is a zip file where TOSCA definitions along with all accompanying artifacts (e.g. scripts, binaries, configuration files) can be packaged together. The zip file format shall conform to the Document Container File format as defined in the ISO/IEC 21320-1 "Document Container File — Part 1: Core" standard [ISO-IEC-21320-1](https://docs.oasis-open.org/tosca/TOSCA/v2.0/csd03/TOSCA-v2.0-csd03.html#CIT_ISO_IEC_21320_1).

A CSAR zip file MUST contain one of the following:
- A TOSCA.meta metadata file that provides entry information for a TOSCA orchestrator processing the CSAR file. The TOSCA.meta file may be located either at the root of the archive or inside a TOSCA-Metadata directory (the directory being at the root of the archive). The CSAR may contain only one TOSCA.meta file.
- a yaml (.yml or .yaml) file at the root of the archive, the yaml file being a valid tosca definition template.

The CSAR file MAY contain other directories and files with arbitrary names and contents.

## TOSCA Meta File

A TOSCA meta file consists of name/value pairs. The name-part of a name/value pair is followed by a colon, followed by a blank, followed by the value-part of the name/value pair. The name MUST NOT contain a colon. Values that represent binary data MUST be base64 encoded. Values that extend beyond one line can be spread over multiple lines if each subsequent line starts with at least one space. Such spaces are then collapsed when the value string is read.
```
<name>: <value>
```
Each name/value pair is in a separate line. A list of related name/value pairs, i.e. a list of consecutive name/value pairs is called a block. Blocks are separated by an empty line. The first block, called block_0, contains metadata about the CSAR itself and is further defined below. Other blocks may be used to represent custom generic metadata or metadata pertaining to files in the CSAR. A TOSCA.meta file is only required to include block_0. The structure of block_0 in the TOSCA meta file is as follows:
```
CSAR-Version: digit.digit
Created-By: string
Entry-Definitions: string
Other-Definitions: string
```
The name/value pairs are as follows:
- CSAR-Version: This is the version number of the CSAR specification. It defines the structure of the CSAR and the format of the TOSCA.meta file. The value MUST be "2.0" for this version of the CSAR specification.
- Created-By: The person or organization that created the CSAR.
- Entry-Definitions: This references the TOSCA definitions file that SHOULD be used as entry point for processing the contents of the CSAR (e.g. the main TOSCA service template).
- Other-Definitions: This references an unambiguous set of files containing substitution templates that can be used to implement nodes defined in the main template (i.e. the file declared in Entry-Definitions). Thus, all the topology templates defined in files listed under the Other-Definitions key are to be used only as substitution templates, and not as standalone services. If such a topology template cannot act as a substitution template, it will be ignored by the orchestrator. The value of the Other-Definitions key is a string containing a list of filenames (relative to the root of the CSAR archive) delimited by a blank space. If the filenames contain blank spaces, the filename should be enclosed by double quotation marks (")

Note that any further TOSCA definitions files required by the definitions specified by Entry-Definitions or Other-Definitions can be found by a TOSCA orchestrator by processing respective imports statements. Note also that artifact files (e.g. scripts, binaries, configuration files) used by the TOSCA definitions and included in the CSAR are fully described and referred via relative path names in artifact definitions in the respective TOSCA definitions files contained in the CSAR.

## Archive without TOSCA-Metadata
In case the archive doesn’t contains a TOSCA.meta file the archive is required to contains a single YAML file at the root of the archive (other templates may exist in sub-directories).

TOSCA processors should recognize this file as being the CSAR Entry-Definitions file. The CSAR-Version is inferred from the tosca_definitions_version keyname in the Entry-Definitions file. For tosca_definitions_version: tosca_2_0 and onwards, the corresponding CSAR-version is 2.0 unless further defined.

Note that in a CSAR without TOSCA-metadata it is not possible to unambiguously include definitions for substitution templates as we can have only one topology template defined in a yaml file.

Example entry file:
```
tosca_definitions_version: tosca_2_0
 
metadata:
  template_name: my_template
  template_author: OASIS TOSCA TC
  template_version: 1.0
```

## TODO
- Currently, validating a CSAR-archive also prints it agains afterwards. While this helps with manual checks, it might be not so useful in the future.
  -> When proper validation takes place (more than just parsing the yaml) remove the printouts and print what was checked instead.
