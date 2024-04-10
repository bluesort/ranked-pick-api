# DB

Access the local development DB:
```bash
psql -d development
```

## Migrations

See [golang-migrate](https://github.com/golang-migrate/migrate/tree/master) for installation and additional documentation.

[Postgres tutorial](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

Create a migration:
```bash
migrate create -ext sql -dir db/migrations -seq sample_migration_name
```

Run up migrations:
```bash
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

To reset the DB, run all down migrations:
```bash
migrate -database ${POSTGRESQL_URL} -path db/migrations down
```

After a migration error is encountered, the DB is marked dirty and a migration version must be forced before any more migrations can be run:
```bash
migrate -database ${POSTGRESQL_URL} -path db/migrations force <DB VERSION BEFORE ERROR>
```
