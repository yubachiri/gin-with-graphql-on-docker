FROM golang:1.14-alpine3.11 as build

LABEL maintainer="https://github.com/nekochans"

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

ENV GO111MODULE on

RUN set -eux && \
  go build -o app main.go

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/app .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/app

RUN touch /var/log/mysql/mysqld.log

CMD ["./app"]
