# try-gh-ost
```sh
docker-compose up mysql-master mysql-replica
```

Load data using `initial_data.sql`

### Normal vanila alter

Iteration 1 (didn't work)
```sql
ALTER TABLE try_gh_ost.baseitem DROP PRIMARY KEY;
ALTER TABLE try_gh_ost.baseitem MODIFY COLUMN id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY;
```

Iteration 2
```sql
ALTER TABLE try_gh_ost.baseitem DROP CONSTRAINT PRIMARY;
alter table try_gh_ost.baseitem modify column id bigint(20) NOT NULL AUTO_INCREMENT;
alter table try_gh_ost.baseitem add primary key (id);
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
ghost            | + gh-ost -verbose -debug -stack -allow-on-master -assume-rbr -host 127.0.0.1 -port 3308 -database try_gh_ost -user repl_user_replica -password toor -assume-master-host 127.0.0.1:3307 -master-user root -master-password toor -table baseitem -alter 'MODIFY id bigint(20) NOT NULL AUTO_INCREMENT;'
ghost            | 2021-02-22 19:09:27 INFO starting gh-ost 1.1.0
ghost            | 2021-02-22 19:09:27 INFO Migrating `try_gh_ost`.`baseitem`
ghost            | 2021-02-22 19:09:27 INFO connection validated on 127.0.0.1:3308
ghost            | 2021-02-22 19:09:27 INFO User has ALL privileges
ghost            | 2021-02-22 19:09:27 INFO binary logs validated on 127.0.0.1:3308
ghost            | 2021-02-22 19:09:27 INFO Inspector initiated on vanila-minikube:3308, version 5.7.33-log
ghost            | 2021-02-22 19:09:27 INFO Table found. Engine=InnoDB
ghost            | 2021-02-22 19:09:27 DEBUG Estimated number of rows via STATUS: 4
ghost            | 2021-02-22 19:09:27 ERROR Found 1 parent-side foreign keys on `try_gh_ost`.`baseitem`. Parent-side foreign keys are not supported. Bailing out
ghost            | 2021-02-22 19:09:27 INFO Tearing down inspector
ghost            | 2021-02-22 19:09:27 FATAL 2021-02-22 19:09:27 ERROR Found 1 parent-side foreign keys on `try_gh_ost`.`baseitem`. Parent-side foreign keys are not supported. Bailing out
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
