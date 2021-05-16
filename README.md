# bukumanga-api

## How to use

### Migration

```sh
docker-compose exec web bash

# create migration
migrate create -ext sql -dir db/migrations -seq create_entries_table

# migrate
migrate -database ${POSTGRESQL_URL} -path db/migrations up

# rollback
migrate -database ${POSTGRESQL_URL} -path db/migrations down
```

### SSH App

```sh
heroku git:remote --app bukumanga-api
heroku run bash
```

### Connect DB

```sh
heroku pg:psql postgresql-sinuous-85818 --app bukumanga-api
```
