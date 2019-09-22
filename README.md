# gotodo

**Simple** todo list app. Stack:

- Backend: Golang and Redis
- Frontend: Sapper (Svelte)

## Getting started

```sh
touch .env
make start
```

## Docker

```
make docker-update
```

This command will ask for loging, bump the version, build the new docker image and push it to the private repo.

## Kubernetes

```
# If `Error: could not find tiller` run:
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

- [x] Find to way to get VERSION to helm Values because tag ":latest" doesn't work
- [x] Change to docker login `--password-stdi`
- [x] Add project to CircleCI
- [x] Read about CircleCI and Kubernetes
- [x] Do we have to use a cloud provider to use CircleCI with Kubernetes?
- [x] Create gcp cluster
- [x] Make sure the db (Redis) tests are passing
- [x] Remove everything from `charts/values`, secret values should be in `.env` and the rest in the Makefile
- [x] Change selectors
- [x] Change docker image registry. Binary will be copied by CirclCI
- [x] Connect to `gcp-cluster` by edit `~/.kube/config` file
- [x] Setup k8s loadbalancer in helm chart and run `curl localhost:3001/todos`
- [x] Run pod `kubectl exec -it gotodo-b6487675f-xtg2q /bin/sh` with `command: ["/bin/sh"]`, `tty: true` and `stdin: true`
- [ ] Is `Chart.yaml` required?
- [x] Add Jenkins and/or [CircleCI](https://circleci.com/pricing/#build-linux)
- [ ] Fix CHART_NAME
- [ ] Inject gcp variables into CircleCi
- [ ] Deploy on GCP
- [x] Run all tests (including the one with Redis) on CircleCI for a specific branch
- [x] Replace NodePort by LoadBalancer
- [ ] Finish the two courses on Pluralsight
- [ ] Deploy manually gotodo with helm
- [ ] In CircleCI dev branch should be deploy if tests pass, deploy to master should need a manuel approval
- [ ] Read more about Ingress
- [ ] Replace LoadBalancer with Ingress
- [ ] Debug go code inside docker container
- [x] Add tests
- [ ] Write some doc about it
- [ ] Find equivalent to package.json/requirements.txt for go
- [ ] Check auth
- [ ] Oauth with google
- [ ] session in Redis
- [ ] Update memory profile commands
- [ ] Add simple gotodo frontend
- [ ] Organise backend, frontend code and devops code
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

## Shut idle pods down 

[source](https://carlosbecker.com/posts/k8s-sandbox-costs)

```sh
#!/bin/bash
set -eo pipefail

ingress="$(kubectl get pods --output=jsonpath='{.items[*].metadata.name}' |
  xargs -n1 | grep "ingress-nginx" | head -n1)"

# cache all hosts that pass through the ingress
hosts="$(kubectl get ingress nginx-ingress \
  --output=jsonpath='{.spec.rules[*].host}' | xargs -n1)"

# cache pods
pods="$(kubectl get pods)"

# cache ingress logs of the last 90min
logs="$(kubectl logs --since=90m "$ingress")"

# iterate over all deployments
kubectl get deployment --output=jsonpath='{.items[*].metadata.name}' |
  xargs -n1 |
  while read -r svc; do

    # skip svc that don't have pods running
    echo "$pods" | grep -q "$svc" || {
      echo "$svc: no pods running"
      continue
    }

    # skip svcs that don't pass through the ingress
    echo "$hosts" | grep -q "$svc" ||  {
      echo "$svc: not passing through ingress"
      continue
    }

    # skip svcs with pods running less than 1h
    echo "$pods" | grep "$svc" | awk '{print $5}' | grep -q h ||  {
      echo "$svc: pod running less than 1h"
      continue
    }

    # check if any traffic to that svc was made through the ingress in the
    # last hour, scale it down case none
    echo "$logs" | grep -q "default-$svc" || {
      echo "$svc: scaling down"
      kubectl scale deployments "$svc" --replicas 0 --record || true
    }
  done
  ```
