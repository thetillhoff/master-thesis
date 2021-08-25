# tests

## opentosca
- `wget -qO- http://install.opentosca.org/install-dockerized | TAG=v3.0.0 sh`
- winery doesn't work -> 500 errors

## cloudify in docker
- sudo docker run --name cfy_manager_local -d --restart unless-stopped -v /sys/fs/cgroup:/sys/fs/cgroup:ro --tmpfs /run --tmpfs /run/lock --security-opt seccomp:unconfined --cap-add SYS_ADMIN -p 80:80 -p 8000:8000 cloudifyplatform/community-cloudify-manager-aio:latest
- works
- good onboarding

## openstack-devstack on hyperv
- works only on some os ? (neither debian10 - not supported & installation error, nor ubuntu 18.04 - not supported & installation error, nor opensuse tumbleweed (os doesn't work properly (on hyperv)) - even though the latter two should work)

## openstack on centos on hyperv
- requires erlang, erlang can't be installed

## openstack-devstack on centos on hyperv
- required manual setting of permissions of /opt with -R to root:root and 0755 -> "can't access *.log
- required manual setting of permissions of /opt/stack with -R to 0777

## aws
- works
- cloudify with aws works
