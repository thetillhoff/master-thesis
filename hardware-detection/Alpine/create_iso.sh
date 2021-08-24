#!/bin/bash

sudo docker build . --tag alpinebuild

sudo docker run --rm -it -v $PWD/container:/container alpinebuild
