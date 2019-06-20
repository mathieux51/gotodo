FROM golang:1.8
WORKDIR /go/src/github.com/mathieux51/gotodo
# For now we copy everything. 
# For production use on the binary should be copied
COPY . .
RUN go get github.com/gomodule/redigo/redis
RUN go get github.com/gorilla/mux
EXPOSE 3001:3001
# EXPOSE 6379:6379
RUN make go-build
CMD ["./main"]