DOCKER_ID = mathieux51
REPOSITORY = gotodo
VERSION = $(shell head -1 VERSION)
DOCKER_REGISTRY = cloud.canister.io:5000

.PHONY: clean
clean: 
		rm -rf main temp

# Docker
.PHONY: docker-login
docker-login:
		docker login $(DOCKER_REGISTRY) -u $(DOCKER_ID)

.PHONY: docker-build
docker-build:
		docker build --tag $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION) . 

.PHONY: docker-run
docker-run: 
		docker run --rm -it --name $(REPOSITORY) -p 3001:3001 $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION)

.PHONY: docker-push
docker-push:
		docker push $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION)

.PHONY: docker-update-version
docker-update-version:
	date '+%Y%m%d.%H%M.%S' > VERSION

.PHONY: docker-update
docker-update: docker-login docker-update-version docker-build docker-push

# Go
.PHONY: go-build
go-build:
		go build cmd/main.go

.PHONY: go-run
go-run:
		main

.PHONY: start
start: go-build go-run

# Kubernetes
.PHONY: k8s-init-remote
k8s-init-remote:
		kubectl apply -f deploy/config/tiller-clusterrolebinding.yaml; \
		helm init --service-account tiller; \