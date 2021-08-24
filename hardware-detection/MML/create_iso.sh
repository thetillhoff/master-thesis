#!/bin/bash

sudo docker build . --tag mllbuild

# privileged=true is required when uefi/both is desired. See https://serverfault.com/questions/701384/loop-device-in-a-linux-container
sudo docker run --privileged=true --rm -it -v $PWD/container:/container mllbuild
