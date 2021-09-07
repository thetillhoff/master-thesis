#!/bin/sh

sudo docker build . --tag dnsmasq \
&& \
sudo docker run \
  --net=host \
  --cap-add=NET_ADMIN \
  --rm -it \
  -v $PWD/isos:/http/isos \
  -e BINDIP=192.168.122.1 \
  -e DHCP=on \
  dnsmasq bash
