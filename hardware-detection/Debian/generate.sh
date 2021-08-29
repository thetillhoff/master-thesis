#!/bin/sh

sudo docker build . --tag geniso

sudo docker run --privileged --rm -it -v $PWD/container:/container geniso #bash
