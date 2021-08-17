> While the iso is successfully generated, it seems to have issues with booting (at least on hyper-v).

# Minimal Linux Live (MLL)
I decided for this OS, since it ticks all those boxes from above.
It supports BIOS & UEFI, is super well documented, intended to be easily extensible, small and can be edited to boot super fast.

## Building the iso
Execute `./create_iso.sh`, wait, enjoy.

## How it works
The MLL iso is generated with additional bundles here.
One is nweb, a super-simple web-server.
The other is self-made and called "autostart" (quite creative, I know).
Both are installed/enabled on the custom-iso.

## Building the iso
Execute `./create_iso.sh`, wait, enjoy.

## Problems
While the original downloaded iso at least boots, keyboard input doesn't work.
Sadly, Hyper-V isn't able to boot those generated iso images:
```
VFS: Cannot open root device "(null)" or unkown-block(0,0): error -6
Please append the correct "root=" boot option: here are the available partitions:
Kernel panic - not syncing:; VFS: Unable to mount root fs on unkown-block(0,0)
```