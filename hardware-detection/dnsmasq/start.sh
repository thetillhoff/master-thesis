#!/bin/sh

sudo docker build . --tag dnsmasq \
&& \
sudo docker run \
  --net=host \
  --cap-add=NET_ADMIN \
  --rm -it \
  -v $PWD/http:/http \
  dnsmasq #bash
