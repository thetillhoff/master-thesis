#!/bin/sh

virt-install \
  --name=pxe-vm \
  --vcpus=2 \
  --memory=2048 \
  --disk size=10 \
  --cdrom $PWD/dnsmasq/isos/debian-live-11.1.0-custom.iso \
  --boot cdrom,hd,menu=on
