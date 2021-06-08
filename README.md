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
migrate -database ${POSTGRESQL_URL} -path db/migrations up

# rollback
migrate -database ${POSTGRESQL_URL} -path db/migrations down

# fix dirty 
migrate -database ${POSTGRESQL_URL} -path db/migrations force N
```

### Production

#### SSH App

```sh
heroku git:remote --app bukumanga-api
heroku run bash
```

#### Debug Log

```sh
heroku logs -t 
```

#### Connect DB

```sh
heroku pg:psql postgresql-sinuous-85818 --app bukumanga-api
```
