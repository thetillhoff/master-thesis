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
