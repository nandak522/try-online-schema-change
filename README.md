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

Iteration 2 (didn't work)
```sql
mysql> ALTER TABLE try_osc.baseitem DROP CONSTRAINT PRIMARY;
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'CONSTRAINT PRIMARY' at line 1
mysql> alter table try_osc.baseitem modify column id bigint(20) NOT NULL AUTO_INCREMENT;
ERROR 1833 (HY000): Cannot change column 'id': used in a foreign key constraint 'baseitem_id_refs_id_2d6ba49a' of table 'try_osc.referringitem'
alter table try_osc.baseitem add primary key (id);
```

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
