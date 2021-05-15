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
FROM ubuntu

WORKDIR /opt

# golang-migrate
# @see https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
RUN apt-get update && apt-get upgrade -y &&\
    apt-get install -y curl gnupg2 vim &&\
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz &&\
    mv ./migrate.linux-amd64 /usr/bin/migrate

COPY --from=builder /go/bukumanga-api/app /opt/app
COPY ./start.sh /opt/start.sh

EXPOSE 5000

CMD ["./start.sh"]
