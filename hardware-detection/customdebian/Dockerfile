FROM ubuntu:20.04

RUN apt update \
  && apt upgrade -y \
  && apt install -y \
  wget \
  xorriso \
  cpio \
  #genisoimage \
  #debootstrap \
  #schroot \
  squashfs-tools \
  dpkg-dev \
  && apt clean

COPY generate_iso.sh /root/generate_iso.sh

#ENV ISOSRC=https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-11.0.0-amd64-netinst.iso
ENV ISOSRC=https://cdimage.debian.org/debian-cd/current-live/amd64/iso-hybrid/debian-live-11.0.0-amd64-standard.iso

ENV ISODST=/container/debian-preseeded.iso

CMD ["/root/generate_iso.sh"]
