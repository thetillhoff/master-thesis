#!/bin/sh

mkdir -p /web/

cat /proc/cpuinfo > /srv/www/cpuinfo.txt
cat /proc/meminfo > /srv/www/meminfo.txt
