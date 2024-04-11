# Migrations for mongodb database

## Tool
https://github.com/golang-migrate/migrate

## Install tool
go get github.com/golang-migrate/migrate
go install github.com/golang-migrate/migrate

## Create up&down migration
```bash
migrate create -ext json -dir ./db/mongodb/migrations migration_name
```