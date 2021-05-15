#!/bin/bash
migrate -database ${POSTGRESQL_URL} -path db/migrations up
./app
