# bukumanga-api

## How to use

### Deploy

```sh
heroku login
git push heroku

heroku container:login
heroku container:push web
heroku container:release web
```

### Migration

```sh
docker-compose exec web bash
migrate create -ext sql -dir db/migrations -seq create_entries_table
```
