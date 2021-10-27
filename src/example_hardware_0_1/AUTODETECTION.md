# Autodetection of new hardware

This document describes how autodetection of new hardware works in hardware-tosca.

## Reference from tosca simple profile

Tosca simple profile defines a type tosca.nodes.Compute which can logically be viewed as a real or virtual server.
Thus, "\*Compute" (or "Compute.\*") refers to some subtypes of physical or virtual servers.

## Unique identifiation of servers

In order to never mix servers up, redeploy the same server etc a way to uniquely identify a server is needed.
Since not every server has a serial number, the second-most intuitive way are the MAC addresses.
In many datacenter environments there is one dedicated management network, with exactly one connection to each server.
If the MAC of that connection is used to identify servers, it is a simple way to uniqely do so.
> ### Sidenode
> IPMI-tools like the one from SuperMicro address hosts via hostname/IP; Direct connections to MACs is not supported.
> IPMI uses RMCP for connections, which in turn uses UDP port 623.

## Bare servers

Servers can be physical or virtual.
NodeName for physical servers: physicalServer / physicalCompute
NodeName for virtual server: virtualServer / virtualCompute

## Servers with operating system

There are three ways to deploy servers (both physical and virtual):
- attach an live- or autoinstall-iso via IPMI
- set server to boot via network and deploy live- or autoinstall-iso via PXE
- manually install an operating system. This includes deployment via potential other mechanisms. Big question here is how to detect them?

NodeName for servers deployed via IPMI-iso-mount: ipmiCompute
NodeName for servers deployed via PXE: pxeCompute
NodeName for manually deployed servers: unmanagedCompute

## Requirements
ipmiCompute:
  Since IPMI is by default only supported in some physical servers, there is a subtype called simulatedIpmiCompute. It contains the same actions as ipmiCompute, but the actions don't call ipmi-tools but other means like hypervisor commands.
  Requires a live- or autoinstall-iso.

pxeCompute:
  Since PXE is supported by most physical and virtual servers, no subtypes are required.
  Requires a configurable DHCP server.
  Requires a configurable TFTP server (HTTP server might be possible as well).
  Requires a live- or autoinstall-iso.

unmanagedCompute:
  This type is considered neither physical or virtual and has no requirements.

## Workflow with ipmiCompute

Normally, the ipmi interface of a server is always on (as long as the server has power). In many datacenter environments, all ipmi interfaces are on a dedicated management network to ensure no conflicts between management and production traffic, and to be able to have dhcp running for ipmi, but static ips for production servers. There are other ways, but this is the cleanest (best-practice) way to mix management dhcp and production workloads.
Anyway - this means, that in order to configure the iso of an ipmiCompute node, the ipmi interface has to be connected to a network with active dhcp.
By pinging the broadcast-address (f.e. 192.168.0.255) and checking the arp table (f.e. with arp -a), it is possible to map MAC- to IP-addresses of those interfaces.


## Workflow with simulatedIpmiCompute

When a server doesn't have a physical ipmi interface, there is an additional requirement to be still able to automate the same actions: When the node is virtual the hypervisor can do the same things but with different commands.

## Issue with IPMI

IPMI only allows for autodetection when all IPMI interfaces are on a dedicated management network with active DHCP. Only then it is possible to check whether a new node joined - namely when a previously unkown IPMI interface wants an ip-address.

# Issue with IPMI and PXE

Both IPMI (as described above) and PXE require an active DHCP server, which somehow runs actions whenever new participants are detected. Or at least the DHCP server has to provide a maintained list of participants.
This means either the lease-list of the active DHCP server must be accessible or the DHCP server must support running actions on new leases or the software listens to dhcp requests in addition to a dhcp server - this is possible since they are broadcasting their request.
The latter seems unfeasable, since the software isn't intended to run all the time.
The second is out of reach, since either DHCP servers support this or they don't - AFAIK they don't.
So the first one is the most intuitive: Access the lease-list of a DHCP server which is running in a dedicated management network where only ipmi-interfaces "live".
There is one functionality where DHCP servers DO "run an action" when a new machine comes up: PXE-booting. For this scenario, it is required to configure an existing /long-living DHCP server and an TFTP (or HTTP) server for the boot-images.

# The good thing about requiring DHCP as central component

While DHCP might be single-point-of-failure, it is relatively easy to activate a new one, or deploy another DHCP server.
PXE-scenario: Since hardware-nodes are mostly not rebooting all the time, it is possible to run them with live-isos. This leaves all of their disks open for actual storage. Since PXE requires an TFTP (or HTTP) server which provides the boot images, the software could monitor its load and issue hardware-starts accordingly.
IPMI-scenario: When machines have IPMI-support it makes sense to use it, too (as it costs extra). 

# How can autodetection and automatic deployment be reflected in TOSCA?

