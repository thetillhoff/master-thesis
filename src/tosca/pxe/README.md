# PXE

The goal of this tosca package is to describe the basic pxe-featureset with ipv4.

# structure
According to CSAR (Cloud Service ARchive), the file-structure is predefined like this:
- TOSCA-Metadata
  - TOSCA.meta # key-value pairs, representing metadata of other files in CSAR, organized in blocks. Each block provides metadata of a certain artifact of the CSAR. An empty line seperates the blocks. The first block describes the CSAR as a whole. Following Names shoudl represent a relative path from the root of the package, f.e. "Name: /Somewhere/Something"
- Definitions
- Types
- Plans
- ...
- VirtualImages
- JARs

# components
- tftp-node / https-node
- dhcp-node
- os-iso
- os-auto-install-?
- node to be installed with capabilities for
  - network plugged in and all the other nodes are reachable (define reachable)
    - mac-address
    - one or more
  - ipmi and/or wol, currently powered off
  - local disk min_size for os
  - some cpu with fitting arch
  - some ram at least min for os
  - power cable attached

- plans:
  - remote-power-on(IP)
  - install-os-on-node(OS,IP)

# notes (taken literally from spec)
- instantiation of all types in simple-profile is expected to be understood by every implementation of TOSCA
  - tosca orchestrators simply select or allocate the correct node (resource) type that fulfills the application topologies requirements using the properties declared in the node and its capabilities
- tosca orchestration engines are expected to validate all property values provided in a node template against the property defintions in their respective node type definitions referenced in the service template
- 4.2.5.1 describes derivation -> "derived_from"
- A Node Template is an instance of a specified Node Type and can provider  customized properties, constraints, relationships or interfaces which complement and change the defaults provided by its Node Type.


# allowed keynames (4.2.1.1 ff of tosca-2.0)
- tosca_definitions_version: <value>  # Mandatory, must be in first line
  allowed values:
  - tosca_2_0
  - tosca_simple_yaml_1_3 <...>

- profile: <string>                   # Optional, java-style reverse-domain notation recommended (f.e. io.kubernetes:1.18)
  This is the name that will be imported in other definitions. There it can be namespaced (python "import x as y")

- metadata: # Optional metadata keyname: value pairs
  - template_name: <value>            # Optional, name of this service template
  - template_author: <value>          # Optional, author of this service template
  - template_version: <value>         # Optional, version of this service template
  #More optional entries of domain or profile specific metadata keynames
  - creation_date
  - date_updated
  - status
  #template_name and template_version should be unique

- description: <template_ description> # Optional description of the definitions inside the file.

- dsl_definitions: # map of YAML alias anchors (or macros)
  #YAML-Style macros like
    ubuntu_image_props: &ubuntu_image_props
      architecture: x86_64
      type: linux
      distribution: ubuntu
      os_version: 14.04

    redhat_image_props: &redhat_image_props
      architecture: x86_64
      type: linux
      distribution: rhel
      os_version: 6.6

- repositories: # map of external repository definitions which host TOSCA artifacts
  #define external repositories
    my_project_artifact_repo:
      description: development repository for TAR archives and Bash scripts
      url: http://mycompany.com/repository/myproject/

- imports: # ordered list of import definitions
  either relative path or url or profile (url can be combined with repository as well (-> relative within that repository))
    #An example import of definitions files from a location relative to the file location of the service template declaring the import.
    - relative_path/my_defns/my_typesdefs_1.yaml
    - url: my_defns/my_typesdefs_n.yaml   
      repository: my_company_repo
      namespace: mycompany

- artifact_types: # map of artifact type definitions
  #example:
  mycompany.artifacttypes.myFileType:
    derived_from: tosca.artifacts.File
  

- data_types: # map of datatype definitions
  #example:
    #A complex datatype definition
    simple_contactinfo_type:
      properties:
        name:
          type: string
        email:
          type: string
        phone:
          type: string
    #datatype definition derived from an existing type
    full_contact_info:
      derived_from: simple_contact_info
      properties:
        street_address:
          type: string
        city:
          type: string
        state:
          type: string
        postalcode:
          type: string

- capability_types: # map of capability type definitions
  #example
  mycompany.mytypes.myCustomEndpoint:
    derived_from: tosca.capabilities.Endpoint
    properties:
      #more details ...
  mycompany.mytypes.myCustomFeature:
    derived_from: tosca.capabilities.Feature
    properties:
      #more details ...
  mycompany.mytypes.myapplication.MyFeature:
    derived_from: tosca.capabilities.Root
    [version]: <string>
    [description]: a custom feature of my company’s application
    [properties]:
      my_feature_setting:
        type: string
      my_feature_value:
        type: integer
    [attributes]: <map_of_attributes>
    [valid_source_types]: <list_of_node_types> #list of node types that this capability type supports a s valid sources for a successful relationship to be established to itself. If undefined, no restrictions.

