#!/bin/sh

cp /container/mkimg.hwinfo.sh .

chmod +x ./mkimg.hwinfo.sh

./mkimage.sh --tag edge --arch x86_64 --outdir /tmp --profile hwinfo --repository http://dl-cdn.alpinelinux.org/alpine/edge/main

cp /tmp/alpine-*.iso /container/alpine-custom.iso
