#!/bin/bash
set -eoux pipefail
which pt-online-schema-change
netstat -peanut | grep 330
# export PTDEBUG=1
pt-online-schema-change \
    --alter "MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;" \
    --progress time,1 \
    --print \
    --alter-foreign-keys-method rebuild_constraints \
    --user root \
    --password toor \
    --execute \
    h=127.0.0.1,P=3307,u=root,p=toor,D=test_schema,t=baseitem 2>&1
