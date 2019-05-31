# gotodo

## Getting started

```sh
go build cmd/main.go
./main
```

# To do list

- [ ] v2 with [redis](https://godoc.org/github.com/gomodule/redigo/redis)
- [ ] `docker-compose`
- [ ] Kubernetes local
- [ ] Kubernetes remote
- [ ] Update memory profile commands

## Coverage report

```sh
go test -coverprofile temp/cover.out ./... && go tool cover -html=temp/cover.out
```

## Memory profile

```sh
go test -memprofilerate 1 -memprofile temp/mem.out ./model
# Get model from ...
go tool pprof -web temp/model.a temp/mem.out
```
