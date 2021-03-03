-- create replication
-- on master
CREATE USER repl_user@'%';
GRANT REPLICATION SLAVE ON *.* TO repl_user@'%' IDENTIFIED BY 'toor';
GRANT ALL PRIVILEGES ON *.* TO root@'%' IDENTIFIED BY 'toor' with grant option;FLUSH PRIVILEGES;
FLUSH PRIVILEGES;
show master status \G

-- on slave
stop slave;
CHANGE MASTER TO
    MASTER_HOST='localhost',
    MASTER_PORT=3307,
    MASTER_USER='repl_user',
    MASTER_PASSWORD='toor',
    MASTER_LOG_FILE='bin.000003',
    MASTER_LOG_POS=1247;
start slave;
CREATE USER repl_user_replica@'%';
GRANT ALL PRIVILEGES ON *.* TO repl_user_replica@'%' IDENTIFIED BY 'toor' with grant option;FLUSH PRIVILEGES;
show master status \G


-- create schema on master
CREATE DATABASE try_osc;
use try_osc;
CREATE TABLE try_osc.baseitem (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `updated_on` datetime NOT NULL,
  `product_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- create schema on master
CREATE TABLE try_osc.referringitem (
  `baseitem_id` int(11) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `updated_on` datetime NOT NULL,
  `some_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `referringitem_23d46617` (`some_id`),
  KEY `referringitem_34e005d0` (`baseitem_id`),
  CONSTRAINT `baseitem_id_refs_id_2d6ba49a` FOREIGN KEY (`baseitem_id`) REFERENCES `baseitem` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- list constraints on master
select * from information_schema.TABLE_CONSTRAINTS where table_name = 'baseitem';
select * from information_schema.TABLE_CONSTRAINTS where table_name = 'referringitem';

-- load data on master
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 1);
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 2);
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 1);
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 4);
-- insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 2147483648);

-- load data on master
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 1);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 2);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 2);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 1);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 4);

-- Do the migration, either through pt-online-schema-change or gh-ost

-- Manually fixing the keys
-- ALTER TABLE `try_osc`.`referringitem`
--   DROP FOREIGN KEY `baseitem_id_refs_id_2d6ba49a`,
--   ADD CONSTRAINT `_baseitem_id_refs_id_2d6ba49a` FOREIGN KEY (`baseitem_id`) REFERENCES `try_osc`.`baseitem` (`id`);

-- ALTER TABLE `try_osc`.`referringitem`
--   DROP FOREIGN KEY `baseitem_id_refs_id_2d6ba49a`;

-- ALTER TABLE try_osc.referringitem
--   ADD CONSTRAINT _baseitem_id_refs_id_2d6ba49a FOREIGN KEY  (baseitem_id) REFERENCES try_osc.baseitem (id);
