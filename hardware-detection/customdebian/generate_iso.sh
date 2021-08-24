#!/bin/sh

# Extract filename from SRCISO
export ISONAME="${ISOSRC##*/}"

# If SRCISO isn't downloaded yet
if [ ! -f /container/${ISONAME} ]; then
  # Download SRCISO
  wget ${ISOSRC} -O /container/${ISONAME}
fi

# Extract SRCISO
export EX_SRCISO=/tmp/extractedSRCISO
mkdir -p ${EX_SRCISO}
xorriso -osirrox on -indev /container/${ISONAME} -extract / ${EX_SRCISO}

# Extracting squashfs
mkdir -p /tmp/squashedfs
unsquashfs -f -d /tmp/squashedfs ${EX_SRCISO}/live/filesystem.squashfs

cp /etc/resolv.conf /tmp/squashedfs/etc/
#mount -t proc -o bind /proc /tmp/squashedfs/proc
mount -o bind /dev/pts /tmp/squashedfs/dev/pts
chroot /tmp/squashedfs # ...
# apt update && apt upgrade -y && apt install -y lighttpd && apt autoremove -y && apt clean
rm /tmp/squashedfs/etc/resolv.conf



# # Editing initrd contents (==filesystem of the installer)
# #echo /container/preseed.cfg | cpio -H newc -o -A -F ${EX_SRCISO}/install.amd/initrd
# #debootstrap --arch=amd64 stable /tmp/initrd http://deb.debian.org/debian/
# chroot /tmp/squashedfs
# cat << EOF | schroot -c custom .
# touch /yoloman
# apt update
# apt install -y \
#   lighttpd \
#   openssh-server
# apt clean
# EOF

bash


# # Regenerating new md5sum.txt
# cd ${EX_SRCISO}
# chmod +w ${EX_SRCISO}/md5sum.txt
# find ${EX_SRCISO}/ -follow -type f ! -name md5sum.txt -print0 | xargs -0 md5sum > ${EX_SRCISO}/md5sum.txt
# chmod -w ${EX_SRCISO}/md5sum.txt

# # Extract first 432 bytes of original iso (they contain an MBR, which enables booting from usb-sticks on legacy BIOS)
# export MBR_TEMPLATE=/mbr_template.bin
# dd if="/container/${ISONAME}" bs=1 count=432 of="/tmp/${MBR_TEMPLATE}"

# # Creating a new bootable iso image
# # genisoimage -r -J -b isolinux/isolinux.bin -c isolinux/boot.cat \
# #   -no-emul-boot -boot-load-size 4 -boot-info-table \
# #   -o /container/preseed-debian-10.2.0-i386-netinst.iso ${EX_SRCISO}
# # ---
# # Create the new ISO image
# xorriso -as mkisofs \
#   -r -V 'Debian_11_amd64_Preseeded' \
#   -o "${ISODST}" \
#   -J -joliet-long -cache-inodes \
#   -isohybrid-mbr /tmp/${MBR_TEMPLATE} \
#   -b isolinux/isolinux.bin \
#   -c isolinux/boot.cat \
#   -boot-load-size 4 -boot-info-table -no-emul-boot \
#   -eltorito-alt-boot \
#   -e boot/grub/efi.img \
#   -no-emul-boot \
#   -isohybrid-gpt-basdat -isohybrid-apm-hfsplus \
#   "${EX_SRCISO}"

# # xorriso -as mkisofs 
# #   -r 
# # #  -checksum_algorithm_iso sha256,sha512 
# #   -V 'Debian 11.0.0 amd64 n' 
# #   -o /srv/cdbuilder.debian.org/dst/deb-cd/out/2bullseyeamd64/debian-11.0.0-amd64-NETINST-1.iso 
# # #  -md5-list /srv/cdbuilder.debian.org/src/deb-cd/tmp/2bullseyeamd64/bullseye/checksum-check 
# #   -J -joliet-long -cache-inodes
# #   -isohybrid-mbr syslinux/usr/lib/ISOLINUX/isohdpfx.bin 
# #   -b isolinux/isolinux.bin 
# #   -c isolinux/boot.cat 
# #   -boot-load-size 4 -boot-info-table -no-emul-boot 
# #   -eltorito-alt-boot 
# #   -e boot/grub/efi.img 
# #   -no-emul-boot 
# #   -isohybrid-gpt-basdat -isohybrid-apm-hfsplus boot1 CD1
