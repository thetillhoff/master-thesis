# Ubuntu

While ubuntu wasn't my first choice due to its size and "bloat", I wanted to see at least one version working and the ecosystem around ubuntu is probably one of the best (for linux systems).

In particular, I saw a reference to the software "Cubic", which goes along something similar to "custom ubuntu iso creator".
I tested it  - and what should I say; It worked perfectly on the first testrun.

## What it contains
- Updated apt-packages
  `apt update && apt upgrade -y && apt autoremove -y`
- Preinstalled `lighttpd`
  `apt install -y lighttpd`
- systemd-service which starts a script after boot at `/etc/systemd/system/hwinfo.service`
  ```
  [Unit]
  After=network.service
  [Service]
  ExecStart=/root/hwinfo.sh
  [Install]
  WantedBy=default.target
  ```
- enabled said systemd-service
  `systemctl enable hwinfo.service`
- said (executable) script at `/root/hwinfo.sh`:
  ```
  echo "hi" > /var/www/html/index.html
  cat /proc/cpuinfo > /var/www/html/cpuinfo.html
  cat /proc/meminfo > /var/www/html/meminfo.html
  ```

## Specs
The size of the iso is 1.3GB.

It was tested on:
- Hyper-V on windows 10
- VMware Workstation on Ubuntu 20.04
  I measured the time until the hwinfos where reachable: Around 40s
