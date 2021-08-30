##########
FROM ubuntu:20.04 AS ipxe-build

# install required software
RUN apt-get update \
  && apt-get install -y \
  git \
  binutils \
  gcc \
  liblzma-dev \
  make \
  mtools \
  genisoimage \
# perl-base on ubuntu, perl on debian
  perl-base \
  syslinux \
  && apt-get clean

# get ipxe src
RUN git clone --depth 1 -b v1.20.1 git://git.ipxe.org/ipxe.git /home

# set workdir
WORKDIR /home/src

## configure ipxe compile settings
### informations taken from https://github.com/SuperClutch/docker-ipxe-build/blob/master/Dockerfile
RUN sed -i "s%//#define POWEROFF_CMD%#define POWEROFF_CMD%" /home/src/config/general.h

# startup command
RUN make bin-x86_64-efi/ipxe.efi && \
  mv bin-x86_64-efi/ipxe.efi /home/; \
  make bin/undionly.kpxe && \
  mv bin/undionly.kpxe /home/


##########
FROM ubuntu:20.04

# install required software
RUN apt-get update \
  && apt upgrade -y \
  && apt install -y \
  wget \
  net-tools \
  dnsmasq \
  #nginx
  && apt-get clean

# Create dedicated folder for tftp files
RUN mkdir -p /tftp /http

# retrieve files from previous stages
RUN wget https://github.com/thetillhoff/serve/releases/download/v0.1.1/serve_darwin_amd64 -O /usr/local/bin/serve && chmod +x /usr/local/bin/serve
COPY --from=ipxe-build /home/ipxe.efi /tftp/
COPY --from=ipxe-build /home/undionly.kpxe /tftp/

## add nginx config file
#COPY ./nginx_site_config /etc/nginx/sites-available/site
## enable nginx site and remove default site
#RUN rm /etc/nginx/sites-enabled/default && ln -s /etc/nginx/sites-available/site /etc/nginx/sites-enabled/
EXPOSE 80/tcp

## add dnsmasq config file
COPY ./dnsmasq.conf /etc/dnsmasq.conf
# 53 dns
# 69/udp tftp
EXPOSE 53/udp
EXPOSE 69/udp
EXPOSE 69/tcp

# add default ipxe-boot-file
RUN mkdir -p /http
COPY ./default /http/default

# add startup script file
COPY ./container_init.sh /container_init.sh

ENV SUBNET=192.168.1.0
ENV ISOSRC=https://cdimage.debian.org/debian-cd/current-live/amd64/iso-hybrid/debian-live-11.0.0-amd64-standard.iso
# available options for DHCP: on, proxy
ENV DHCP=proxy

# startup command
CMD /container_init.sh
