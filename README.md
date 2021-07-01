# bukumanga-api

## How to use

### Develop

#### Connect Web

```sh
docker-compose exec web bash
```

#### Connect DB

```sh
docker-compose exec db bash
psql -U pguser bukumanga
```

#### Debug Log

```sh
docker-compose logs -f
```

#### Migration

```sh
docker-compose exec web bash

# create migration
migrate create -ext sql -dir db/migrations -seq {migration_name}

# migrate
migrate -database ${POSTGRESQL_URL} -path db/migrations up N

# rollback
migrate -database ${POSTGRESQL_URL} -path db/migrations down N

# fix dirty 
migrate -database ${POSTGRESQL_URL} -path db/migrations force N
```

#### Seed

```sh
docker-compose exec web bash
psql -h db -U pguser bukumanga -f ./db/seeds/sites.sql
```
