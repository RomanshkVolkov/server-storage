FROM golang:latest as base
WORKDIR /srv

# install libweb for chai2010/webbp
RUN apt-get update && apt-get install -y libwebp-dev

RUN go install github.com/air-verse/air@latest

ENV GO111MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64 \
  PKG_CONFIG_PATH=/usr/lib/pkgconfig


