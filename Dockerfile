FROM golang:1.8
# For now we copy everything. 
# For production use on the binary should be copied
COPY . .
WORKDIR /go/src/github.com/mathieux51/gotodo
RUN go get github.com/gomodule/redigo/redis
RUN go get github.com/gorilla/mux
EXPOSE 3001:3001
# Open a new port for debugging purposes
# EXPOSE 3002:3002
RUN make go-build
CMD ["./main"]