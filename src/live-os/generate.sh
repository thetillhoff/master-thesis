#!/bin/sh

sudo docker build . --tag live-os-builder

sudo docker run --rm -it -v $PWD/container:/container live-os-builder
