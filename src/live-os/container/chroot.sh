#!/bin/sh

# installing lighttpd
apt update
apt upgrade -y
apt autoremove -y
apt install -y lighttpd openssh-server
apt clean

# creating autostart script
echo """#!/bin/sh
cat /proc/cpuinfo > /var/www/html/cpuinfo.txt
cat /proc/meminfo > /var/www/html/meminfo.txt
""" > /root/autostart.sh
chmod +x /root/autostart.sh

# creating service file
echo """[Unit]
After=network.service

[Service]
ExecStart=/root/autostart.sh

[Install]
WantedBy=default.target
""" > /etc/systemd/system/autostart.service

# manually enabling systemd service
ln -s /etc/systemd/system/autostart.service /etc/systemd/system/multi-user.target.wants/autostart.service

# Set custom hostname
# hostnamectl set-hostname live-os # Doesn't work during chroot, since it's not booted with systemd
