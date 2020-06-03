# container setup

```
docker-compose up
```

## create database

```
echo 'CREATE DATABASE IF NOT EXISTS `test-db_development` DEFAULT CHARACTER SET `utf8mb4`' | mysql -h db -uroot
```

# migration

## create sample table

migration 作るとき
```
migrate create -ext sql -dir db/migrations -seq create_hoges_table
```

```
migrate -source file://db/migrations -database 'mysql://root:@tcp(db:3306)/test-db_development' up
```

# db check
## connect mysql

```
mysql -h db -uroot -t test-db_development
```

# test data

```
insert into hoges (name) values ('name-1');
insert into hoges (name) values ('name-2');
insert into hoges (name) values ('name-3');
insert into hoges (name) values ('name-4');

insert into foos (field1, created_at, updated_at) values ('field1-1', now(), now())
insert into foos (field1, created_at, updated_at) values ('field1-2', now(), now());
insert into foos (field1, created_at, updated_at) values ('field1-3', now(), now());
```

