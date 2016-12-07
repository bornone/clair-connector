# set base image
FROM postgres:latest

MAINTAINER XiaoLiang Hu <xlhuxa@cn.ibm.com>

Run apt-get update \
    && apt-get install -y wget \
    && wget -P /tmp 9.111.250.208/clair.data 2>/dev/null \
    && echo "#!/bin/bash" >> /docker-entrypoint-initdb.d/init.sh \
    && echo "/usr/lib/postgresql/9.6/bin/psql -h localhost -U postgres -d postgres < /tmp/clair.data" >> /docker-entrypoint-initdb.d/init.sh
