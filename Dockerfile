# stage1 builder
FROM golang:1.15 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/bukumanga-api

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

# stage2 final
FROM ubuntu:20.04

WORKDIR /opt

# golang-migrate
# @see https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
RUN apt-get update -y && apt-get upgrade -y &&\
    apt-get install -y curl gnupg2 vim &&\
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz &&\
    mv ./migrate.linux-amd64 /usr/bin/migrate

# postgresql-client
# @see https://www.postgresql.org/download/linux/ubuntu/
RUN echo 'deb http://apt.postgresql.org/pub/repos/apt/ bionic-pgdg main' > /etc/apt/sources.list.d/pgdg.list &&\
    curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - &&\
    apt-get update && apt-get install -y postgresql-client-12

COPY --from=builder /go/bukumanga-api/app /opt/app
COPY ./db /opt/db/
COPY ./scripts/start.sh /opt/

EXPOSE 5000

ENTRYPOINT [ "./start.sh" ]
