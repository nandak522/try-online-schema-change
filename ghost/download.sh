#!/bin/bash
set -eoux pipefail

url="https://github.com/github/gh-ost/releases/download/v${GHOST_VERSION}/gh-ost_${GHOST_VERSION}_amd64.deb"
wget --no-check-certificate -q "${url}" -O "/tmp/gh-ost_${GHOST_VERSION}_amd64.deb"
dpkg -i /tmp/gh-ost_${GHOST_VERSION}_amd64.deb
