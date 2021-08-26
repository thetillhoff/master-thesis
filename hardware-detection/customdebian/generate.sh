#!/bin/sh

sudo docker build . --tag geniso

# --privileged needed for schroot
sudo docker run --privileged --rm -it -v $PWD/container:/container geniso #bash
