#!/bin/sh

# Extract filename from source ISO
export ISONAME="${ISOSRC##*/}"

# If source ISO isn't downloaded yet
if [ ! -f /container/${ISONAME} ]; then
  # Download source ISO to volume, so on rerun it doesn't need to download again
  wget ${ISOSRC} -O /container/${ISONAME}
fi

# Extract source ISO
export EX_ISO=/tmp/extractedISO
mkdir -p ${EX_ISO}
xorriso -osirrox on -indev /container/${ISONAME} -extract / ${EX_ISO}

# Unpacking squashfs
export UNSQUASHEDFS=/tmp/unsquashedfs
mkdir -p ${UNSQUASHEDFS}
unsquashfs -f -d ${UNSQUASHEDFS} ${EX_ISO}/live/filesystem.squashfs

# Prepare editing live system
cp /etc/resolv.conf ${UNSQUASHEDFS}/etc/
cp /container/chroot.sh ${UNSQUASHEDFS}/
#mount -t proc -o bind /proc ${UNSQUASHEDFS}/proc
mount -o bind /dev/pts ${UNSQUASHEDFS}/dev/pts
#mount none -t devpts /dev/pts

# Edit live system
chroot ${UNSQUASHEDFS} /chroot.sh # for manual edits run `chroot ${UNSQUASHEDFS}` and later on `exit`
# TODO install open-ssh as well?
#bash # for debugging; have shell open for manual changes, after `exit`, this script continues

# Finish editing live system
umount "${UNSQUASHEDFS}/dev/pts"
rm ${UNSQUASHEDFS}/chroot.sh
rm ${UNSQUASHEDFS}/etc/resolv.conf

# Repacking squashfs (and overwrite the previous one)
mksquashfs ${UNSQUASHEDFS}/ ${EX_ISO}/live/filesystem.squashfs -comp xz -noappend

# Editing isolinux configuration (boot menu for bios)
#   The only change is for timeout; no timeout -> 1s timeout (units of 1/10s)
sed -i 's/^timeout 0$/timeout 10/g' ${EX_ISO}/isolinux/isolinux.cfg

# Editing grub2 configuration (boot menu for efi)
#   default=0 makes the first option the default, timeout=1 (unit of 1s)
sed -i '$adefault=0\ntimeout=1' ${EX_ISO}/boot/grub/grub.cfg

# Recreating hash (required for efi-boot)
cd ${EX_ISO}
find . -type f -print0 | xargs -0 md5sum | tee ${EX_ISO}/md5sum.txt
cd -

# Repacking SRCISO to new iso - with bios- and efi-boot-support
xorriso -as mkisofs -o ${ISODST} -isohybrid-mbr /usr/lib/ISOLINUX/isohdpfx.bin -c isolinux/boot.cat -b isolinux/isolinux.bin -no-emul-boot -boot-load-size 4 -boot-info-table -eltorito-alt-boot -e boot/grub/efi.img -no-emul-boot -isohybrid-gpt-basdat ${EX_ISO}/

#bash # for debugging; leave shell open after everything else
