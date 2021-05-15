FROM golang:1.15 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/bukumanga-api

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/bukumanga-api/app /app

EXPOSE 5000

CMD ["./app"]
