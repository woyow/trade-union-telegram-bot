# Migrations for mongodb database

## Tool
https://github.com/golang-migrate/migrate

## Install tool
go install github.com/golang-migrate/migrate/v4

## Create up&down migration
```bash
migrate create -ext json -dir ./db/mongodb/migrations migration_name
```