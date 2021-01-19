## DB

- postgres

## migration

- flyway

1. Add some sql file to `migrate/migrations`. (ex: `V3__create_table.sql`)
2. Execute the following command.

```sh
docker-compose run --rm db-migrator
```

## backend

- go

```sh
docker-compose up
docker-compose exec api /bin/bash
go run main.go
```

### generate code

1. Add any file to `backend/view/model/`. (ex: `food.go`)
2. Execute the following command

```sh
cd backend
go generate ./...
```
