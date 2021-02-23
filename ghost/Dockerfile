FROM ubuntu:20.04
ENV GHOST_VERSION 1.1.0
COPY download.sh /tmp/
RUN apt-get update && \
    apt-get install --no-install-recommends -y wget net-tools telnet netcat && \
    chmod a+x /tmp/download.sh && /tmp/download.sh && \
    rm -rf /var/lib/apt/lists/*
