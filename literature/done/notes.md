# notes on literature

## Development of an Ontology-Based Configuration Management System
new keywords:
  - configuration management
  - IT service management
  - rules
content:
  - standards, like ITIL [3]
  - purpose of configuration management [4]: 
    - identify configuration items
    - control changes to configuration items
    - build/provide spec to build product form cms
    - maintain integrity of baseline
    - provide accurate status & configuration for developers, end users & customers
  - configuration item types:
    - hardware:
      - server:
        - cpu
        - disk
        - ram
    - software:
      - system software
      - application software
    - document:
      - requirement
      - testing
    - it service management support:
      - change
      - incident
  - data format: XML/OWL
  - query language: SPARQL
potentially interesting sources:
  - Cabinet Office, “ITIL Continual Service Improvement”, The Stationery Office , Norwich ©2011 , 2011 Edition
  - Brenner, M., Gillmeister, M., "Designing CMDB data models with good utility and limited complexity ", Network Operations and Management Symposium (NOMS), 2014 IEEE.
  - Ontology Working Group , “Web Ontology Language (OWL)” , Available: http://www.w3.org/2001/sw/wiki/OWL
  - Stanford Center for Biomedical Informatics Research , “Why Protégé” , Available: http://protege.stanford.edu
  - Freitas, J., Correia,A., Brito eAbreu, F., “An Ontology for IT
  - Services”, In: the 13th Conference on Software Engineering and Databases, 2008
  - Kleiner ,F., Abecker ,A., Liu, A. , "Automatic Population and
  - Updating of a Semantic Wiki-based Configuration Management Database" ,"Informatik 2009:Im Focus das Leben", 28 November – 2 October 2009

## Containerized Docker Application Lifecycle with Microsoft Platform and Tools
has no sources!
definitions & definitions:
  - what is docker
  - state and data in docker applications
  - orchestrating microservices and multicontainer applications for high scalability and availability
  - language and framework choices
  - Build, CI, CD Deploy, Run Manage, Monitor
  - Introduction to orchestrators, schedulers, container clusters

## Das ultimative Handbuch für Windows Server 2016
cites:
- horizontal scaling of a cluster to scale [its] power is history.

## itgruene_pdf -> IT security criteria
- based on the "Orange Book" (original title: Trusted Computer System Evaluation Criteria) of US Department of Defense
- trust can be achieved with
  - trusting the manufacturer and its QA
  - testing the product oneself
  - have the product tested by an institution which one trusts
- three basic threats:
  - loss of confidentiality
  - loss of integrity
  - loss of availability
- basic security functions of "secure systems":
  - identification and authentication
  - administration of rights
  - verification of rights
  - audit
  - object reuse
  - error recovery
  - continuity of service
  -  data communication security.

## The Benefits of Agentless Architecture (ansible)
cites:
- You shouldn't have to manage your management system
- A management tool should not impose additional demands on one's environment - in fact, one should have to think about it as little as possible.
notes:
- communication via openssh or winrm (most critically reviewed programs in the entire world)
- push only what is required to nodes
- secure file copy
- "credential segregation": dev access dev machines, QA engs QA machines, admins production machines -> no risk for dev to accidently push to prod
- "zero bootstrapping": no managing the management; no software installation on remote hosts -> already part of OS' and no agent crash scenario
  -> Also no lock-in, since customers might not want to commit to one particular new system for management
- "push-based": immune to "thundering herd problem" -> no scaling issues
- kerberos-auth: clean

## An architecture for Modular Data Centers
- containers (physical) for modular data center
  -> might be a perfect example for a larger scale use-case
- "the only way that clusters of 10.000 to 100.000 nodes can be empoyed cost effectively is to automate all software administration. Those services that have adopted this highly distributed architecture, have automated most aspects of software administration including installation, upgrade, problem detection, and the majority of problem correction"

## Container Storage for dummies
- what is a container
- why software-defined-storage
- storage for containers (works with most SANs, NASs, SDSs...) vs storage in containers (containerized Gluster, Ceph, Rook)
- "Gartner predicts that by 2019 70% of existing storage array solutions will be available as a software-only version"
  "also predicts that by 2020 70-80% of unstructured data will be stored in less-expensive storage hardware managed by SDS systems"
