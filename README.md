## DB

- postgres

## migration

- flyway

1. Add some sql file to `migrate/migrations` (ex: `V3__create_table.sql`)
2. Execute the following command

```sh
docker-compose run --rm db-migrator
```
