FROM golang:1.12-alpine

ENV GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -ex \
  && apk add \
    make

RUN mkdir -p gotodo

COPY . ./gotodo

WORKDIR gotodo

RUN go build -o gotodo cmd/main.go

