FROM ubuntu:20.04
RUN apt-get update && \
    apt-get install --no-install-recommends -y wget net-tools telnet netcat gnupg2 lsb-release curl && \
    wget --no-check-certificate https://repo.percona.com/apt/percona-release_latest.generic_all.deb && \
    dpkg -i percona-release_latest.generic_all.deb && \
    cat /etc/apt/sources.list.d/percona-original-release.list
RUN apt-get update && apt-get install -y --no-install-recommends percona-toolkit && \
    rm -rf /var/lib/apt/lists/*
