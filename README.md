# gotodo

Simple todo list app. Stack:

- Backend: Golang and Redis
- Frontend: Sapper (Svelte)

## Getting started

```sh
make start
```

# WIP

- [x] v2 with [redis](https://godoc.org/github.com/gomodule/redigo/redis)
- [x] [`docker-compose` vs `Kompose`](https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/#install-kompose)
- [x] Create docker image
- [x] Add docker image to [registry](https://cloud.canister.io)
- [x] Init Helm in minikube
- [x] Create Chart
- [x] Check CRUD locally
- [x] Create helm release
- [x] Read about Makefile
- [x] Fix Makefile
- [ ] Deploy on DigitalOcean
- [ ] Debug go code inside docker container
- [ ] Organise backend, frontend code and devops code
- [ ] Find equivalent to package.json/requirements.txt for go
- [ ] Write some doc about it
- [ ] Add simple gotodo frontend
- [ ] Check auth
- [ ] Oauth with google
- [ ] session in Redis
- [ ] Update memory profile commands
- [ ] Update frontend with [tailwindcss](https://tailwindcss.com/docs/controlling-file-size/#app)

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
