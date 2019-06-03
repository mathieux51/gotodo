# gotodo

## Getting started

```sh
make start
```

# To do list

- [x] v2 with [redis](https://godoc.org/github.com/gomodule/redigo/redis)
- [ ] [`docker-compose` vs `Kompose`](https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/#install-kompose)
- [ ] Create docker image
- [ ] Add docker image to [registry](https://cloud.canister.io)
- [ ] Init Helm in minikube
- [ ] Create Chart
- [ ] Check CRUD locally
- [ ] Create helm release
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
