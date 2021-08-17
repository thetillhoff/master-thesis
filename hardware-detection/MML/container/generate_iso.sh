#!/bin/bash

# Copying files for new overlay
cp /container/autostart /root/minimal_overlay/bundles/autostart -r

# Building new overlay
cd /root/minimal_overlay/ && /root/minimal_overlay/overlay_build.sh autostart

# Adding overlay_bundles
sed -i 's/OVERLAY_BUNDLES=dhcp,mll_hello,mll_logo,mll_source/OVERLAY_BUNDLES=dhcp,mll_hello,mll_logo,mll_source,nweb,autostart/g' /root/.config

# Setting firmware type
sed -i 's/FIRMWARE_TYPE=bios/FIRMWARE_TYPE=both/g' /root/.config

# Building mml iso
cd /root && /root/build_minimal_linux_live.sh

# Copying mml iso to host
cp /root/minimal_linux_live.iso /container/mml.iso
