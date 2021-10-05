# TODO

## presentation

- live-demo
  - vagrant set up or alternative

## weekly BG
- livedemo with real deployment instead of dummy

## weekly ST
- notes:
  - less tooling more features
  - Hypervisors/VMs?

## self
- write down notes on story with cluster API, tinkerbell, OPA etc.
- [related work] what's the problem with opentosca, cloudify, ...?
- [outlook] QR-code for mgmt MAC whitelisting or whitelist _all_ MACs by default
- migrate tosca notes to document, print that as reference
- Check and document which levels were achieved:
  Check qualification of implementaiton at https://docs.oasis-open.org/tosca/TOSCA/v2.0/csd03/TOSCA-v2.0-csd03.html#_Toc56506177
  Check Conformance targets at https://docs.oasis-open.org/tosca/TOSCA/v2.0/csd03/TOSCA-v2.0-csd03.html#_Toc56506774
- Check and document which levels were achieved:
  Check qualification of implementation at https://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.3/os/TOSCA-Simple-Profile-YAML-v1.3-os.html#_Toc26969414
  Also check conformance targets at https://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.3/os/TOSCA-Simple-Profile-YAML-v1.3-os.html#_Toc26969508

## tex
- replace "f.e." with "for example"
- replace "\&" with "and"
- unify license and licence
- unify bare-metal and bare metal
- add more stuff of IaC book (available at oreilly with hs access) - f.e. chapter Challenges and Principles
- [DSL for IaC] page 1, problems arising while deploying infrastructures in cloud
  - heterogeneity of cloud resources
  - lack of compatibility of cloud resource model and api in different clouds
  - lack of integrations between different cloud environments
  - transferring applications between clouds causes unreasonable complexity and errors
  - multiple infrastructure deployments demand too much humand resources

## general code
- "Orchestrators MUST" are not yet implemented. Example in 5.3.8.4 of simple-profile spec
- "Orchestrators SHALL" are never implemented. Example in 5.3.6.3 of simple-profile spec

## code (tosca)
- [x] replace '"' with '"'
- [?] replace all <double spaces> with <tabs>
- [ ] 4.3.5.8 PropertyFilterDefinition is not implemented
- [ ] where to place conditions? [4.6.7.6ff]
- tosca spec [4.8.4] why the hell is this here and not in simple-profile?
- [ ] add examples to service_template.yaml
  - 5.4.1.3
  - 5.4.2.3
- [ ] check SubstitutionMapping
- [ ] search for _name and check whether that should be added as variable to those structs
- "AssertionDefinition" not used
- "AttributeAssignment" not used
- "ParameterAssignment" and "ParameterMappingAssignment" not used
- "InterfaceMapping" unused
- dsl_definitions are completely ignored (intended by spec, but worth noting specifically)
- PropertyFilterDefinition not implemented
- [x] on mandatory ("[mandatory]") attributes remove the omitempty
- [5.1] "Reserved Function Keywords" not implemented
- [5.2] "Reserved Environment Variable Names and Usage" not implemented
- [5.3] "Intrinsic Functions" not implemented (functions.go)
- [ ] Add normative types like string, list etc as "default" data types
- [ ] Test with metadata not embedded in service template but in dedicated file
- `Equal` should consider constraints etc as well!
- [ ] Size.Parse uses bytefmt which declares f.e. TB==TiB which is not true -> reimplement, but too much effort for MA
- [x] Imports not resolved right now -> should be done, and resolving derivation should use this as well
- [ ] InterfaceAssignment is never used?
- [ ] Are Assignments the same as "storing a value according to definition"? This would result in making all of them interfaces...
- [ ] As stated in spec 4.4.1, 64-bit precision is recommended for all values
- [ ] It would be great to work with pointers instead of direct variables (at least at some points). This would later enable checking against nil.

## issues in tosca spec
- 4.3.5.6.3.3 contains indentation error
- 4.2.1.3.16.3 missing capability definition name("mytypes.myfeatures.transactSQL")
- 4.2.6.2.7.2 contains invalid yaml ("constraints: equals: ...")
- 4.2.6.2.7.2 contains indentation error
- while in [4.4] / [4.4.1] primitive types are described, the implementation is partly influenced by the simple-profile spec [5.3]
- 4.4.7.2 inconsistent "external-schema"
- 4.5.5.2 inconsistent type of "properties" and "attributes", only for requirement mappings these are lists, else maps
- 5.2.1 yaml-snippet contains indentation error
- 5.3.1.3 missing output name
- 4.4.2 get_property is defined with "[entity_name, optional_req_or_cap_name, property_name, nested_property_name_or_index]"
  It is not possible to detect in a sane way whether optional_req_or_cap_name or nested_property_name_or_index is set.
  Therefore it is assumed, that optional_req_or_cap_name is NEVER set.
<!-- - 3.8.3.1 and 3.8.3.2 contain conflicting information about the content of the field 'interfaces': Is it interface_definitions or interface_assignments? -->
- There are quite many assumptions to make when working with nodeTemplates (same goes for relationshipTempaltes) - mostly related with how properties etc are retrieved from nodeType and relationshipType


