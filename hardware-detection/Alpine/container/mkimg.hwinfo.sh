#!/bin/sh

profile_hwinfo() {
  profile_standard
  kernel_cmdline="unionfs_size=512M console=tty0 console=ttyS0,115200"
  syslinux_serial="0 115200"
  apks="$apks lighttpd"

  local _k _a
  for _k in $kernel_flavors; do
    apks="$apks linux-$_k"
    for _a in $kernel_addons; do
      apks="$apks $_a-$_k"
    done
  done
  apks="$apks linux-firmware"
}
