# live-demo

## prerequisites

- some kind of hardware, either real or vms. (can this be done with vagrant? how many nodes can vagrant run a workstation?)
- one initial machine, like a linux laptop.
- cabling set up (power, network).

## vision

- have some policies in place, which are explained during the presentation
- test policies
- start some kind of software on initial machine
- boot all hardware pretty much at the same time
- make it possible to watch what happens and how fast everything is set up
  - PXE server running,
  - requests answered,
  - files sent,
  - names/ids created/used,
  - state(-updates) of all machines
  - state of (all) clusters
  - state of storage
  - application deployment
- when is everything done? When the exeutable exits? Or when a message is displayed?
- retrieve cluster access information, test application deployment
(- tear-down)

## concrete workflow

Assumptions:
- No preexisting active DHCP server
- No preexisting DNS server

<!-- - Explain tosca, show the hardware requirement of the demo-webserver, where it requires 4 gb (or more, depends on host capabilities) -->
- manually create some vms (configured to netboot) to simulate new hardware - two with 2GB ram, two with 4GB ram (or double that)
- run software on host, which runs PXE on the network and sends/simulates wake-on-lan to vms - MACs are inputted via tosca input params.
  - [x] requires semi-customizable PXE with DHCP (or proxyDHCP)
  - [x] requires wol function (and/or simulated wol for vms)
- By default a live-iso ("service-OS") is provided via PXE, which allows for hardware detection. This live-iso:
  - [x] ... runs a webserver, and writes/updates some information-files, which are available via said webserver
  - ... also contains a public key as authorized ssh-access; The couterpart private key lies only at the software.
    - [ ] requires ssh-key (-path) as input to software, which is checked for existance before generation is started
    - [WIP] requires generating/editing live-iso
      (add param for preexisting public-key-path)
  - software can now choose fitting host -> important to display the MAC of all possible nodes, and the selected node.
- select fitting host
  - [WIP] requires retrieving ip for service-os (arp or dhcp)
  - [ ] requires matching of requirements against existing hardware
- EITHER/OR
  - - ssh into all booted hosts and shut them down
    - for only the selected host: wol it again with new boot parameters;
    - this time the base OS will be installed. Further configuration can happen via SSH, or prescripted in the base OS.
  - - [ ] just shutdown all hosts except the selected host
    - [ ] ssh into it and configure it (live-os can install packages etc.)
- Configuration consists of:
  - [ ] Install and start a/the webserver
  - [ ] Provide index.html
  OR/AND
  - [ ] Install a DBMS on two nodes (for cluster)
  - [ ] Setup a ha-database
  - [ ] Install a webserver on one node
  - [ ] Install wordpress on webserver, use previous database
  OR Install a k8s cluster
- Enjoy the deployed website.


-> dhcp leases instead of arp
-> srv not recommended, directly check whether desired state is reached
-> check whether approach works for linux and windows, but only linux in demo
-> define requirement per system like netboot, wol and show statistics
