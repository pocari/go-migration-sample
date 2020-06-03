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

