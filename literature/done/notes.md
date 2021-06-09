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
    - in "their" (OASIS') words: loanguage which uses the service topology for describing service components and their interrelations.
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

