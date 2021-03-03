#!/bin/bash
set -eoux pipefail
# which gh-ost
# netstat -peanut | grep 330

# # Approach 1 : Bails out. Either with parent-side or child-side of FKs.
# gh-ost \
#     -verbose \
#     -debug \
#     -stack \
#     -allow-on-master \
#     -assume-rbr \
#     -host 127.0.0.1 \
#     -port 3308 \
#     -database try_osc \
#     -user repl_user_replica \
#     -password toor \
#     -assume-master-host 127.0.0.1:3307 \
#     -master-user root \
#     -master-password toor \
#     -table baseitem \
#     -alter "MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;"

# Approach-2 => Drop FKs first
# ALTER TABLE try_osc.referringitem DROP FOREIGN KEY baseitem_id_refs_id_2d6ba49a;

# # Continuing Approach-2 => Evolve child table(s) next
gh-ost \
    -verbose \
    -debug \
    -stack \
    -allow-on-master \
    -assume-rbr \
    -host 127.0.0.1 \
    -port 3308 \
    -database try_osc \
    -user repl_user_replica \
    -password toor \
    -assume-master-host 127.0.0.1:3307 \
    -master-user root \
    -master-password toor \
    -table referringitem \
    -alter "MODIFY baseitem_id bigint(20) NOT NULL;" \
    --execute

# # Continuing Approach-2 => Evolve parent table next
gh-ost \
    -verbose \
    -debug \
    -stack \
    -allow-on-master \
    -assume-rbr \
    -host 127.0.0.1 \
    -port 3308 \
    -database try_osc \
    -user repl_user_replica \
    -password toor \
    -assume-master-host 127.0.0.1:3307 \
    -master-user root \
    -master-password toor \
    -table baseitem \
    -alter "MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;" \
    --execute

# # Continuing Approach-2 => Add FKs back
# # mysql>
# # -- With this flag, re-adding the FKs is very quick because it doesn't have to evaluate if there are any violations.
# set foreign_key_checks=off;
# ALTER TABLE try_osc.referringitem ADD CONSTRAINT baseitem_id_refs_id_2d6ba49a FOREIGN KEY  (baseitem_id) REFERENCES try_osc.baseitem (id);
# set foreign_key_checks=on;

# BIG NOTE: So you have a tiny window where data integrity is not enforced.
# We have to manually fix it.
