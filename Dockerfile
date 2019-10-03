FROM golang:1.12-alpine

RUN apk update && apk upgrade && apk add --no-cache git

ENV GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -ex \
  && apk add \
    make

RUN mkdir -p gotodo

COPY . ./gotodo

WORKDIR gotodo

RUN go build -o gotodo cmd/main.go

RUN chmod +x ./gotodo 
# ENTRYPOINT ["/.sh"]
