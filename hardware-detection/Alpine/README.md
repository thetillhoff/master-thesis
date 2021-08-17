# Alpine

Since there is a tutorial in the alpine wiki (https://wiki.alpinelinux.org/wiki/How_to_make_a_custom_ISO_image_with_mkimage), my initial thought

## Building the iso
Execute `./create_iso.sh`, wait, enjoy.

## setting up network (manually)
ifconfig eth0 up
udhcpc eth0

## setting up network (properly)
```
auto eth0
  iface eth0 inet dhcp
```

## adding online apk repository
`vi /etc/apk/repositories`
i -> start edit (insert mode)
ESC -> stop edit (back to nomal mode)
:wq -> save exit
http://dl-cdn.alpinelinux.org/alpine/edge/main

## update apk sources
apk update

## install and enable lighttpd
apk add lighttpd
rc-update add lighttpd default

## enable root ssh
vi /etc/ssh/sshd_config
change `#PermitRootLogin prohibit-password` to `PermitRootLogin yes`
service sshd restart

## Problems
I wasn't yet able to insert files like scripts, because I wasn't even able to install additional packages (or even build the iso in a way that is bootable).
Strangely, the generated iso files (even with 'standard' profile) are almost 4 times larger than the actual standard iso downloaded from the alpine website.
Additionally, and quite aadly, Hyper-V can't boot those images, since vmlinuz seems to be missing.
To resolve that problem I tried several things:
- Run from an alpine machine (VM) -> no change
- Run several different (predefined) profiles -> no change
- Manually inserted the vmlinuz file (3 different ways) -> rendered the iso unbootable (grub doesn't know what to do any more)
