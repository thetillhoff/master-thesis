#!/bin/sh

export ISONAME="${ISOSRC##*/}"

# If source ISO isn't downloaded yet
if [ ! -f /http/isos/${ISONAME} ]; then
  # Download source ISO to volume, so on rerun it doesn't need to download again
  wget ${ISOSRC} -O /http/isos/${ISONAME} --show-progress
fi

echo "IPs: $(hostname -I)"
export IP=$(hostname -I | cut -d' ' -f1)
echo "Running on IP: ${IP}"

# Set Subnet variable according to host ip
if [ -z "${SUBNET}" ]; then
  if [ "$(echo ${IP} | cut -d'.' -f1)" = "192" ]; then
    # When we are in an 192.168.*.0 network, make sure to set * accordingly since most private networks use subnet mask 255.255.255.0 and not 255.255.0.0
    export SUBNET=$(echo ${IP} | cut -d'.' -f1)\.$(echo ${IP} | cut -d'.' -f2)\.$(echo ${IP} | cut -d'.' -f3)\.\0
  elif [ "$(echo ${IP} | cut -d'.' -f1)" = "10" ]; then
    export SUBNET=10.0.0.0
  fi
fi
echo "Subnet: ${SUBNET}"

# configure ipxe for current ip
sed -i "s%set boot-url http://.*$%set boot-url http://$(hostname -I | cut -d' ' -f1)/%" /http/default
# configure ipxe for iso
if [ ! -z "${ISONAME}" ]; then # The same as [ -n "${ISONAME}" ]
  sed -i "s%sanboot --no-describe \${boot-url}/.*$%sanboot --no-describe \${boot-url}/${ISONAME} || goto failed%" /http/default
fi

# configure dnsmasq for current ip
sed -i "s%dhcp-boot=tag:ipxe,http://xxx.xxx.xxx.xxx/default%dhcp-boot=tag:ipxe,http://$(hostname -I | cut -d' ' -f1)/default?mac=\${net0/mac:hexhyp}%" /etc/dnsmasq.conf

# configure dnsmasq's dhcp & its subnet
if [ -z "${SUBNET}" ]; then
  echo "Invalid value for \$SUBNET variable: ${SUBNET}"
  exit 1
fi
if [ "${DHCP}" = "proxy" ]; then
  sed -i "s%dhcp-range=.*$%dhcp-range=${SUBNET},proxy%" /etc/dnsmasq.conf
elif [ "${DHCP}" = "on" ]; then
  sed -i "s%dhcp-range=.*$%dhcp-range=${SUBNET}%" /etc/dnsmasq.conf
else
  echo "Invalid value for \$DHCP variable: ${DHCP}."
  exit 1
fi

# start dnsmasq with config file and as root (since no dnsmasq-user exists)
dnsmasq -C /etc/dnsmasq.conf -u root

# start `serve`
/usr/local/bin/serve -d /http -p 80