- static vs dynamic provisioning in k8s
cites:
  - "Like the kid who says 'It's not that I don't like tomatoes; I just odn't like to eat them'"
- "Developers are certainly aware of the need for storage, but dealing with it can be a real pain. They simply want storage to be agile, reliable, and persistent. The mindset of developers regarding storage, or infrastructure in general, is that 'it's best when it's out of my way'. Cloud-native storage puts developers in control."
- container-native storage & management plane
- why container-native storage (ch4 and ch5)

## Fallacies of Distributed Computing Explained
- "The more things change the more they stay the same"
- 1994 Peter Deutsch drafted 7 assumptions and 1997 James Gosling added another one:
  - network is reliable
  - latency is zero
    Cool cite about why latency doesn't improve in contrast to bandwidth; speed of light limits to >30ms from US to EU and back
  - bandwidth is infinite
  - network is secure
    cool cite/source on how time consuming security reports are and how secure companies are and how trends show increasing number of attacks
  - topology doesn't change
    make it possible to add/remove nodes easily, don't hardcode endpoints
  - one admin
    -> who is reponsible for faults, SLAs etc
  - transport cost is zero
  - network is homogeneous
    -> XML & web services are popular because they help alleviate the affect of heterogeneity

## Ansible for devops
- explains vagrant, virtualbox and ansible setup

## Domain-specific language for infrastructure as code
- advantages of cloud:
  - pay for using and payload
  - nearly unlimited resources
  - simplified resource scalability
- problems with cloud:
  - heterogenity of cloud resources
  - lack of compatibility between cloud apis
  - lack of integration capacity between cloud environments
  - complex & errorprone cloud migration
  - too much human resources required for multiple infrastructure deployments
- two types of domain specific languages:
  - declarative with final state
  - imperative with operations
- two types of IaC:
  - push
  - pull
