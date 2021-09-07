#!/bin/sh

export ISONAME="${ISOSRC##*/}"

# alternative value for BINDIP, but since the demo-network is not the physical one, this doesn't work here
#export BINDIP="$(hostname -I | cut -d' ' -f1)"
#echo "IPs: $(hostname -I)"

# If source ISO isn't downloaded yet
if [ ! -f /http/isos/${ISONAME} ]; then
  # # Download source ISO to volume, so on rerun it doesn't need to download again
  # wget ${ISOSRC} -O /http/isos/${ISONAME} --show-progress
  echo "The iso-file at /http/isos/${ISONAME} does not exist!"
  exit 1
fi
echo "The path to the iso-file is '${ISONAME}'."

# Configure DHCP for current ip
sed -i "s%listen-address=.*$%listen-address=${BINDIP}%" /etc/dnsmasq.conf

# configure PXE for current ip
sed -i "s%dhcp-boot=tag:ipxe,http://xxx.xxx.xxx.xxx/default%dhcp-boot=tag:ipxe,http://${BINDIP}/default?mac=\${net0/mac:hexhyp}%" /etc/dnsmasq.conf

# Configure DHCP (range & proxy or not)
#   Assumption: Mask is 255.255.255.0
if [ "${DHCP}" = "proxy" ]; then
  # proxy DHCP configuration requires subnet and ',proxy'
  sed -i "s%dhcp-range=.*$%dhcp-range=${BINDIP},proxy%" /etc/dnsmasq.conf
elif [ "${DHCP}" = "on" ]; then
  # active DHCP configuration requires lower and upper bound of range and lease-time
  sed -i "s%dhcp-range=.*$%dhcp-range=$(echo ${BINDIP} | cut -d'.' -f1-3).2,$(echo ${BINDIP} | cut -d'.' -f1-3).254,2m%" /etc/dnsmasq.conf
else
  echo "Invalid value for \$DHCP: ${DHCP}."
  exit 1
fi

# start dnsmasq with config file and as root (since no dnsmasq-user exists)
dnsmasq -C /etc/dnsmasq.conf -u root
echo "DHCP running on IP ${BINDIP}."

# configure ipxe for current ip
sed -i "s%set boot-url http://.*$%set boot-url http://${BINDIP}/%" /http/default
# configure ipxe for live-iso
sed -i "s%sanboot --no-describe \${boot-url}.*$%sanboot --no-describe \${boot-url}isos/${ISONAME} || goto failed%" /http/default

# start `serve`
/usr/local/bin/serve -d /http -i ${BINDIP} -p 80
