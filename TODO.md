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


## general code
- "Orchestrators MUST" are not yet implemented. Example in 5.3.8.4 of simple-profile spec
- "Orchestrators SHALL" are never implemented. Example in 5.3.6.3 of simple-profile spec

## code (tosca)
- [ ] replace '“' with '"'
- [ ] replace all <double spaces> with <tabs>
- [ ] where to place conditions? [4.6.7.6ff]
- tosca spec [4.8.4] why the hell is this here and not in simple-profile?
- [ ] add examples to service_template.yaml
  - 5.4.1.3
  - 5.4.2.3
- check SubstitutionMapping
- status as const
- search for _name and check whether that should be added as variable to those structs
- OtherServiceTemplates of CSAR?
- "AssertionDefinition" not used
- "AttributeAssignment" not used
- check if abstractType is used everywhere it should
- "ParameterAssignment" and "ParameterMappingAssignment" not used
- "InterfaceMapping" unused
- dsl_definitions are completely ignored (intended by spec, but worth noting specifically)
- PropertyFilterDefinition not implemented
- on mandatory ("[mandatory]") attributes remove the omitempty
- [5.1] "Reserved Function Keywords" not implemented
- [5.2] "Reserved Environment Variable Names and Usage" not implemented
- [5.3] "Intrinsic Functions" not implemented (functions.go)
- Add normative types like string, list etc as "default" data types

## issues in tosca spec
- 4.4.7.2 inconsistent "external-schema"
- 4.3.5.6.3.3 contains indentation error
- 5.2.1 yaml-snippet contains indentation error
- 4.2.6.2.7.2 contains invalid yaml ("constraints: equals: ...")
- 4.2.1.3.16.3 missing capability definition name("mytypes.myfeatures.transactSQL")
- 5.3.1.3 missing output name
- 4.5.5.2 inconsistent type of "properties" and "attributes", only for requirement mappings these are lists, else maps
- 4.2.6.2.7.2 contains indentation error
- while in [4.4] / [4.4.1] primitive types are described, the implementation is partly influenced by the simple-profile spec [5.3]


## code (simple profile)
- Simple profile uses yaml-timestamps. tosca has own timestamp types. Decide for one.
- See issue on "operation implementation" and "notification implementation", omitted the distinction since all values are nullable.
- on mandatory ("[mandatory]") attributes remove the omitempty
- [dataTypes.yaml] check metadata and other initial stuff
- *dataTypes often have default values or "only this is allowed" - this is not implemented yet
- 5.8.4.3 TOSCA Orchestrator processing of Deployment artifacts - what?

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
- 5.8.4.2ff up to 5.8.4.4.2 is incomplete (and the only provided diagram seems to be faulty)
- 5.8.5 missing property/attribute descriptions
- 5.8.5.2 missing illustration
- 5.9.8 1:1 connection between database and user-password combination. Additionally: Why does the Credential type exist, when username and password are then dedicated attributes. (Also observed in other occurences. Just search for "password" or "user")
- there was one occurence where a constraint was larger_or_equal than 0.1 MB or something like that. Why no 0B?
  example at 5.9.10.1
- 5.9.9.2 capabilities are TBD
- 5.9.1.3 contains indentation error
- 5.7.1.2 contains indentation error & derive_from is not defined
- sometimes the map[string]* forces for name finding, but its sometimes unnecessary (as it has a description). While it allows for reference elsewhere, it should be described clearly where this is possible and whats the use-case is.
- 5.8.3 is the first one to contain a type definitions, earlier subchapters don't contain types - this is inconsistent...


## code (csar?)
- when parsing a serviceTemplate, what happens when stuff is ignored? Can UnMarshallStrict help here?
  Example: dsl_definitions

## goal:
- grow infrastructure with minimal manual effort
- deploy/restore fully automated