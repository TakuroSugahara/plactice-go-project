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

### generate code by schema.graphql

1. Write `graphql/schema.graphql`.
2. Execute the following command

```sh
cd backend
go run github.com/99designs/gqlgen
```

#### No Edit files

- backend/api/graph/exec.go
- backend/view/model/models_gen.go

### install or update dependencies modules

```
go get [module-name] // Install module.
go mod tidy // Delete unused modules.
```