- standards
  - OCCI (Open Cloud Computing Interface) describes API for IaaS providers
    - developed by Open Grid Forum in 2010
    - RESTful
    - API for all types of management tasks like deployment, scaling & monitoring
  - TOSCA (Topology and Orchestration Specification for Cloud Applications)
    - developed by OASIS (Organization for Advancement of Structured Information Standards) in 2013
    - template language for defining required topology of a software application in the cloud
    - focus on defining application topology for the cloud orchestrator, not how the desired topology is deployed in the cloud
    - in "their" (OASIS') words: language which uses the service topology for describing service components and their interrelations.
  - common capabilities:
    - unified infrastructure model for different cloud providers
    - description of physical and virutal resources in clouds
    - high level of abstraction of infrastructure components
  - comparison:
    - TOSCA is higher level than OCCI
    - OCCI unifies cloud resources of different providers, while TOSCA describes them indepenetent of cloud computing concepts -> TOSCA advantage
    - OCCI uses sequential order of resource creation, which makes debugging easier
    - OCCI & TOSCA are both incomplete & ambiguous at interpreting model parameters and objects
    - last update: TOSCA: 2019 (time of writing), OCCI: 2016
- clients requires lots of updates, since cloud providers update their apis all the time
  - incomplete openTOSCA
  - PoC TOSCA2Chef uses openStack and Chef (not open source)
TOSCA:
  - uses YAML files with nested templates
  - basic elements:
    - datatype
    - artifact
    - capability; f.e. can run containers, ip-port-endpoint, protocol
    - attribute; allowed-input-ip-address, allowed-remote-ip-address (range/single), networks, public-address
    - property; num_cpus, os type, os version, architecture
    - relationship
    - interface
    - node
    - group
    - policy
  - property vs attribute; property is the desired value, attribute is the actual value
  - ansible can interpret TOSCA
  - "normative definitions":
    - blockstorage is attached to compute
    - network port is bound to compute
    - network port links to network
    - compute connects to compute
    - software component is hosted on compute
    - webserver, dbms, container runtime are all software components
    - webapp is hosted on webserver
    - database is hosted on dbms
    - container app is hosted on container runtime
  - -> is incomplete, f.e. no volume creation
  - generic core model and cloud provider specific model
  - have a mapping between generic format and cloud-provider-specific format - which can be customized
- Have generic types, with type-specific custom attributes (like template-extension) and instance-specific custom attributes
  - load-balancer
  - SecurityGroup
  - SecurityGroupRule
  #- PrivateSubnet
  - NetworkInterface
  - Image (OS)
  - Volume
  - IP
  - Network
  - Port
  - Server/Compute
    - KeyPair
    - Flavor
- result:
  - impossible to carry out complete mapping between cloud providers, since even basic parameters of tosca standard aren't implemented
  - tosca standard needs development
  - tosca allows for quickly adding updates
  - tosca wants to add support for software configuration

## A Comparative Study of Baremetal Provisioning Frameworks
- register hardware at provisioner either manually via MAC or autodiscover
- when request for node allocation comes, the provisioner picks appropriate set of nodes of available list
- Use IPMI as power driver
- DHCP/PXE with TFTP download of a bootstrap image (f.e. ipxe)
- download actual OS image, restart after installation to boot new OS
- in most cases a configuration management system (CMS) is used along with provisioning, as installation of software is often required, too
- frameworks:
  - emulab; network topology "simulator"
  - openStack Ironic
  - Dell Crowbar; also for OpenStack deployment
  - Cobbler; focus on RedHat related linux-os, has its own CMS & supports Chef
  - Canonical MaaS; focus on Ubuntu with Juju as CMS
  - Razor; by Puppet labs, complement to puppet
- evaluation of:
  - vendor lock-in:
    - all frameworks are open source
    - Razor & crowbar have an enterprise version with more features and support
  - maturity
    - emulab is older, but focus on academic
    - cobbler is older
    - maas, crowbar, razor and ironic are (almost equally) relatively new
  - number of deployments/users
    - emulab developed by Flux research group at uah university
    - crowbar developed by dell, rackspace, suse, some telecom and finance companies
    - razor developed by emc and puppet labs
    - ironic developed by HP cloud and diverse community
    - maas developed by canonical
  - difficulty of installation
    - emulab is complex af, need FreeBSD in version 8.3, needs to be built from source
    - crowbar has built-in iso for provisioner, so not complex
    - maas is simple, because only ubuntu is considered
    - razor is simple, because its deployed as a puppet module
    - cobbler seems simple
    - ironic is complex, because it needs a lot of other services like nova, glance, keystone
  - stability
    - mostly unknown (only breaking changes are considered, and for that almost no data was found)
  - difficulty of maintenance
    - all have a web-ui
    - ironic might not be able to autodiscover hosts, current state?
  - hardware requirements
    - IPMI
    - PXE
    - VLAN after allocation?
    - crowbar required BMC (iDRAC) -> extended IPMI
  - breadth of features
    - VLAN
      - emulab supports it, manual and automated
      - Openstack Ironic does support it, but requires manual configuration as "neutron" (network component) is not fully integrated
      - others have VLAN support integrated in their CMS counterpart
    - network topology
      - only emulab allows configuration
    - os support
      - maas only ubuntu
      - ironic & cobbler only linux, no windows
      - crowbar, emulab, razor support linux and windows
    - high availability (allocate new machine when needed)
      - on roadmap for maas
      - manually in crowbar
      - planned for ironic
      - emulab no automatic failover but can preempt node(s)
    - performance (time for provisioning)
      - emulab and ironic: 5-6min
      - parallel deployment, but network is bottleneck
  - scheduling
    - f.e. allocate node with GPU

## A review of existing cloud automation tools
- "virtualization has reduced teh time required to deploy computing resources from weeks to few minutes"
- cloud promises access to flexible and elastic compute resources at minimal cost
- Managing the growing infrastructure is one of the major challenges
- [1] system to create, configure and manage the CM deployments in the cloud
  Juve, Deelman, Automating Applicatoin Deployment in Infrastructure Clouds
- [2] automatic deployment mechanism with openstack
  Zhang, Shang, An Automatic Deployment Mechanism on Cloud Computing Platform
- [3] migration framework for infrastructure in the cloud
  Callanan, O'Shea, Automated Environment Migration to the Cloud
- [4] discussion of modules & their integration for mgmt & automation of cloud-based infrastructure
  Wibowo, Cloud Management and Automation
- terraform
  - execution plans; how to reach desired state
  - resource graph; parallelize wherever possible
- cloudformation
  - aws specific
  - resource graph with automatic dependency detection

## Bare-Metal Marketplace At The Bottom Of The Cloud
- OpenCloudExchange (OCX) where multiple organizations freely cooperate & compete with each other for offering different hardware resources while customers can choose from numberous competing services instead of a single provider.
- goals:
  - bare-metal allocation and isolation service
  - diskless rapid provisioning service
  - security model
  - market based incentive system
- base questions:
  - move hardware between clusters on-demand
  - cluster set up fast enough to respond to rapid fluctuations in demand
  - single system for a wide variety of scenarios, like multiple clusters of a single company to different tenants in the same colocation facility to new model of cloud with multiple providers
  - design system that provides cluster owners incentives to offer their hardware resources to other clusters
- provisioning systems:
  - ironic
  - maas
  - emulab
  - geni
  - foreman
  - xCat
  - [25,8,28,4,11]
- "Clusters are typically stood up with sufficient capacity to deal with peak demand; resulting in silos of under-utilized hardware."
- "A tenant of the ESI [Elastic Secure Infrastructure] has to trust only a minimal functionality of the tenant that offers the hardware resources. Rest of the services can be deployed by each tenant themselves."
- [7,17] custom deployment practices of organizations with on-prem requirements
- [5] OCX Open Cloud eXchange model
- note to self: OCX allows for easier 3-2-1 rule of data backups
- "The goal of a cloud is to maximize its profits"
- important bits:
  - isolation & bare-metal access
  - fast provisioning / migration of hosts
  - security model; mechanism to verify whether a server matches the security standard for a cluster
  - incentive model; offering unused resources has to pay off somehow
- design principles & architecture
  - give control to tenants as much as possible
  - minimize shared services for security purposes and to enable new capabilities
  - partition components into micro-services for maintenance and tenant support
     - better scaling
     - each service has its own development life-cycle
     - each service has a well defined API interface
     - loose coupling
  - incentives instead of top-down decision making
- services
  - isolation service; allocates server to cluster, deallocates servers when unused
    - Exokernel-like approach [19]
    - partitions physical hardware & connectivity
    - enables direct access to physical resources
  - attestation service; checks integrity of server
  - provisioning service; sets up os and apps
    - serves os images with applications from remote-mounted boot-drives
    - uses ceph as underlying storage [21,22]
  - security model
    - tenants can control trade-offs between security, price and performance
    - self-hosting provisioning possible & therefore increased security
  - incentive model
    - tenant earn revenue for offering resources
    - change of demand & supply is reflected by dynamic changes in the price of the resources
    - auctions decide best placement of resources among competing demands
- hardware isolation layer (HIL)
  - exokernel-like approach [10]
  - network isolation, so that multiple provisioning services can work in parallel for a different set of nodes
  - "drivers" for different Oout-of-band-modules (OBM) used for power-cycling & switch specific drivers for network isolation
  - fundamental operations
    - allocation of physical nodes
    - allocation of networks
    - connecting nodes and networks
  - 3000 loc that have to be trusted
  - "ipmitool" checking whether all consoles are disconnected
- bare-metal provisioning service (BMI)
  - non-local boot-drives, and therefore no "installation" time; only boot
- security model
  - threat phases
    - prior to occupancy; bad firmware threatens integrity, for example from prior tenant, or node provider
    - during occupancy; communication between servers, communication between server and storage; denial of service
    - after occupancy; storage and memory contents
  - attestation service -> measuring firmware & software (compare "all values" plus signature)
    - checking all values helps also against firmware with bugs [16,15,27,6,22,12]
    - time incease between un-attested and full-attested (LUKS & IPSec) is from a total of 300s to 450s
    - initrd, then network change, then kexec
- uses
  - LinuxBoot vs UEFI (Power-On-Self-Test on UEFI takes longer, but apart from that there is not much idfference)
  - pxe
  - ipxe
  - keylime
  - IPsec between node and boot-os-storage
  - LUKS for disk encryption
- foreman takes 11min for provisioning, 5min of that are os installation
  - no security procedures and therefore faster than most cloud provisioning systems
- incentive system
  - centralized system more efficient [9,26,23,14,3,18]
  - centralized may not suit a multi-provider cloud
  - transparent "matching" of resources requests with resource availabilities
  - minimum is same cost as before (when it was unused)
  - orgs can set their minimum price
- open questions
  - marketplace feasible? proposal: PoC
  - clusters are optimizied towards a specific goal; can this still work in a marketplace? proposal: probably yes
  - get average "idle" resources in current implementation, how much more load can be taken on and is it generally feasable
- allocation and isolation service: [13]
- provisioning service [24,19]
- security framework [21,20]
- incentive model OSDI 2020, spring 2020
sources:
- [3] Anderson, D. P. Boinc: A system for public-resource computing and storage., 2004
- [4] Berman, Chase: A federated testbed for innovative network experiments., 2014
- [5] Bestavros,Krieger,: Toward an open cloud marketplace: Vision and first steps, 2014
- [6] Bulygin,: Summary of attacks against BIOS and secure boot., 2014
- [7] Butler, "Which is cheaper: Public or private clouds", 2016
- [8] Canonical MaaS
- [9] Duplyakin &  Johnson, "The part-time cloud: Enabling balanced elasticity between diverse computing environments", 2017
- [10] Engler, D. R., Kaashoek, M. F., and Otoole, J. Exokernel: an operating system architecture for application-level resource management., 1995
- [11] foreman
- [12] heasman: Rootkit threats, 2006
- [13] Hennessey,: Designing an exokernel for the data center., 2016
- [14]Hindman,: Mesos: A platform for fine-grained resource sharing in the data center
- [15]Hudson,: ThunderStrike 2: Sith Strike., 2015
- [16]Hudson,: Thunderstrike: EFI firmware bootkits for Apple Macbooks., 2015
- [17] Kirkwood & Suarez, "Cloud Wars! Public vs Private Cloud Economics", 2017
- [18] Lai,: Tycoon: A distributed market-based resource allocation system: 2004
- [19] Mohan & Turk, "M2: Malleable Metal as a Service", 2018
- [20] Mosayyebzadeh & Mohan, "Supporting security sensitive tenants in a bare-metal cloud", 2019
- [21] Mosayyebzadeh & Mohan, "A secure cloud with minimal provider trust", 2018
- [22] Rutkowska,: Intel x86 considered harmful, 2015
- [23] Schwarzkopf, "Omega ...", 2013
- [24] Turk & Gudimetla, "An experiment on bare-metal bigdata provisioning", 2016
- [25] openstack ironic wiki
- [26] Verma, "Large-scale cluster management at google with borg", 2015
- [27] Wagner: BIOS-rootkit LightEater
- [28] White: An Integrated Experimental Environment for Distributed Systems and Networks, 2002

## TOSCA wikipedia
- approved by OASIS since 16.01.2014
- related:
  - AWS Cloudformation (JSON)
  - OpenStack Heat (adopted TOSCA for standardized templating)
  - Cloudify (TOSCA-based) (incorporates Alien4Cloud, which is an TOSCA-design-tool)
  - Ubicity (tooling & orchestrators based on TOSCA)
  - MiCADOscale (TOSCA-based resource orchestration framework for containerized apps)
  - SeaClouds (multi-cloud management of service-based apps, supports TOSCA) (EU FP7 funded)
  - DICE (modeldriven devops; TOSCA as pivot language between modelling & ops) (EU H2020 funded)
  - Project COLA (pluggable framework for optimal & secure deployment & orchestration of cloud applications - uses MiCADOscale) (EU H2020 funded)

## OASIS wikipedia
standards:
- AMQP: Advanced message queuing protocol
- CAMP: Cloud application management for platforms
- CAP: Common alerting protocol
- DocBook: markup language for technical documentation
- EML: Election Markup Language
- MQTT: Message Queuing telemetry transport
- OpenDocument: document format
- PKCS #11: standard that defines a platofrm independent api to cryptographic tokens such as hardware security modules and smart cards (api is named cryptoki)
- SAML: Security Assertion Markup Language
- SARIF: Static Analysis Results Interchange Format
- UBL: Universal Business Language; standard for electronic documents like invoice in xml
- VirtIO: standard for paravirtualized devices
- WSDM: Seb Services Distributed Management
members:
- Dell, GM, IBM, ISO/IEC, Microsoft, Oracle, RedHat,universities, government agencies, ...
competition:
- w3c
addtional notes:
- https://en.wikipedia.org/wiki/Service_choreography#Web_Service_Choreography
  BPMN: https://en.wikipedia.org/wiki/Business_Process_Modeling_Notation
  WS-CDL: https://en.wikipedia.org/wiki/WS-CDL
- https://github.com/PrivateSky/privatesky
- https://en.wikipedia.org/wiki/YAWL

## https://www.cloudcomputing-insider.de/was-ist-oasis-tosca-a-883140/ 2019
- initial specification by Cisco, Citrix, IBM, EMC, SAP, RedHAt, Capgemini, Software AG
- XML & YAML are supported
- highest level: topology template, node template, relationship template and management plan
- structure of app is shown as graph in topology template: nodes and relationships are nodes and edges
- management plans are workflows, f.e. to start or end a service
- in total: cloud service is fully described by service template (which contains all the other templates and plans)
- "supports TOSCA" means service template can be read & parsed
- service definition is fully machine readable, management and orchestration are well defined
- indepentent from cloud-providers
- apps can be migrationed between environments & automated provisioning or management is possible
- supporting TOSCA as cloud-provider allows for easier migration TO your cloud, not only from it
- unique advantage of your platform if you support TOSCA

## https://www.admin-magazin.de/Das-Heft/2018/02/TOSCA-Standard-fuer-die-Cloud 2018 - BOUGHT
goals of tosca:
- automation of deployment and management
- portability of app(-description)
- interoperability & reusablity of components
notes:
- XML-based metamodell for formal description of app-structure -> servicetemplate
- servicetemplate contains topologytemplate, nodetype, relationshiptype and plans
- topology is a graph
- app-components and their relationships equivalents are Node templates and Relationship templates
- Node type defines the strucutre of an node template
  - more specific: node type defines the structure of the "watchable" properties of the app-component, its management functions, the potential states, the system requirements and the abilities it publishes
  - properties are described with property definitions, requirement definitions and capacity definitions
  - example contents for node type for a web-based crm: licence key, admin user&password, system requirements like versions of apache, mysql & php
- relationship type has validSource and ValidTarget; it allows to describe action between valid interfaces
- artefacts are a content-element or the information for executing a deployment or management operation of a app-component
  - scripts, binaries, images, config files, libs etc are all artefacts
  - metadata for those are contained in ArtifactType, which will then be put together in Artifact template
- tosca requests to use plans, but doesn't care about the language for descripting them
- plantype, with two predfined ones
  - buildplan; describes the creation of service tmeplates
  - terminationplan; ends a service instance
- CSAR (cloud service archive) is a archive format for storing all tosca files
  - actually a zip archive
  - two folders have to exist in there:
    - "Defintions": tosca defintions in files with the extension "tosca". One of those contains the service template, which acts as the entry point
    - "TOSCA-Metadata": contains a file "TOSCA.meta" which holds metainformation about the CSAR file itself and other components of the environment
  - The CSAR file could be processed on a third-party-platform. This requires an "TOSCA-container". Currently this is only implemented in OpenTOSCA and OpenStack
advantages:
- modelling of components, of processes and of dependencies
- in order to use tosca, someone had to understand the app at some point at least
- XML (&YAML) are more or less easy to use
- non-functional stuff can also be described with aid of special policies
- independent of cloud provider
disadvantages:
- doesn't reuse current specifications for describing infrastructure; those would make service description faster and easier
- since 2013 there are almost no implementations; user friendly tools are missing
- unknown collaboration with other standards, like CAMP or modelling technology like WS-BPEL

## https://www.admin-magazin.de/Das-Heft/2018/02/Apache-ARIA-TOSCA 2018 - BOUGHT
- two parts of tosca: topology to describe the relationship between apps, and orchestration, especially model-management of the app
- "Apache-Incubator-Project" ARIA TOSCA: lib, cli tool
  - orchestration-tool
  - vendor & technology independend implementation
  - cli for template development
  - lib (sdk) for development of tosca-able software
  - TOSCA DSL Parser; validates template & creates deployment-diagram & retrieves relationships
  - declarative YAML
  - technologyspecific types can be created & used via ARIA-plug-ins -> no need to change parser code
  - supported: TOSC Simple profile 1.0 & TOSCA Simple Profile for NFV (Network Functions Virtualization)
  - ARIA workflows: dynamic interactions with app template (define tasks and their ordner - tasks can be implemented as plug-ins)
    Can be part of tosca template & therefore access the graph easily - at runtime, with dedicated API for runtime context
    examples: installation, deinstallation, adjusting installation, repair installation
  - implemented in python
  - it is possible to combine different tech of single provider, but also tech of different providers
  - existing plug-ins: 
    - IaaS: OpenStack, VMware, AWS, GCP, Azure
    - CM: Puppet, Chef, Ansible, Saltstack
    - Container: Kubernetes, Docker, Mesos, Swarm
    - SND: ODL, ONOS
    - Skript: Bash, Python, ...
  - plug-ins are part of template, and can be loaded dynamically
  - aria tosca can work with cloudify - the most advanced tosca implementation
  - plugins in WGN-format
example-process:
- store/load service-template
- create service (based on template)
- run installation workflow
cloudify:
- blueprint-composer: drag-and-drop service editor
- again python
openTOSCA: alternative to cloudify:
- german software (uni stuttgart)
- OpenTOSCA Container, a tosca runtime
- Winery, a graphical modelling tool
- vinothek, a portal for apps
conclusion:
- not user-friendly
- webbased admin-interface & tools for modelling & tosca-editors are missing
- could be cool, but currently is not so much

## https://docs.vmware.com/en/VMware-Telco-Cloud-Automation/1.9/com.vmware.tca.userguide/GUID-43644485-9AAE-410E-89D2-3C4A56228794.html 2021
- tosca on vmware
- doc is incomplete af

## https://www.oasis-open.org/committees/tosca/charter.php
- should be compatible with BPMN 2.0 and WS-BPEL
- interfaces should be expressed in a proper REST-style based on HTTP and specified via WSDL 1.1 & allows for the use of scripts


# https://ahmet.im/blog/cloud-instance-provisioning/
- aws uses cloud-init
- digitalocean uses cloud-init
- gcp uses selfimplemented guest agents written in python (running as systemd services)
- azure uses cloud-init (and waagent (azure linux guest agent, https://github.com/Azure/WALinuxAgent))


# Comparative Study of DSL Tools
dimensions [DSL Implementation in MetaOCaml, Template Haskell, and C++ by Czarnecki]:
- approach: what is the primary approach supported by the DSL tool (f.e. translation/templating, term rewriting)
- guarantees: what guarantees are provided by the DSL tool in terms of syntactic and semantic well-formedness of the transformed-to constructs (f.e. well-typed, syntax valid at run-time, none)
- reuse: can the 'user-defined' aspects of the DSL implementation be reused
- context-sensitive transofmration: can the DSL tool perform context-sensitive transformation
- error reporting: can the DSL tool report errors in terms of the DSL source (line number and column offset) (run-time/compile-time)
metric:
- lines of code: given a case study, how many lines of code are required to represent the domain-specific information
- aspects to learn: given a case study, how many aspects need to be learned to implement a DSL
potential sources:
- Domain-Specific Languages: An Annotated Bibliography by Deursen
- Modular Domain Specific Languages and Tools by Hudak


# Comparing General-Purpose and Domain-Specific Languages: An Empirical Study
dimensions [Usability analysis of visual programming environments: a "cognitive dimensions" framework by Green and Petre, Ten years of cognitive dimensions in visual languages and computing: Guest editor’s introduction to special issue by Blackwell]:
- closeness of mapping: languages should be task-specific
- viscosity: revisions should be painless
- hidden dependencies: the consequences of changes should be clear
- hard mental operations: no enigmatic is allowed
- imposed guess-ahead: no premature commitment
- secondary notation: allow to encompass additional infomration
- visibility: search trails should be short
- consistency: user expectations should not be broken
- diffuseness: language should not be too verbose
- error-proneness: notation should catch mistakes avoiding erros
- progressive evaluation: get immediate feedback
- role expressiveness: see the relations among components clearly
- abstraction gradient: languages should allow different abstaction levels


# Usability Evaluation of Domain-Specific Languages
- instances of DSL == instance model == sentences of DSL


# Comparison between internal and external DSLs via RubyTL and Gra2MoL
comparison measures:
- three elements of DSL: abstract syntax, concrete syntax and semantics (executability and optimizations)
- quality criteria: extensibility and efficiency
- DSL tooling: tools for developing DSL and tools for using DSL
->
- concrete syntax: does the DSL require a specialized syntax? Is the host language syntax suitable for the DSL? How much effort is needed to embed the DSL in comparison to building the DSL from scratch?
- abstract syntax: In which cases mgiht an abstract syntax be necessary, and in which is it possible to manage without it? How different is it to support an abstract syntax in each case? This last issue is related to the following aspect
- executability: How much does the host language assist inthe executability of the DSL? Do we need to adapt the (internal) DSL to facilitate its executability? In which cases is it most recommended to create an interpreted/compiled language?
- optimizations: Can the execution process be optimized to improve the efficiency?
- language extension: How difficult is it to incorporate new constructs into the language?
- integration and library availability: How can an internal/external language facilitate integration with other tools such as editors? Are there libraries required available in the choses host language?
- DSL development tools: Are there tools that facilitate the creation of internal/external DSLs? How much freedom do they offer in the creation of the language? Do these tools upport the aspects identified in this comparison?
- target audience and usability: Does the target audience expect a language with a special syntax? Are they already used to the host language syntax?

# The Mythical Man-Month, 1985, Frederic Brooks
"There is no Silver Bullet in Software Engineering"

# Understanding OASIS Tosca [https://www.youtube.com/watch?v=C75LBxsQNsc]
design tools:
- vnomic service designer
- ibm owrkload deployer
- zenoss cloud monitoring
service marketplaces:
- ibm cloud marketplace
- sap marketplace
cloud managers:
- fujitu flexframe orchestrator
- hp cloud management and automation
- huawei telco cloud solution
- ibm cloud orchestrator
- vnomic supported clouds

architects: model services, policies, requirements
developers: develop, unit test scripts, plans, artifacts
qa-teams: build and test releases, updates and configurations
operations: deploy, manage and monitor application lifecycle

application requirements are modelled independently form cloud infrastructure capabilities

tosca orchestrator automatically matches tosca to cloud-provider-specific stuff and optimizes

forrester names tosca as top four cloud open standard (mar 2014)

multi-cloud interoperability demonstrated at:
- eurocloud 2013 with ibm, sap, fujitsu, huawei, hp, vnomic, zenoss, ...
- open data center alliance (jan 2014)

open source projects:
- openstack
- eclipse
- cloudify
- apache
- celar
- canonical Juju is an tosca orchestrator which supports aws, openstack, azure, hp cloud, joyent and bare metal - NOT TRUE (any more)

# Introduction to TOSCA - Yaron Parasol [https://www.youtube.com/watch?v=EgzwaQi03l8]
## impact of human error
- 80% of outages impacting mission-critical services willb e caused by people and process issues
- 50% of those issues will be caused by change/configuraiton/release integration and hand-off issues
## the path to orchestration
- standardize
  - lower cost
  - increase performance
  - reduce complexity
- consolidate
  - release assetes
  - improve efficiencies
  - improve management and control
- virtualize
  - lower cost
  - increate utilization
  - higher flexiblity
- automate
  - lower cost
  - improve user experience
  - speed to market
- orchestrate

rackspace is also a suppporter
v1 of tosca was xml
v2 yaml
tosca:
- can describe any topology
- any automation process
tosca building blocks:
- application topologies
- workflows
- policies
bpmn for workflow descriptions

Every components is a node

An interface is a set of hooks ("operations")

tosca.interfaces.node.lifecycle: create, configure, start, stop delete operations
capabilities == what it can provider to other nodes
requirements == what capability is required

basic relationship types are:
- dependsOn with subtypes:
  - hostedOn (can be displayed as containment)
  - connectsTo (linked connection)

node template:
- is an instance of type (like object to class)
- has specific properties
- has artifacts (what to install, how to install (mapped to interface hooks))
- has requirements and capabilities (for relationships)
example from node template:
  requirements:
    - database: some_db
      interfaces:
        tosca.interfaces.relationships.Configure:
          pre_configure_source: scripts/some_db_configure.sh

workflows: requires workflow engine
policies:
- brings monitoring as input -> ongoing evaluation of rules
- can invoke more processes

Cloudify contains gui, workflow engine, blueprint and runtime data, task manager and an agent.
Spawned vms have an agent that gives feedback to task manager and manages the software on that vm. Task manager outputs logs & events
An optional additional monitoring agent is on each vm, which reports to policy engine, which reports to metrics-vm with monitoring data.

# https://smartbear.com (creator of openapi/swagger)
"Quality isn't just a goal. It's the whole point."
With rapid transformation comes risk.
Make sure your software works when it's needed most.