- interface_types: # map of interface type definitions
    mycompany.mytypes.myinterfaces.MyConfigure:
      derived_from: tosca.interfaces.relationship.Root
      [version]: <string>
      [metadata]: <map_of_string>
      [description]: My custom configure Interface Type
      [inputs]:
        mode:
          type: string
      [operations]: #operations must be defined in the interface type definition
        pre_configure_service:
          description: pre-configure operation for my service
        post_configure_service:
          description: post-configure operation for my service
      [notifications]: <map_of_notification_defintions> #must be defined in the interface type definition
    mycompany.interfaces.service.Signal:
      operations:
        signal_begin_receive:
          description: Operation to signal start of some message processing.
        signal_end_receive:
          description: Operation to signal end of some message processed.

- relationship_types: # map of relationship type definitions
  #example
    mycompany.mytypes.myCustomClientServerType:
      derived_from: tosca.relationships.HostedOn
      properties:
        #more details ...
    mycompany.mytypes.myCustomConnectionType:
      derived_from: tosca.relationships.ConnectsTo
      properties:
        #more details ...  
    mycompanytypes.myrelationships.AppDependency:
      derived_from: tosca.relationships.DependsOn
      [version]: <string>
      [metadata]: <map_of_string>
      [description]: <string>
      [properties]: <map_of_properties>
      [attributes]: <map_of_attributes>
      [interfaces]: <map_of_interfaces>
      [valid_target_types]: [ mycompanytypes.mycapabilities.SomeAppCapability ] #valid target capability types for the relationship, if undefined, the valid target types are not restricted at all. Should this be derived, all elements must be set in parent type list or derived from an element in the parten type list.

- node_types: # map of node type definitions
  #example
    my_webapp_node_type:
      derived_from: WebApplication
      properties:
        my_port:
          type: integer
    my_database_node_type:
      derived_from: Database
      capabilities:
        mytypes.myfeatures.transactSQL
    my_company.my_types.my_app_node_type: #Derived content can be redefined or new stuff added
      derived_from: tosca.nodes.SoftwareComponent
      description: My company’s custom applicaton
      version: <version_number>
      metadata:
        <map of string>
      properties:
        my_app_password:
          type: string
          description: application password
          constraints:
            - min_length: 6
            - max_length: 10
      attributes:
        my_app_port:
          type: integer
          description: application port number
      capabilities:
        # A capability defintion defines a typed set of data that a node can expose and is used to describe a relevant feature of the component described by a node. A Capability is defed part of a node type defintion and my be refined during node type derivation
        <capability_definitions>
        some_capability: mytypes.mycapabilities.MyCapabilityTypeName #short style
        some_capability: #long style
          type: mytypes.mycapabilities.MyCapabilityTypeName
          [description]: <string>
          [properties]: #no new ones
            limit:
              default: 100
          [attributes]: <map_of_attributes> #no new ones
          [valid_source_types]: <list_of_node_type_names>
          [occurences]: <range> #min/max can be defined, must be within range of parent if node is derived, default is [1, inf]
      requirements:
        - some_database:
            capability: EndPoint.Database
            node: Database   
            relationship: ConnectsTo
      interfaces:
        <interface_def_name>:
          type: <interface_type_name>
          [description]: <string>
          [inputs]: <map_of_parameter_definitions_and_refinements>
          [operations]: <map_of_operation_refinements> #TBD
          [notifications]: <map_of_notification_refinements> #TBD
        Standard:
          start: scripts/start_server.sh
        Configure:
          pre_configure_source:
            implementation:
              primary: scripts/pre_configure_source.sh
              dependencies:
                - scripts/setup.sh
                - binaries/library.rpm
                - scripts/register.py
        Configure2: #equal to above, but more verbose
          pre_configure_source:
            implementation:
              primary:
                file: scripts/pre_configure_source.sh
                type: tosca.artifacts.Implementation.Bash
                repository: my_service_catalog
              dependencies:
                - file : scripts/setup.sh
                  type : tosca.artifacts.Implementation.Bash
                  repository : my_service_catalog
        do_something:
          operations: 
            an_operation: <operation_implementation-definition>
        do_something_with_input_and_output:
          operations:
            an_extended_operation:
              description: <string>
              implementation: <operation_implementation-definition>
              inputs:
                <map_of_parameter_definitions>
              outputs:
                <map_of_parameter_definitions>
      artifacts: #deriving doesn't make sense in a lot of cases
        <artifact_definitions>

