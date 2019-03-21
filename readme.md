# gotodo


## setup

### db
```
CREATE DATABASE gotodo CHARACTER SET utf8mb4;
```

### migration

```
$ brew install golang-migrate

## create migration
migrate create -ext "sql" -dir "./db/migrate" -format "unix" example_create_user

## exec migration
migrate -database 'mysql://root:@tcp(localhost:3306)/gotodo' -path db/migrate up
```
