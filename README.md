# try-online-schema-change
```sh
docker-compose up mysql-master mysql-replica
```

Load data using `initial_data.sql`

### Normal vanila alter

Iteration 1 (didn't work)
```sql
ALTER TABLE try_osc.baseitem DROP PRIMARY KEY;
ALTER TABLE try_osc.baseitem MODIFY COLUMN id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY;
```

---
Iteration 2 (didn't work)
```sql
mysql> ALTER TABLE try_osc.baseitem DROP CONSTRAINT PRIMARY;
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'CONSTRAINT PRIMARY' at line 1
mysql> alter table try_osc.baseitem modify column id bigint(20) NOT NULL AUTO_INCREMENT;
ERROR 1833 (HY000): Cannot change column 'id': used in a foreign key constraint 'baseitem_id_refs_id_2d6ba49a' of table 'try_osc.referringitem'
alter table try_osc.baseitem add primary key (id);
```

---
Iteration 3 (didn't work)
```sh
# Using hhost
cd <this folder>
docker-compose up ghost
Building with native build. Learn about native build in Compose here: https://docs.docker.com/go/compose-native-build/
Starting ghost ... done
Attaching to ghost
ghost            | + set -eoux pipefail
ghost            | + which gh-ost
ghost            | + netstat -peanut
ghost            | + grep 330
ghost            | /usr/bin/gh-ost
ghost            | tcp        0      0 127.0.0.1:46642         127.0.0.1:3307          ESTABLISHED 27         32815      -
ghost            | tcp        0      0 127.0.0.1:56198         127.0.0.1:3308          TIME_WAIT   0          0          -
ghost            | tcp6       0      0 :::3308                 :::*                    LISTEN      27         31607      -
ghost            | tcp6       0      0 :::3307                 :::*                    LISTEN      27         32238      -
ghost            | tcp6       0      0 127.0.0.1:3307          127.0.0.1:46642         ESTABLISHED 27         32415      -
ghost            | + gh-ost -verbose -debug -stack -allow-on-master -assume-rbr -host 127.0.0.1 -port 3308 -database try_osc -user repl_user_replica -password toor -assume-master-host 127.0.0.1:3307 -master-user root -master-password toor -table baseitem -alter 'MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;'
ghost            | 2021-02-22 19:09:27 INFO starting gh-ost 1.1.0
ghost            | 2021-02-22 19:09:27 INFO Migrating `try_osc`.`baseitem`
ghost            | 2021-02-22 19:09:27 INFO connection validated on 127.0.0.1:3308
ghost            | 2021-02-22 19:09:27 INFO User has ALL privileges
ghost            | 2021-02-22 19:09:27 INFO binary logs validated on 127.0.0.1:3308
ghost            | 2021-02-22 19:09:27 INFO Inspector initiated on vanila-minikube:3308, version 5.7.33-log
ghost            | 2021-02-22 19:09:27 INFO Table found. Engine=InnoDB
ghost            | 2021-02-22 19:09:27 DEBUG Estimated number of rows via STATUS: 4
ghost            | 2021-02-22 19:09:27 ERROR Found 1 parent-side foreign keys on `try_osc`.`baseitem`. Parent-side foreign keys are not supported. Bailing out
ghost            | 2021-02-22 19:09:27 INFO Tearing down inspector
ghost            | 2021-02-22 19:09:27 FATAL 2021-02-22 19:09:27 ERROR Found 1 parent-side foreign keys on `try_osc`.`baseitem`. Parent-side foreign keys are not supported. Bailing out
ghost            | goroutine 1 [running]:
ghost            | runtime/debug.Stack(0xa8, 0x100, 0xc000184090)
ghost            | 	/usr/local/go/src/runtime/debug/stack.go:24 +0x9d
ghost            | runtime/debug.PrintStack()
ghost            | 	/usr/local/go/src/runtime/debug/stack.go:16 +0x22
ghost            | github.com/github/gh-ost/vendor/github.com/outbrain/golib/log.logErrorEntry(0x0, 0x8951a0, 0xc00001c6c0, 0xc00009b2a0, 0xc00009b290)
ghost            | 	/go/src/github.com/github/gh-ost/vendor/github.com/outbrain/golib/log/log.go:178 +0xe7
ghost            | github.com/github/gh-ost/vendor/github.com/outbrain/golib/log.Fatale(0x8951a0, 0xc00001c6c0, 0x7dd701, 0xc0000fc120)
ghost            | 	/go/src/github.com/github/gh-ost/vendor/github.com/outbrain/golib/log/log.go:255 +0x3e
ghost            | github.com/github/gh-ost/go/base.(*simpleLogger).Fatale(0xb67148, 0x8951a0, 0xc00001c6c0, 0x0, 0x0)
ghost            | 	/go/src/github.com/github/gh-ost/go/base/default_logger.go:62 +0x35
ghost            | main.main()
ghost            | 	/go/src/github.com/github/gh-ost/go/cmd/gh-ost/main.go:296 +0x2a43
ghost exited with code 1
```

---
Iteration 4 (ptonline-schema-change)
```sh
cd <this folder>
docker-compose up ptkit
Building with native build. Learn about native build in Compose here: https://docs.docker.com/go/compose-native-build/
Starting ptkit ... done
Attaching to ptkit
ptkit            | + set -eoux pipefail
ptkit            | + which pt-online-schema-change
ptkit            | /usr/bin/pt-online-schema-change
ptkit            | + netstat -peanut
ptkit            | + grep 330
ptkit            | tcp        0      0 127.0.0.1:46816         127.0.0.1:3307          ESTABLISHED 27         129184     -
ptkit            | tcp6       0      0 :::3308                 :::*                    LISTEN      27         127870     -
ptkit            | tcp6       0      0 :::3307                 :::*                    LISTEN      27         128097     -
ptkit            | tcp6       0      0 127.0.0.1:3307          127.0.0.1:46816         ESTABLISHED 27         128381     -
ptkit            | tcp6       0      0 127.0.0.1:3307          127.0.0.1:46820         TIME_WAIT   0          0          -
ptkit            | tcp6       0      0 127.0.0.1:3307          127.0.0.1:46818         TIME_WAIT   0          0          -
ptkit            | + pt-online-schema-change --alter 'MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;' --progress time,1 --print --alter-foreign-keys-method rebuild_constraints --user root --password toor --execute h=127.0.0.1,P=3307,u=root,p=toor,D=try_osc,t=baseitem
ptkit            | Cannot connect to P=3308,h=,p=...,u=root
ptkit            | No slaves found.  See --recursion-method if host vanila-minikube has slaves.
ptkit            | Not checking slave lag because no slaves were found and --check-slave-lag was not specified.
ptkit            | Operation, tries, wait:
ptkit            |   analyze_table, 10, 1
ptkit            |   copy_rows, 10, 0.25
ptkit            |   create_triggers, 10, 1
ptkit            |   drop_triggers, 10, 1
ptkit            |   swap_tables, 10, 1
ptkit            |   update_foreign_keys, 10, 1
ptkit            | Child tables:
ptkit            |   `try_osc`.`referringitem` (approx. 5 rows)
ptkit            | Will use the rebuild_constraints method to update foreign keys.
ptkit            | Altering `try_osc`.`baseitem`...
ptkit            | Creating new table...
ptkit            | CREATE TABLE `try_osc`.`_baseitem_new` (
ptkit            |   `id` int(11) NOT NULL AUTO_INCREMENT,
ptkit            |   `created_on` datetime NOT NULL,
ptkit            |   `updated_on` datetime NOT NULL,
ptkit            |   `product_id` int(11) NOT NULL,
ptkit            |   PRIMARY KEY (`id`)
ptkit            | ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1
ptkit            | Created new table try_osc._baseitem_new OK.
ptkit            | Altering new table...
ptkit            | ALTER TABLE `try_osc`.`_baseitem_new` MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;
ptkit            | Altered `try_osc`.`_baseitem_new` OK.
ptkit            | 2021-02-23T07:42:38 Creating triggers...
ptkit            | -----------------------------------------------------------
ptkit            | Event : DELETE
ptkit            | Name  : pt_osc_try_osc_baseitem_del
ptkit            | SQL   : CREATE TRIGGER `pt_osc_try_osc_baseitem_del` AFTER DELETE ON `try_osc`.`baseitem` FOR EACH ROW BEGIN DECLARE CONTINUE HANDLER FOR 1146 begin end; DELETE IGNORE FROM `try_osc`.`_baseitem_new` WHERE `try_osc`.`_baseitem_new`.`id` <=> OLD.`id`; END
ptkit            | Suffix: del
ptkit            | Time  : AFTER
ptkit            | -----------------------------------------------------------
ptkit            | -----------------------------------------------------------
ptkit            | Event : UPDATE
ptkit            | Name  : pt_osc_try_osc_baseitem_upd
ptkit            | SQL   : CREATE TRIGGER `pt_osc_try_osc_baseitem_upd` AFTER UPDATE ON `try_osc`.`baseitem` FOR EACH ROW BEGIN DECLARE CONTINUE HANDLER FOR 1146 begin end; DELETE IGNORE FROM `try_osc`.`_baseitem_new` WHERE !(OLD.`id` <=> NEW.`id`) AND `try_osc`.`_baseitem_new`.`id` <=> OLD.`id`; REPLACE INTO `try_osc`.`_baseitem_new` (`id`, `created_on`, `updated_on`, `product_id`) VALUES (NEW.`id`, NEW.`created_on`, NEW.`updated_on`, NEW.`product_id`); END
ptkit            | Suffix: upd
ptkit            | Time  : AFTER
ptkit            | -----------------------------------------------------------
ptkit            | -----------------------------------------------------------
ptkit            | Event : INSERT
ptkit            | Name  : pt_osc_try_osc_baseitem_ins
ptkit            | SQL   : CREATE TRIGGER `pt_osc_try_osc_baseitem_ins` AFTER INSERT ON `try_osc`.`baseitem` FOR EACH ROW BEGIN DECLARE CONTINUE HANDLER FOR 1146 begin end; REPLACE INTO `try_osc`.`_baseitem_new` (`id`, `created_on`, `updated_on`, `product_id`) VALUES (NEW.`id`, NEW.`created_on`, NEW.`updated_on`, NEW.`product_id`);END
debug2: channel 0: window 999185 sent adjust 49391
ptkit            | Suffix: ins
ptkit            | Time  : AFTER
ptkit            | -----------------------------------------------------------
ptkit            | 2021-02-23T07:42:38 Created triggers OK.
ptkit            | 2021-02-23T07:42:38 Copying approximately 4 rows...
ptkit            | INSERT LOW_PRIORITY IGNORE INTO `try_osc`.`_baseitem_new` (`id`, `created_on`, `updated_on`, `product_id`) SELECT `id`, `created_on`, `updated_on`, `product_id` FROM `try_osc`.`baseitem` LOCK IN SHARE MODE /*pt-online-schema-change 10 copy table*/
ptkit            | Cannot connect to P=3308,h=,p=...,u=root
ptkit            | 2021-02-23T07:42:38 Copied rows OK.
ptkit            | 2021-02-23T07:42:38 Analyzing new table...
ptkit            | 2021-02-23T07:42:38 Swapping tables...
ptkit            | RENAME TABLE `try_osc`.`baseitem` TO `try_osc`.`_baseitem_old`, `try_osc`.`_baseitem_new` TO `try_osc`.`baseitem`
ptkit            | 2021-02-23T07:42:38 Swapped original and new tables OK.
ptkit            | 2021-02-23T07:42:38 Rebuilding foreign key constraints...
ptkit            | ALTER TABLE `try_osc`.`referringitem` DROP FOREIGN KEY `baseitem_id_refs_id_2d6ba49a`, ADD CONSTRAINT `_baseitem_id_refs_id_2d6ba49a` FOREIGN KEY (`baseitem_id`) REFERENCES `try_osc`.`baseitem` (`id`)
ptkit            | Error updating foreign key constraints: 2021-02-23T07:42:38 DBD::mysql::db do failed: Cannot add foreign key constraint [for Statement "ALTER TABLE `try_osc`.`referringitem` DROP FOREIGN KEY `baseitem_id_refs_id_2d6ba49a`, ADD CONSTRAINT `_baseitem_id_refs_id_2d6ba49a` FOREIGN KEY (`baseitem_id`) REFERENCES `try_osc`.`baseitem` (`id`)"] at /usr/bin/pt-online-schema-change line 11222.
ptkit            | 	(in cleanup) 2021-02-23T07:42:38 DBD::mysql::db do failed: Cannot add foreign key constraint [for Statement "ALTER TABLE `try_osc`.`referringitem` DROP FOREIGN KEY `baseitem_id_refs_id_2d6ba49a`, ADD CONSTRAINT `_baseitem_id_refs_id_2d6ba49a` FOREIGN KEY (`baseitem_id`) REFERENCES `try_osc`.`baseitem` (`id`)"] at /usr/bin/pt-online-schema-change line 11222.
ptkit            | 2021-02-23T07:42:38 Dropping triggers...
ptkit            | DROP TRIGGER IF EXISTS `try_osc`.`pt_osc_try_osc_baseitem_del`
ptkit            | DROP TRIGGER IF EXISTS `try_osc`.`pt_osc_try_osc_baseitem_upd`
ptkit            | DROP TRIGGER IF EXISTS `try_osc`.`pt_osc_try_osc_baseitem_ins`
ptkit            | 2021-02-23T07:42:38 Dropped triggers OK.
ptkit            | Altered `try_osc`.`baseitem` but there were errors or warnings.
ptkit exited with code 15
```

Iteration 5 (using gh-ost after dropping foreign keys)
```sh
cd <this folder>
docker-compose up ghost
Building with native build. Learn about native build in Compose here: https://docs.docker.com/go/compose-native-build/
Starting ghost ... done
Attaching to ghost
ghost            | + set -eoux pipefail
ghost            | + gh-ost -verbose -debug -stack -allow-on-master -assume-rbr -host 127.0.0.1 -port 3308 -database try_osc -user repl_user_replica -password toor -assume-master-host 127.0.0.1:3307 -master-user root -master-password toor -table referringitem -alter 'MODIFY baseitem_id bigint(20) NOT NULL;' --execute
ghost            | 2021-04-01 17:33:57 INFO starting gh-ost 1.1.0
ghost            | 2021-04-01 17:33:57 INFO Migrating `try_osc`.`referringitem`
ghost            | 2021-04-01 17:33:57 INFO connection validated on 127.0.0.1:3308
ghost            | 2021-04-01 17:33:57 INFO User has ALL privileges
ghost            | 2021-04-01 17:33:57 INFO binary logs validated on 127.0.0.1:3308
ghost            | 2021-04-01 17:33:57 INFO Inspector initiated on vanila-minikube:3308, version 5.7.33-log
ghost            | 2021-04-01 17:33:57 INFO Table found. Engine=InnoDB
ghost            | 2021-04-01 17:33:57 DEBUG Estimated number of rows via STATUS: 25
ghost            | 2021-04-01 17:33:57 DEBUG Validated no foreign keys exist on table
ghost            | 2021-04-01 17:33:57 DEBUG Validated no triggers exist on table
ghost            | 2021-04-01 17:33:57 INFO Estimated number of rows via EXPLAIN: 25
ghost            | 2021-04-01 17:33:57 DEBUG Potential unique keys in referringitem: [PRIMARY (auto_increment): [id]; has nullable: false]
ghost            | 2021-04-01 17:33:57 INFO Master forced to be 127.0.0.1:3307
ghost            | 2021-04-01 17:33:57 INFO log_slave_updates validated on 127.0.0.1:3308
ghost            | 2021-04-01 17:33:57 INFO connection validated on 127.0.0.1:3308
ghost            | [2021/04/01 17:33:57] [info] binlogsyncer.go:133 create BinlogSyncer with config {99999 mysql 127.0.0.1 3308 repl_user_replica    false false <nil> false UTC true 0 0s 0s 0 false}
ghost            | [2021/04/01 17:33:57] [info] binlogsyncer.go:354 begin to sync binlog from position (bin.000003, 78817)
ghost            | [2021/04/01 17:33:57] [info] binlogsyncer.go:203 register slave for master server 127.0.0.1:3308
ghost            | 2021-04-01 17:33:57 DEBUG Streamer binlog coordinates: bin.000003:78817
ghost            | 2021-04-01 17:33:57 INFO Connecting binlog streamer at bin.000003:78817
ghost            | [2021/04/01 17:33:57] [info] binlogsyncer.go:723 rotate to (bin.000003, 78817)
ghost            | 2021-04-01 17:33:57 DEBUG Beginning streaming
ghost            | 2021-04-01 17:33:57 INFO rotate to next log from bin.000003:0 to bin.000003
ghost            | 2021-04-01 17:33:57 INFO connection validated on 127.0.0.1:3307
ghost            | 2021-04-01 17:33:57 INFO connection validated on 127.0.0.1:3307
ghost            | 2021-04-01 17:33:57 INFO will use time_zone='SYSTEM' on applier
ghost            | 2021-04-01 17:33:57 INFO Examining table structure on applier
ghost            | 2021-04-01 17:33:57 INFO Applier initiated on vanila-minikube:3307, version 5.7.33-log
ghost            | 2021-04-01 17:33:57 INFO Dropping table `try_osc`.`_referringitem_ghc`
ghost            | 2021-04-01 17:33:57 INFO Table dropped
ghost            | 2021-04-01 17:33:57 INFO Creating changelog table `try_osc`.`_referringitem_ghc`
ghost            | 2021-04-01 17:33:57 INFO Changelog table created
ghost            | 2021-04-01 17:33:57 INFO Creating ghost table `try_osc`.`_referringitem_gho`
ghost            | 2021-04-01 17:33:57 INFO Ghost table created
ghost            | 2021-04-01 17:33:57 INFO Altering ghost table `try_osc`.`_referringitem_gho`
ghost            | 2021-04-01 17:33:57 DEBUG ALTER statement: alter /* gh-ost */ table `try_osc`.`_referringitem_gho` MODIFY baseitem_id bigint(20) NOT NULL;
ghost            | 2021-04-01 17:33:57 INFO Ghost table altered
ghost            | 2021-04-01 17:33:57 INFO Waiting for ghost table to be migrated. Current lag is 0s
ghost            | 2021-04-01 17:33:57 INFO Intercepted changelog state GhostTableMigrated
ghost            | 2021-04-01 17:33:57 INFO Handled changelog state GhostTableMigrated
ghost            | 2021-04-01 17:33:57 DEBUG ghost table migrated
ghost            | 2021-04-01 17:33:57 DEBUG Potential unique keys in _referringitem_gho: [PRIMARY (auto_increment): [id]; has nullable: false]
ghost            | 2021-04-01 17:33:57 INFO Chosen shared unique key is PRIMARY
ghost            | 2021-04-01 17:33:57 INFO Shared columns are baseitem_id,id,created_on,updated_on,some_id
ghost            | 2021-04-01 17:33:57 INFO Listening on unix socket file: /tmp/gh-ost.try_osc.referringitem.sock
ghost            | 2021-04-01 17:33:57 DEBUG Reading migration range according to key: PRIMARY
ghost            | 2021-04-01 17:33:57 INFO Migration min values: [1]
ghost            | 2021-04-01 17:33:57 DEBUG Reading migration range according to key: PRIMARY
ghost            | 2021-04-01 17:33:57 INFO Migration max values: [25]
ghost            | 2021-04-01 17:33:57 INFO Waiting for first throttle metrics to be collected
ghost            | 2021-04-01 17:33:57 INFO First throttle metrics collected
ghost            | 2021-04-01 17:33:57 DEBUG Operating until row copy is complete
ghost            | 2021-04-01 17:33:57 DEBUG Getting nothing in the write queue. Sleeping...
ghost            | # Migrating `try_osc`.`referringitem`; Ghost table is `try_osc`.`_referringitem_gho`
ghost            | # Migrating vanila-minikube:3307; inspecting vanila-minikube:3308; executing on vanila-minikube
ghost            | # Migration started at Thu Apr 01 17:33:57 +0000 2021
ghost            | # chunk-size: 1000; max-lag-millis: 1500ms; dml-batch-size: 10; max-load: ; critical-load: ; nice-ratio: 0.000000
ghost            | # throttle-additional-flag-file: /tmp/gh-ost.throttle
ghost            | # Serving on unix socket: /tmp/gh-ost.try_osc.referringitem.sock
ghost            | Copy: 0/25 0.0%; Applied: 0; Backlog: 0/1000; Time: 0s(total), 0s(copy); streamer: bin.000003:80931; Lag: 0.03s, State: migrating; ETA: N/A
ghost            | 2021-04-01 17:33:58 DEBUG Issued INSERT on range: [1]..[25]; iteration: 0; chunk-size: 1000
ghost            | Copy: 0/25 0.0%; Applied: 0; Backlog: 0/1000; Time: 1s(total), 1s(copy); streamer: bin.000003:85152; Lag: 0.03s, State: migrating; ETA: N/A
ghost            | 2021-04-01 17:33:58 DEBUG Iteration complete: no further range to iterate
ghost            | 2021-04-01 17:33:58 DEBUG Getting nothing in the write queue. Sleeping...
ghost            | 2021-04-01 17:33:58 INFO Row copy complete
ghost            | 2021-04-01 17:33:58 DEBUG checking for cut-over postpone
ghost            | 2021-04-01 17:33:58 DEBUG checking for cut-over postpone: complete
ghost            | Copy: 25/25 100.0%; Applied: 0; Backlog: 0/1000; Time: 1s(total), 1s(copy); streamer: bin.000003:85152; Lag: 0.03s, State: migrating; ETA: due
ghost            | 2021-04-01 17:33:58 INFO Grabbing voluntary lock: gh-ost.93.lock
ghost            | 2021-04-01 17:33:58 INFO Setting LOCK timeout as 6 seconds
ghost            | 2021-04-01 17:33:58 INFO Looking for magic cut-over table
ghost            | 2021-04-01 17:33:58 INFO Creating magic cut-over table `try_osc`.`_referringitem_del`
ghost            | 2021-04-01 17:33:58 INFO Magic cut-over table created
ghost            | 2021-04-01 17:33:58 INFO Locking `try_osc`.`referringitem`, `try_osc`.`_referringitem_del`
ghost            | 2021-04-01 17:33:58 INFO Tables locked
ghost            | 2021-04-01 17:33:58 INFO Session locking original & magic tables is 93
ghost            | 2021-04-01 17:33:58 INFO Writing changelog state: AllEventsUpToLockProcessed:1617298438725936820
ghost            | 2021-04-01 17:33:58 INFO Waiting for events up to lock
ghost            | 2021-04-01 17:33:58 INFO Intercepted changelog state AllEventsUpToLockProcessed
ghost            | 2021-04-01 17:33:58 INFO Handled changelog state AllEventsUpToLockProcessed
ghost            | Copy: 25/25 100.0%; Applied: 0; Backlog: 1/1000; Time: 2s(total), 1s(copy); streamer: bin.000003:91780; Lag: 0.03s, State: migrating; ETA: due
ghost            | 2021-04-01 17:33:59 INFO Waiting for events up to lock: got AllEventsUpToLockProcessed:1617298438725936820
ghost            | 2021-04-01 17:33:59 INFO Done waiting for events up to lock; duration=980.452461ms
ghost            | # Migrating `try_osc`.`referringitem`; Ghost table is `try_osc`.`_referringitem_gho`
ghost            | # Migrating vanila-minikube:3307; inspecting vanila-minikube:3308; executing on vanila-minikube
ghost            | # Migration started at Thu Apr 01 17:33:57 +0000 2021
ghost            | # chunk-size: 1000; max-lag-millis: 1500ms; dml-batch-size: 10; max-load: ; critical-load: ; nice-ratio: 0.000000
ghost            | # throttle-additional-flag-file: /tmp/gh-ost.throttle
ghost            | 2021-04-01 17:33:59 DEBUG Getting nothing in the write queue. Sleeping...
ghost            | # Serving on unix socket: /tmp/gh-ost.try_osc.referringitem.sock
ghost            | Copy: 25/25 100.0%; Applied: 0; Backlog: 0/1000; Time: 2s(total), 1s(copy); streamer: bin.000003:92234; Lag: 0.03s, State: migrating; ETA: due
debug2: channel 0: window 999374 sent adjust 49202
ghost            | 2021-04-01 17:33:59 INFO Setting RENAME timeout as 3 seconds
ghost            | 2021-04-01 17:33:59 INFO Session renaming tables is 91
ghost            | 2021-04-01 17:33:59 INFO Issuing and expecting this to block: rename /* gh-ost */ table `try_osc`.`referringitem` to `try_osc`.`_referringitem_del`, `try_osc`.`_referringitem_gho` to `try_osc`.`referringitem`
ghost            | Copy: 25/25 100.0%; Applied: 0; Backlog: 0/1000; Time: 3s(total), 1s(copy); streamer: bin.000003:96646; Lag: 0.03s, State: migrating; ETA: due
ghost            | 2021-04-01 17:34:00 DEBUG Getting nothing in the write queue. Sleeping...
ghost            | 2021-04-01 17:34:00 INFO Found atomic RENAME to be blocking, as expected. Double checking the lock is still in place (though I don't strictly have to)
ghost            | 2021-04-01 17:34:00 INFO Checking session lock: gh-ost.93.lock
ghost            | 2021-04-01 17:34:00 INFO Connection holding lock on original table still exists
ghost            | 2021-04-01 17:34:00 INFO Will now proceed to drop magic table and unlock tables
ghost            | 2021-04-01 17:34:00 INFO Dropping magic cut-over table
ghost            | 2021-04-01 17:34:00 INFO Releasing lock from `try_osc`.`referringitem`, `try_osc`.`_referringitem_del`
ghost            | 2021-04-01 17:34:00 INFO Tables unlocked
ghost            | 2021-04-01 17:34:00 INFO Tables renamed
ghost            | 2021-04-01 17:34:00 INFO Lock & rename duration: 2.00988116s. During this time, queries on `referringitem` were blocked
ghost            | 2021-04-01 17:34:00 INFO Looking for magic cut-over table
ghost            | [2021/04/01 17:34:00] [info] binlogsyncer.go:164 syncer is closing...
ghost            | 2021-04-01 17:34:00 DEBUG done streaming events
ghost            | 2021-04-01 17:34:00 DEBUG Done streaming
ghost            | 2021-04-01 17:34:00 INFO Closed streamer connection. err=<nil>
ghost            | [2021/04/01 17:34:00] [error] binlogstreamer.go:77 close sync with err: sync is been closing...
ghost            | [2021/04/01 17:34:00] [info] binlogsyncer.go:179 syncer is closed
ghost            | 2021-04-01 17:34:00 INFO Dropping table `try_osc`.`_referringitem_ghc`
ghost            | 2021-04-01 17:34:00 INFO Table dropped
ghost            | 2021-04-01 17:34:00 INFO Am not dropping old table because I want this operation to be as live as possible. If you insist I should do it, please add `--ok-to-drop-table` next time. But I prefer you do not. To drop the old table, issue:
ghost            | 2021-04-01 17:34:00 INFO -- drop table `try_osc`.`_referringitem_del`
ghost            | 2021-04-01 17:34:00 INFO Done migrating `try_osc`.`referringitem`
ghost            | 2021-04-01 17:34:00 INFO Removing socket file: /tmp/gh-ost.try_osc.referringitem.sock
ghost            | 2021-04-01 17:34:00 INFO Tearing down inspector
ghost            | 2021-04-01 17:34:00 INFO Tearing down applier
ghost            | 2021-04-01 17:34:00 DEBUG Tearing down...
ghost            | 2021-04-01 17:34:00 INFO Tearing down streamer
ghost            | 2021-04-01 17:34:00 INFO Tearing down throttler
ghost            | 2021-04-01 17:34:00 DEBUG Tearing down...
ghost            | # Done
ghost            | + gh-ost -verbose -debug -stack -allow-on-master -assume-rbr -host 127.0.0.1 -port 3308 -database try_osc -user repl_user_replica -password toor -assume-master-host 127.0.0.1:3307 -master-user root -master-password toor -table baseitem -alter 'MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;' --execute
ghost            | 2021-04-01 17:34:00 INFO starting gh-ost 1.1.0
ghost            | 2021-04-01 17:34:00 INFO Migrating `try_osc`.`baseitem`
ghost            | 2021-04-01 17:34:00 INFO connection validated on 127.0.0.1:3308
ghost            | 2021-04-01 17:34:00 INFO User has ALL privileges
ghost            | 2021-04-01 17:34:00 INFO binary logs validated on 127.0.0.1:3308
ghost            | 2021-04-01 17:34:00 INFO Inspector initiated on vanila-minikube:3308, version 5.7.33-log
ghost            | 2021-04-01 17:34:00 INFO Table found. Engine=InnoDB
ghost            | 2021-04-01 17:34:00 DEBUG Estimated number of rows via STATUS: 20
ghost            | 2021-04-01 17:34:00 DEBUG Validated no foreign keys exist on table
ghost            | 2021-04-01 17:34:00 DEBUG Validated no triggers exist on table
ghost            | 2021-04-01 17:34:00 INFO Estimated number of rows via EXPLAIN: 20
ghost            | 2021-04-01 17:34:00 DEBUG Potential unique keys in baseitem: [PRIMARY (auto_increment): [id]; has nullable: false]
ghost            | 2021-04-01 17:34:00 INFO Master forced to be 127.0.0.1:3307
ghost            | 2021-04-01 17:34:00 INFO log_slave_updates validated on 127.0.0.1:3308
ghost            | 2021-04-01 17:34:00 INFO connection validated on 127.0.0.1:3308
ghost            | 2021-04-01 17:34:00 DEBUG Streamer binlog coordinates: bin.000003:98192
ghost            | 2021-04-01 17:34:00 INFO Connecting binlog streamer at bin.000003:98192
ghost            | [2021/04/01 17:34:00] [info] binlogsyncer.go:133 create BinlogSyncer with config {99999 mysql 127.0.0.1 3308 repl_user_replica    false false <nil> false UTC true 0 0s 0s 0 false}
ghost            | [2021/04/01 17:34:00] [info] binlogsyncer.go:354 begin to sync binlog from position (bin.000003, 98192)
ghost            | [2021/04/01 17:34:00] [info] binlogsyncer.go:203 register slave for master server 127.0.0.1:3308
ghost            | 2021-04-01 17:34:00 DEBUG Beginning streaming
ghost            | 2021-04-01 17:34:00 INFO rotate to next log from bin.000003:0 to bin.000003
ghost            | [2021/04/01 17:34:00] [info] binlogsyncer.go:723 rotate to (bin.000003, 98192)
ghost            | 2021-04-01 17:34:00 INFO connection validated on 127.0.0.1:3307
ghost            | 2021-04-01 17:34:00 INFO connection validated on 127.0.0.1:3307
ghost            | 2021-04-01 17:34:00 INFO will use time_zone='SYSTEM' on applier
ghost            | 2021-04-01 17:34:00 INFO Examining table structure on applier
ghost            | 2021-04-01 17:34:00 INFO Applier initiated on vanila-minikube:3307, version 5.7.33-log
ghost            | 2021-04-01 17:34:00 INFO Dropping table `try_osc`.`_baseitem_ghc`
ghost            | 2021-04-01 17:34:00 INFO Table dropped
ghost            | 2021-04-01 17:34:00 INFO Creating changelog table `try_osc`.`_baseitem_ghc`
ghost            | 2021-04-01 17:34:00 INFO Changelog table created
ghost            | 2021-04-01 17:34:00 INFO Creating ghost table `try_osc`.`_baseitem_gho`
ghost            | 2021-04-01 17:34:00 INFO Ghost table created
ghost            | 2021-04-01 17:34:00 INFO Altering ghost table `try_osc`.`_baseitem_gho`
ghost            | 2021-04-01 17:34:00 DEBUG ALTER statement: alter /* gh-ost */ table `try_osc`.`_baseitem_gho` MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;
ghost            | 2021-04-01 17:34:00 INFO Ghost table altered
ghost            | 2021-04-01 17:34:00 INFO Waiting for ghost table to be migrated. Current lag is 0s
ghost            | 2021-04-01 17:34:00 INFO Intercepted changelog state GhostTableMigrated
ghost            | 2021-04-01 17:34:00 INFO Handled changelog state GhostTableMigrated
ghost            | 2021-04-01 17:34:00 DEBUG ghost table migrated
ghost            | 2021-04-01 17:34:00 DEBUG Potential unique keys in _baseitem_gho: [PRIMARY (auto_increment): [id]; has nullable: false]
ghost            | 2021-04-01 17:34:00 INFO Chosen shared unique key is PRIMARY
ghost            | 2021-04-01 17:34:00 INFO Shared columns are id,created_on,updated_on,product_id
ghost            | 2021-04-01 17:34:00 INFO Listening on unix socket file: /tmp/gh-ost.try_osc.baseitem.sock
ghost            | 2021-04-01 17:34:00 DEBUG Reading migration range according to key: PRIMARY
ghost            | 2021-04-01 17:34:00 INFO Migration min values: [1]
ghost            | 2021-04-01 17:34:00 DEBUG Reading migration range according to key: PRIMARY
ghost            | 2021-04-01 17:34:00 INFO Migration max values: [20]
ghost            | 2021-04-01 17:34:00 INFO Waiting for first throttle metrics to be collected
ghost            | 2021-04-01 17:34:00 INFO First throttle metrics collected
ghost            | 2021-04-01 17:34:00 DEBUG Operating until row copy is complete
ghost            | 2021-04-01 17:34:00 DEBUG Getting nothing in the write queue. Sleeping...
ghost            | # Migrating `try_osc`.`baseitem`; Ghost table is `try_osc`.`_baseitem_gho`
ghost            | # Migrating vanila-minikube:3307; inspecting vanila-minikube:3308; executing on vanila-minikube
ghost            | # Migration started at Thu Apr 01 17:34:00 +0000 2021
ghost            | # chunk-size: 1000; max-lag-millis: 1500ms; dml-batch-size: 10; max-load: ; critical-load: ; nice-ratio: 0.000000
ghost            | # throttle-additional-flag-file: /tmp/gh-ost.throttle
ghost            | # Serving on unix socket: /tmp/gh-ost.try_osc.baseitem.sock
ghost            | Copy: 0/20 0.0%; Applied: 0; Backlog: 0/1000; Time: 0s(total), 0s(copy); streamer: bin.000003:100272; Lag: 0.02s, State: migrating; ETA: N/A
ghost            | 2021-04-01 17:34:01 DEBUG Issued INSERT on range: [1]..[20]; iteration: 0; chunk-size: 1000
ghost            | Copy: 0/20 0.0%; Applied: 0; Backlog: 0/1000; Time: 1s(total), 1s(copy); streamer: bin.000003:104439; Lag: 0.02s, State: migrating; ETA: N/A
ghost            | 2021-04-01 17:34:01 DEBUG Iteration complete: no further range to iterate
ghost            | 2021-04-01 17:34:01 DEBUG Getting nothing in the write queue. Sleeping...
ghost            | 2021-04-01 17:34:01 INFO Row copy complete
ghost            | Copy: 20/20 100.0%; Applied: 0; Backlog: 0/1000; Time: 1s(total), 1s(copy); streamer: bin.000003:104439; Lag: 0.02s, State: migrating; ETA: due
ghost            | 2021-04-01 17:34:01 DEBUG checking for cut-over postpone
ghost            | 2021-04-01 17:34:01 DEBUG checking for cut-over postpone: complete
ghost            | 2021-04-01 17:34:01 INFO Grabbing voluntary lock: gh-ost.95.lock
ghost            | 2021-04-01 17:34:01 INFO Setting LOCK timeout as 6 seconds
ghost            | 2021-04-01 17:34:01 INFO Looking for magic cut-over table
ghost            | 2021-04-01 17:34:01 INFO Creating magic cut-over table `try_osc`.`_baseitem_del`
ghost            | 2021-04-01 17:34:01 INFO Magic cut-over table created
ghost            | 2021-04-01 17:34:01 INFO Locking `try_osc`.`baseitem`, `try_osc`.`_baseitem_del`
ghost            | 2021-04-01 17:34:01 INFO Tables locked
ghost            | 2021-04-01 17:34:01 INFO Session locking original & magic tables is 95
ghost            | 2021-04-01 17:34:01 INFO Writing changelog state: AllEventsUpToLockProcessed:1617298441925002545
ghost            | 2021-04-01 17:34:01 INFO Waiting for events up to lock
ghost            | 2021-04-01 17:34:01 INFO Intercepted changelog state AllEventsUpToLockProcessed
ghost            | 2021-04-01 17:34:01 INFO Handled changelog state AllEventsUpToLockProcessed
ghost            | Copy: 20/20 100.0%; Applied: 0; Backlog: 1/1000; Time: 2s(total), 1s(copy); streamer: bin.000003:110778; Lag: 0.02s, State: migrating; ETA: due
ghost            | 2021-04-01 17:34:02 DEBUG Getting nothing in the write queue. Sleeping...
ghost            | # Migrating `try_osc`.`baseitem`; Ghost table is `try_osc`.`_baseitem_gho`
ghost            | # Migrating vanila-minikube:3307; inspecting vanila-minikube:3308; executing on vanila-minikube
ghost            | # Migration started at Thu Apr 01 17:34:00 +0000 2021
ghost            | # chunk-size: 1000; max-lag-millis: 1500ms; dml-batch-size: 10; max-load: ; critical-load: ; nice-ratio: 0.000000
ghost            | # throttle-additional-flag-file: /tmp/gh-ost.throttle
ghost            | # Serving on unix socket: /tmp/gh-ost.try_osc.baseitem.sock
ghost            | 2021-04-01 17:34:02 INFO Waiting for events up to lock: got AllEventsUpToLockProcessed:1617298441925002545
ghost            | 2021-04-01 17:34:02 INFO Done waiting for events up to lock; duration=979.492945ms
ghost            | Copy: 20/20 100.0%; Applied: 0; Backlog: 0/1000; Time: 2s(total), 1s(copy); streamer: bin.000003:110778; Lag: 0.02s, State: migrating; ETA: due
ghost            | 2021-04-01 17:34:02 INFO Setting RENAME timeout as 3 seconds
ghost            | 2021-04-01 17:34:02 INFO Session renaming tables is 97
ghost            | 2021-04-01 17:34:02 INFO Issuing and expecting this to block: rename /* gh-ost */ table `try_osc`.`baseitem` to `try_osc`.`_baseitem_del`, `try_osc`.`_baseitem_gho` to `try_osc`.`baseitem`
ghost            | 2021-04-01 17:34:02 INFO Found atomic RENAME to be blocking, as expected. Double checking the lock is still in place (though I don't strictly have to)
ghost            | 2021-04-01 17:34:02 INFO Checking session lock: gh-ost.95.lock
ghost            | 2021-04-01 17:34:02 INFO Connection holding lock on original table still exists
ghost            | 2021-04-01 17:34:02 INFO Will now proceed to drop magic table and unlock tables
ghost            | 2021-04-01 17:34:02 INFO Dropping magic cut-over table
ghost            | 2021-04-01 17:34:02 INFO Releasing lock from `try_osc`.`baseitem`, `try_osc`.`_baseitem_del`
ghost            | 2021-04-01 17:34:02 INFO Tables unlocked
ghost            | 2021-04-01 17:34:02 INFO Tables renamed
ghost            | 2021-04-01 17:34:02 INFO Lock & rename duration: 1.003557289s. During this time, queries on `baseitem` were blocked
ghost            | 2021-04-01 17:34:02 INFO Looking for magic cut-over table
ghost            | [2021/04/01 17:34:02] [info] binlogsyncer.go:164 syncer is closing...
ghost            | 2021-04-01 17:34:02 INFO Closed streamer connection. err=<nil>
ghost            | 2021-04-01 17:34:02 INFO Dropping table `try_osc`.`_baseitem_ghc`
ghost            | 2021-04-01 17:34:02 DEBUG done streaming events
ghost            | 2021-04-01 17:34:02 DEBUG Done streaming
ghost            | [2021/04/01 17:34:02] [error] binlogstreamer.go:77 close sync with err: sync is been closing...
ghost            | [2021/04/01 17:34:02] [info] binlogsyncer.go:179 syncer is closed
ghost            | 2021-04-01 17:34:02 INFO Table dropped
ghost            | 2021-04-01 17:34:02 INFO Am not dropping old table because I want this operation to be as live as possible. If you insist I should do it, please add `--ok-to-drop-table` next time. But I prefer you do not. To drop the old table, issue:
ghost            | 2021-04-01 17:34:02 INFO -- drop table `try_osc`.`_baseitem_del`
ghost            | 2021-04-01 17:34:02 INFO Done migrating `try_osc`.`baseitem`
ghost            | 2021-04-01 17:34:02 INFO Removing socket file: /tmp/gh-ost.try_osc.baseitem.sock
ghost            | 2021-04-01 17:34:02 INFO Tearing down inspector
ghost            | 2021-04-01 17:34:02 INFO Tearing down applier
ghost            | 2021-04-01 17:34:02 DEBUG Tearing down...
ghost            | # Done
ghost            | 2021-04-01 17:34:02 INFO Tearing down streamer
ghost            | 2021-04-01 17:34:02 INFO Tearing down throttler
ghost            | 2021-04-01 17:34:02 DEBUG Tearing down...
ghost exited with code 0
```