- group_types: # map of group type definitions
  #example
    mycompany.mytypes.myScalingGroup:
      derived_from: tosca.groups.Root
- policy_types: # map of policy type definitions
  #example
    mycompany.mytypes.myScalingPolicy:
      derived_from: tosca.policies.Scaling
- topology_template:
  #topology template definition of the cloud application or service
  description: <template_description>
  inputs:
    SiteName:
    type: string
    description: string typed parameter definition with constraints
    default: My Site
    constraints:
      - min_length: 9
  outputs:
    server_address:
      description: The first private IP address for the provisioned server.
      value: { get_attribute: [ nodes, networks, private, addresses, 0 ] }
  node_templates:
    my_webapp_node_template:
      type: WebApplication
      requirements: # In this case the requirement definition has a range of [2,2] instead of [1,1]
        - redundant_database: db1
        - redundant_database: db2
    my_database_node_template:
    mariadb:
      type: tosca.nodes.DBMS.MariaDB
      properties:
        # omitted here for brevity
      requirements:
        - host:
            node: tosca.nodes.Compute
            node_filter:
              capabilities:
                - host:
                    properties:
                      - num_cpus: { in_range: [ 1, 4 ] }
                      - mem_size: { greater_or_equal: 512 MB }
                - os:
                    properties:
                      - architecture: { equal: x86_64 }
                      - type: { equal: linux }
                      - distribution: { equal: ubuntu }
                - mytypes.capabilities.compute.encryption:
                    properties:
                      - algorithm: { equal: aes }
                      - keylength: { valid_values: [ 128, 256 ] }
    mysql:
      type: tosca.nodes.DBMS.MySQL
      description: <string>
      [directives]: <list_of_directives>
      [properties]:
        root_password: { get_input: my_mysql_rootpw }
        port: { get_input: my_mysql_port }
      [attributes]: <map_of_attributes>
      [requirements]: #there are no requirement types - they match capability types / definitions
        - host: #example for when <node> is of type tosca.nodes.WebApplication
            node: tosca.nodes.WebServer
        - host:
            node_filter:
              capabilities:
                - host:
                    properties:
                      - num_cpus: { in_range: [1,4] }
                      - mem_size: {greater_or_equal: 512 MB }
        - database:
            node: my_database
            capability: Endpoint.Database
            relationship: my.types.CustomDbConnection
        - [description]: <string>
          capability: <symbolic_name_of_capability_definition> or <capability_type_name> # capability required for target node
          [node] : <name_of_node_type> # required when <name_of_capability_defintion> has been used for "capability"
          [relationship]: <string/map> #create named relationship when fulfilling the requirement (4.3.5.5.1.1 describes how to extend this when using map (type&interfaces))
          [node_filter]: #target node filtering
          [occurences]: <range_of_integer>
      [capabilities]:
        some_capability:
          [properties]:
            limit: 100
          [attributes]
          [occurences] #how many relationships can be made to this capability
      [interfaces]:
        Standard:
          operations:
            configure: scripts/my_own_configure.sh
      [artifacts]: <map_of_artifacts>
      [node_filter]:
        some_node_filter:
          [properties]: <list> #must match node template, node type, capability type etc
          [capabilities]: #actual filters
            - cap1
            - cap2: #this cap requires properties
                properties:
                  - cap_property_filter1
                  - ...
      [copy]: <source_node_template_name> #Source-node must not be a copy
      

  relationship_templates:
    my_connectsto_relationship:
      type: tosca.relationships.ConnectsTo
      interfaces:
        Configure:
          inputs:
            speed: { get_attribute: [ SOURCE, connect_speed ] }  
    storage_attachment:
      type: AttachesTo
      [description]: <string>
      [metadata]: <map_of_string>
      [properties]:
        location: /my_mount_point
      [attributes]: <map_of_attributes>
      [interfaces]: <map_of_interfaces>
      [copy]: <source_relationship_template_name> #Source-rl must not be a copy

  groups:
    #server2 and server3 are instances in node_templates
    server_group_1:
      type: tosca.groups.Root
      members: [ server2, server3 ]
  policies:
    - my_placement_policy:
        type: mycompany.mytypes.policy.placement
  workflows: <workflows>
  # Optional declaration that exports the Topology Template
  # as an implementation of a Node Type.
  substitution_mappings:
    <substitution_mappings>