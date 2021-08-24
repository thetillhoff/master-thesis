#!/bin/sh

#cp /container/mkimg.hwinfo.sh .
cp /container/mkimg.nas.sh .

#chmod +x ./mkimg.hwinfo.sh
chmod +x ./mkimg.nas.sh

#./mkimage.sh --tag edge --arch x86_64 --outdir /tmp --profile hwinfo --repository http://dl-cdn.alpinelinux.org/alpine/edge/main
./mkimage.sh --tag edge --arch x86_64 --outdir /tmp --profile nas --repository http://dl-cdn.alpinelinux.org/alpine/edge/main

cp /tmp/alpine-*.iso /container/alpine-custom.iso
