# gotodo

Simple todo list app. Stack:

- Backend: Golang and Redis
- Frontend: Sapper (Svelte)

## Getting started

```sh
make start
```

## Docker

```
make docker-update
```

This command will ask for loging, bump the version, build the new docker image and push it to the private repo.

## Kubernetes

```
# If this command fails with `Error: could not find tiller`
kubectl -n kube-system delete deployment tiller-deploy
kubectl -n kube-system delete service/tiller-deploy
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
- [x] Replace Minikube with docker-for-desktop
- [x] Replace cloud.canister.io with registry.gitlab.com
- [x] Command for

```
kubectl create secret generic regcred \ --from-file=.dockerconfigjson=<path/to/.docker/config.json> \ --type=kubernetes.io/dockerconfigjson
```

- [ ] Start using Helm values and `_helpers.tpl` because tag ":latest" doesn't work
- [ ] Change to docker login `--password-stdi`
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
