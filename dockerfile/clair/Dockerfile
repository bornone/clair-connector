# set base image
FROM quay.io/coreos/clair:latest

MAINTAINER XiaoLiang Hu <xlhuxa@cn.ibm.com>

# set crt for clair to access to cfc docker registry, since current mount host dir doesn't
# work due to the host path /etc/docker/certs.d/master.cfc:8500/ca.crt contains colon

RUN wget -P /usr/local/share/ca-certificates 9.111.250.208/ca.crt 2>/dev/null \
    && update-ca-certificates -f