# Developer: Romanshk Volkov - https://github.com/RomanshkVolkov
# Customer: Dwit México - https://dwitmexico.com
# Project: server-storage

FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GO111MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

RUN apt-get update && apt-get install -y libwebp-dev
RUN go build -o /server-storage ./cmd/

# drop build enviroment
FROM alpine:latest

WORKDIR /srv

RUN apk update && apk add --no-cache libc6-compat libwebp
COPY --from=builder /server-storage .
COPY ./files ./files

EXPOSE 8080

CMD ["/srv/server-storage"]