- Add state. If there is a maintained state, it is easy to detect which nodes are new (since last run).
- Require components like an configurable DHCP-server (with PXE), and an always-running application like a TFTP- (or HTTP-) server.
  At minimum the latter can be achieved with k8s.
  Anyway, to have the whole cluster able to adapt to scaling during unmaintained times (orchestrator software not running), the desired state has to be stored somewhere.
  k8s does it the exact same way - it stores the desired state in etcd.
  cloudify and all other TOSCA-orchestrators store the desired state themselves which requires them to be up and running at all times.
  This means, TOSCA-orchestrators are currently by design the backplane of the infrastructure (often described in TOSCA as well).
  But the orchestrators are (also by design) not meant to instantiate new components, but only request a new instance by another backplane (for example aws,openstack etc).
  So we require two backplanes, first the provider, then the TOSCA-orchestrator to be up and running at all times.
  For bare-metal there are only two providers (which are supported by opentosca, cloudify and non-tosca terraform): openstack and vmware.
  VMware is quite costly and closed source, which makes it hard to extend. Yet, it might be fine for many environments.
  Openstack might be fine for some environments, but it is extremely complicated (more than a dozen components, all deployed seperately, all configured seperately, ...)
  But what is really required of the provider backplane?
  - detect new (powered off) nodes and make them accessible (i.e. via ssh) -> power on, (install &) run os with predefined ssh-settings
  - delete (poweroff? proper deprovision with deleting disk contents?) nodes
  - instantiate more complex objects like vlans, firewalls, (virtual-/shared-/...)disks, load-balancers, managed databases
  The first two actions are relatively easy to accomplish. The later seems hard at first, but isn't this the perfect use case for TOSCA as well?
  Public cloud providers instantiate those with software-defined everything. Given the fitting hardware (supporting completely software-defining them), it must be possible to create a provider backend easily extensible via TOSCA. And this TOSCA-described provider backend requires only
  - detect new nodes and make them accessible
  - delete (poweroff?) nodes
  - run actions to deploy and configure software on nodes
  And now the later seems easy as well. Not only easy, but also doesn't require an always-running backplane. Instead it can be run once, where it might take longer, since it is deploying the tosca-defined provider backend. This backend must then be able to store "the state" - which means other tosca-templates, and where/how they are deployed.

# Does this have advantages over vmware, openstack?

- easily extensible
- simpler than openstack
- free

# How will autodetection and automatic deployment be implemented in this MA?

Since time is limited, only one approach will be taken. The optimal case would be to select the approach which requires less work.
Do be able to decide that, a comparison of required work follows.

For both PXE and IPMI there is an active DHCP required.
For PXE it needs to monitor the MAC addresses of "available" devices, so it can power-on on-demand.
For IPMI it needs to monitor "availabe" ipmi interfaces, so it can power-on on-demand.

## PXE

PXE also works for almost all machines.

- Autodetect new nodes
  - Implement DHCPnode in TOSCA, where the PXE is fully configurable
    MAC addresses are provided at input, then detection/provisioning takes place. New hosts are not in arp-table.
- Automatic deployment of nodes
  - Implement TFTP- (or HTTP-) node, where the boot-images can be placed.

## IPMI interface on dedicated network

IPMI works only on the minority of computers. Also, for hypervisors there is no "proper" solution (except for kvm there is a basic implementation).

- Autodetect new nodes
  - Both of the below could be merged into the same thing.
  - Implement DHCPnode in TOSCA, where the ipmi-interfaces are autodetected.
  - Implement "ipmi interface list in dhcp"
    It needs to monitor available ipmi interfaces, so it can provision the server when it is needed (and power it on)
- Automatic deployment of nodes
  - implement ipmi-calls - or simulate them, so it can be used with hypervisor

## Conclusion

The PXE variant is less complex and requires less work. Hence, the PXE-approach will be selected for the live-demo.

## Implementation details

- Describe a DHCPnode in TOSCA, where subnet and PXE is configurable
  - This node doesn't need to be actually deployed, but could have interfaces like configure, where an preexisting DHCP server is configured.
  - Or, the DHCP is included in `eat`, and a list of MAC addresses has to be handed in as inputs.
    These MACs are then waked-on-lan and request their netboot stuff from the DHCP
  - Assume no preexisting stuff. Which means we have the DHCP included in eat.
    When the vms are created, they are configured to netboot. In a real world environment this could be done via IPMI.
    Since the live-demo will use vms instead of actual machines, wake-on-lan won't work, so wake-on-lan has to be script-simulated (alternatively, there is a powershell script to accept wake-on-lan for hyperv https://deploymentpros.wordpress.com/2016/11/28/wake-on-lan-for-hyper-v-guests/, nothing found for kvm though).
- Describe a TFTP- or HTTP-server for the actual netboot stuff.
  This server also needs to be included in `eat`.
  Research whether HTTP-only netboot is possible. Should be, with ipxe, which runs on bios and efi.
  boot iso can be placed here as artifacts, or their creation script can be placed here. <- latter is cool show for live-demo, but prepare the custom iso just in case.
- Describe an application like webserver and/or database to be deployed. This includes the underlying OS.
---
Yet, to reduce workload during this MA, an external DHCP server will be run. In this case, dnsmasq in a docker container.
The reason for this is, that the creator (me) already has experience with PXE-booting with dnsmasq, and dnsmasq supports running in proxyDHCP mode, which helps developing in environments with an existing dhcp server.

