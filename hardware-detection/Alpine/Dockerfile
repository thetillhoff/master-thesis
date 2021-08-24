FROM alpine:edge

RUN apk update

# Make sure we have the needed tools
RUN apk add \
  alpine-sdk \
  build-base \
  apk-tools \
  alpine-conf \
  busybox \
  fakeroot \
  syslinux \
  xorriso \
  squashfs-tools \
  sudo \
  # For efi you should add the following
  mtools \
  dosfstools \
  grub-efi \
  # Own additions
  make

# Create a user (e.g. build) and add them to the abuild group
#RUN adduser -h /home/build -D build -G abuild
RUN adduser build -G abuild

# Give administrative access to the abuild group
#RUN echo "build ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/abuild
RUN echo "%abuild ALL=(ALL) ALL" > /etc/sudoers.d/abuild

# Change to the build user
USER build

# # Then create signing keys (-i installs them in /etc/apk/keys which is required for later)
RUN abuild-keygen -i -a

# # Tip: Make sure your public keys are placed in /etc/apk/keys/ (example: build-xxxxxxxx.rsa.pub)
# RUN ls /etc/apk/keys/

# Clone the git repository
RUN git clone https://gitlab.alpinelinux.org/alpine/aports.git /root

WORKDIR /root/scripts

CMD ["/container/generate_iso.sh"]
