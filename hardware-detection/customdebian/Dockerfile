FROM ubuntu:20.04

RUN apt update \
  && apt upgrade -y \
  && apt install -y \
  wget \
  xorriso \
  cpio \
  squashfs-tools \
  dpkg-dev \
  syslinux-utils \
  isolinux \
  && apt clean

COPY generate_iso.sh /root/generate_iso.sh

ENV ISOSRC=https://cdimage.debian.org/debian-cd/current-live/amd64/iso-hybrid/debian-live-11.0.0-amd64-standard.iso

ENV ISODST=/container/debian-live-11.0.0-custom.iso

CMD ["/root/generate_iso.sh"]
