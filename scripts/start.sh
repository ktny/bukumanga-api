#!/bin/bash
set -euxo pipefail

# wait for postgres
until PGPASSWORD=${POSTGRES_PASSWORD} psql -h ${POSTGRES_HOST} -U ${POSTGRES_USER} ${POSTGRES_DB} -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
>&2 echo "Postgres is up - executing command"

# start app
migrate -database ${POSTGRESQL_URL} -path db/migrations up
./app
