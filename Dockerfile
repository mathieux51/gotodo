FROM golang:1.8
WORKDIR /go/src/github.com/mathieux51/gotodo
COPY . .
RUN go get github.com/gomodule/redigo/redis
RUN go get github.com/gorilla/mux
RUN make build
CMD ["./main"]