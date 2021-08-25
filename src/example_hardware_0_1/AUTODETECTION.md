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
