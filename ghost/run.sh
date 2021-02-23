#!/bin/bash
set -eoux pipefail
which gh-ost
netstat -peanut | grep 330
gh-ost \
    -verbose \
    -debug \
    -stack \
    -allow-on-master \
    -assume-rbr \
    -host 127.0.0.1 \
    -port 3308 \
    -database try_gh_ost \
    -user repl_user_replica \
    -password toor \
    -assume-master-host 127.0.0.1:3307 \
    -master-user root \
    -master-password toor \
    -table baseitem \
    -alter "MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;"
