#!/bin/bash
set -eoux pipefail
# which pt-online-schema-change
# netstat -peanut | grep 330
# export PTDEBUG=1

# Approach-1 => Straight approach.
# # Not possible. Errors out if FKs are there
# pt-online-schema-change \
#     --alter "MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;" \
#     --progress time,1 \
#     --print \
#     --alter-foreign-keys-method rebuild_constraints \
#     --user root \
#     --password toor \
#     --execute \
#     h=127.0.0.1,P=3307,u=root,p=toor,D=try_osc,t=baseitem 2>&1

# Approach-2 => Drop FKs first
# ALTER TABLE try_osc.referringitem DROP FOREIGN KEY baseitem_id_refs_id_2d6ba49a;

# Continuing Approach-2 => Evolve child table(s) next
pt-online-schema-change \
    --alter "MODIFY baseitem_id bigint(20) NOT NULL;" \
    --progress time,1 \
    --print \
    --alter-foreign-keys-method rebuild_constraints \
    --user root \
    --password toor \
    --execute \
    h=127.0.0.1,P=3307,u=root,p=toor,D=try_osc,t=referringitem 2>&1

# Continuing Approach-2 => Evolve parent table next
pt-online-schema-change \
    --alter "MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;" \
    --progress time,1 \
    --print \
    --alter-foreign-keys-method rebuild_constraints \
    --user root \
    --password toor \
    --execute \
    h=127.0.0.1,P=3307,u=root,p=toor,D=try_osc,t=baseitem 2>&1

# Continuing Approach-2 => Add FKs back
# mysql>
# -- With this flag, re-adding the FKs is very quick because it doesn't have to evaluate if there are any violations.
# set foreign_key_checks=off;
# ALTER TABLE try_osc.referringitem ADD CONSTRAINT _baseitem_id_refs_id_2d6ba49a FOREIGN KEY  (baseitem_id) REFERENCES try_osc.baseitem (id);
# set foreign_key_checks=on;

# BIG NOTE: So you have a tiny window where data integrity is not enforced.
# We have to manually fix it.
