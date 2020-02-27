# gotodo

**Simple** todo list app. Stack:

- Backend: Golang and Redis
- Frontend: Svelte with Sapper (not built yet)
- Platform: auto deployements with CircleCI, Vault, Kubernetes and Docker.

## Getting started

### Requirements

```
brew install vault
```

### Vault, Github and Circle

To get started we need to setup Vault and CircleCI (CCI). CCI is using Vault so
we need to configure it. For now with CCI we are using a project token but we
might change to a personal token approach in the future. Let's first configure
the Vault. Go to [https://github.com/settings/tokens/new](https://github.com/settings/tokens/new) and create a vault token:

```
 - [] admin:org Full control of orgs and teams, read and write org projects
    - [] write:org Read and write org and team membership, read and write org projects
    - [x] read:org Read org and team membership, read org projects
```

Copy your token and login to Vault with your token by running

```
make vaut-login
```


## CircleCI

```
curl \
  --header "Content-Type: application/json" \
  --data '{"build_parameters": {"param1": "value1", "param2": 500}}' \
  --request POST \
  https://circleci.com/api/v1.1/project/github/circleci/mongofinil/tree/master?circle-token=$CIRCLE_TOKEN
```

## Docker

```
make docker-update
```

This command will ask for loging, bump the version, build the new docker image and push it to the private repo.

## Kubernetes

# WIP

## Done 
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
- [x] Is `Chart.yaml` required?
- [x] Add Jenkins and/or [CircleCI](https://circleci.com/pricing/#build-linux)
- [x] Fix CHART_NAME
- [x] Inject gcp variables into CircleCi
- [x] Deploy on GCP
- [x] Run all tests (including the one with Redis) on CircleCI for a specific branch
- [x] Replace NodePort by LoadBalancer
- [x] Finish the two courses on Pluralsight
- [x] Deploy manually gotodo with helm
- [x] Read more about Ingress
- [x] Add tests
- [x] Write some doc about it
- [x] Find equivalent to package.json/requirements.txt for go
- [x] Check auth

## To do

### General

- [ ] Update memory profile commands
- [ ] Organise backend, frontend code and devops code

### Backend

- [ ] session in Redis

### Frontend

- [ ] Add simple gotodo frontend
- [ ] Update frontend with [tailwindcss](https://tailwindcss.com/docs/controlling-file-size/#app)

### DevOps

- [ ] Replace LoadBalancer with Ingress
- [ ] Debug go code inside docker container

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


## Vault

```
$ kubectl exec -it vault-0 sh
/ $ vault operator init -n 1 -t 1
Unseal Key 1: uFaAR6AsAgYL2f0wm9FdmTBzn/y5WpEYIDNwts6z4ro=

Initial Root Token: s.6dFIKedAWGVTwRyWMcwQ0kWa

Vault initialized with 1 key shares and a key threshold of 1. Please securely
distribute the key shares printed above. When the Vault is re-sealed,
restarted, or stopped, you must supply at least 1 of these keys to unseal it
before it can start servicing requests.

Vault does not store the generated master key. Without at least 1 key to
reconstruct the master key, Vault will remain permanently sealed!

It is possible to generate new unseal keys, provided you have a quorum of
existing unseal keys shares. See "vault operator rekey" for more information.
/ $ vault operator unseal uFaAR6AsAgYL2f0wm9FdmTBzn/y5WpEYIDNwts6z4ro=
Key             Value
---             -----
Seal Type       shamir
Initialized     true
Sealed          false
Total Shares    1
Threshold       1
Version         1.3.2
Cluster Name    vault-cluster-9fc2c848
Cluster ID      d32ff548-e6a6-33f5-5b70-d466ea300b2a
HA Enabled      false
```
