# Database

`ranked-pick-api` uses a SQLite3 database for storage.

## Console

To access the local database:
```bash
export DB_URL="sqlite3.db"
sqlite3 $DB_URL
```

## Migrations

[golang-migrate](https://github.com/golang-migrate/migrate/tree/master) is used for migration management.

[Command installation](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
```bash
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

To create a migration:
```bash
migrate create -ext sql -dir migrations -seq sample_migration_name
```

To run up migrations:
```bash
migrate -database sqlite3://${DB_URL} -path db/migrations up
```

To roll back a single migration:
```bash
migrate -database sqlite3://${DB_URL} -path db/migrations down 1
```

To drop the db:
```bash
migrate -database sqlite3://${DB_URL} -path db/migrations drop
```

After a migration error is encountered, the DB is marked dirty and a migration version must be forced before any more migrations can be run:
```bash
migrate -database sqlite3://${DB_URL} -path db/migrations force <DB VERSION BEFORE ERROR>
```
