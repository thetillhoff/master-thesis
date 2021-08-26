#!/bin/sh

# installing lighttpd
apt update
apt install -y lighttpd
apt clean

# creating autostart script
echo """
#!/bin/sh
cat /proc/cpuinfo > /var/www/html/cpuinfo.txt
cat /proc/meminfo > /var/www/html/meminfo.txt
""" > /root/autostart.sh
chmod +x /root/autostart.sh

# creating service file
echo """
[Unit]
After=network.service

[Service]
ExecStart=/root/autostart.sh

[Install]
WantedBy=default.target
""" > /etc/systemd/system/autostart.service

# manually enabling systemd service
ln -s /etc/systemd/system/autostart.service /etc/systemd/system/multi-user.target.wants/autostart.service
