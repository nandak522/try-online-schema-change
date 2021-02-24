
-- load the below data on master after the migration
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 2);
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 3);
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 1);
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 4);
insert into try_osc.baseitem(created_on, updated_on, product_id) values(now(), now(), 5);

-- load the below data on master after the migration
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 2);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 3);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 1);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 1);
insert into try_osc.referringitem(created_on, updated_on, some_id, baseitem_id) values(now(), now(), 10, 5);