<!-- ## code (simple profile)
- Simple profile uses yaml-timestamps. tosca has own timestamp types. Decide for one.
- See issue on "operation implementation" and "notification implementation", omitted the distinction since all values are nullable.
- on mandatory ("[mandatory]") attributes remove the omitempty
- [dataTypes.yaml] check metadata and other initial stuff
- *dataTypes often have default values or "only this is allowed" - this is not implemented yet
- 5.8.4.3 TOSCA Orchestrator processing of Deployment artifacts - what? -->

## issues in simple profile spec
- 2.1 invalid tosca_definitions_version
- most of stuff in chapter 3 is already defined in the tosca spec -> skipped over some parts
- simple profile "operation implementation" has attribute "operation_host", while the tosca-part of it doesn't
- simple profile "operation implementation" and "notification implementation" are completely distinct, while in tosca they are the same.
- "interface definition" keynames doesn't contain "type", but is listed in grammar notations. Since tosca-spec contains the same _with_ type, left as is.
- 5.3.4.2.2 bad indentation (3 spaces)
- datatypes sometimes start with capital letters, sometimes not (Root, json, xml, Credential) - they are case sensitive, so this feels inconsistent
- 5.3.6.5.1 contains indentation error
- 5.3.6.5.4 is followed by 5.3.6.5.6 - WTF
- 5.3.8 - 5.3.10 don't contain the column "required" - assuming none are required
- 5.4 first line says thre are three categories of artifacts. Listed are 4.
- 5.7.1.2 contains indentation error & derive_from is not defined
- 5.8.3 is the first one to contain a type definitions, earlier subchapters don't contain types - this is inconsistent...
- 5.8.4.2ff up to 5.8.4.4.2 is incomplete (and the only provided diagram seems to be faulty)
- 5.8.5 missing property/attribute descriptions
- 5.8.5.2 missing illustration
- 5.9.1.3 contains indentation error
- 5.9.8 1:1 connection between database and user-password combination. Additionally: Why does the Credential type exist, when username and password are then dedicated attributes. (Also observed in other occurences. Just search for "password" or "user")
- 5.9.9.2 capabilities are TBD
- there was one occurence where a constraint was larger_or_equal than 0.1 MB or something like that. Why no 0B?
  example at 5.9.10.1
- sometimes the map[string]* forces for name finding, but its sometimes unnecessary (as it has a description). While it allows for reference elsewhere, it should be described clearly where this is possible and whats the use-case is.
- 5.8.5 Text is copy-paste error

## code (csar)
- [ ] Resolving OtherDefinitions should:
  - parse them as service templates, resolve imports and profiles
  - resolve substitutions
- Imports will only work via absolute and relative paths. Paths beginning with "https:" and "file:" will result in an error (file doesn't exist).
  "Profile" keyword is not supported (yet).
  Sufficient for MA, but important to note somewhere.

## code tosca_orchestrator
- [ ] inputs should also be providable via .yaml file. Multiple possible (merge with override). Command-line input must also override those values.
- [-] use opa for policy/... checking (requirements, capabilities, filters, ...)
  https://www.openpolicyagent.org/docs/latest/#example
  Won't do, since tosca has its own requirements, capabilities etc. opa would simply introduce a new layer of complexity

## notes from prof
Level 0:
  App: Webshop

Level 1:
  Physical: Compute Type x1
  Admin:    SSH
  Software: Webshop v2.1
  Scale:    Europe
  Provider: Private

Level 2 (IaS):
  Physical (Compute Type x1)
    Location
      Room: Central
      Power: USV
    Compute
      CPU: 4 Cores
      Memory: 16GB
    Network (Standard)
      Speed: 1G
      Loadbalancer: Yes
    Storage
      Disk: 1TB SSD
  Software
    Code: Generic-webshop v1
    Deploy: Docker
    Database: Postgres
    Auth: Webtoken
    Userdatabase: Kerberos

-> "Mehrdimensionales Systemmodell"?


## goal:
- grow infrastructure with minimal manual effort
- deploy/restore fully automated

## code cli 
- init command generates definitions
- minify removes unused definitions
- check whether all used components have corresponding plugins available (and download them)
- `(--)list-inputs` should list the inputs of a CSAR package, with examples
  inputs are provided at runtime OR via file, help flag shows a list with available entries like
  `ram_size=<bytesize> # f.e. 4 GB`
  - defaults can be specified in tosca files
  - allow variable-files as well (.env files) - create a minimum file via cli-command?

## other
- contact tosca devs; questions:
  - why no hardware?
  - why no state?
  - why properties and attributes and types and templates and instances?
    Isn't it simpler to create templates with properties (with defaults) and instantiate those and extend the predefined properties?
    That way, whenever a new value is set, all verifications (type, range, etc) can be applied, instead of a whole lot of deriving and resolving and then still doing the same.
- when it doesn't matter whether a machine is physical or virtual, it is referred to as 'machine', else 'phyiscal machine' or 'virtual machine'.
