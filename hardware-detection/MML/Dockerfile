FROM ubuntu:20.04

RUN apt update \
  && apt upgrade -y \
  && apt install -y \
    # dependencies for building mll:
    wget \
    make \
    gawk \
    gcc \
    bc \
    bison \
    flex \
    xorriso \
    libelf-dev \
    libssl-dev \
    # own additions (needed for building mll as well):
    xz-utils \
    rsync \
    python3 \
    dosfstools \
  && apt autoremove -y \
  && apt clean \
  # Set python3 as default - not sure if necessary
  && ln -s /usr/bin/python3 /usr/bin/python

RUN wget http://github.com/ivandavidov/minimal/releases/download/15-Dec-2019/minimal_linux_live_15-Dec-2019_src.tar.xz -O /root/mll.tar.xz \
  && tar xf /root/mll.tar.xz -C /root/ --strip 1 \
  && rm /root/mll.tar.xz

# Inserting additional parameter for building glibc. Known error, which should already be fixed in newer versions, but still seems to exist for me.
# '\ ' escapes the space character, '\\' escapes the backslash character
# https://github.com/ivandavidov/minimal/issues/236
RUN sed -i '/^  CFLAGS="$CFLAGS"/i \ \ --enable-cet \\' /root/04_build_glibc.sh

CMD ["/container/generate_iso.sh"]
