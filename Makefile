-include .env
export
#
# docker
# ID to login to docker
DOCKER_ID ?= 
# name of the docker image in the registry
DOCKER_REPOSITORY ?=
# remote server where the docker image is stored
DOCKER_REGISTRY ?= 
DOCKER_REGISTRY_PWD ?= 
IMAGE_NAME = golang:1.12.7-alpine
IMAGE_PORT = 3001
REDIS_NAME = redis
REDIS_IMAGE = redis:alpine
REDIS_PORT = 6379
# kubernetes
CLUSTER_NAME ?=
GCLOUD_SERVICE_KEY ?=
GOOGLE_COMPUTE_ZONE ?=
GOOGLE_PROJECT_ID ?=
# go
BINARY_NAME = gotodo
RELEASE_NAME = dev
# app
APP_NAME = gotodo
APP_VERSION = 0.0.1

# DOCKER_IMAGE_VERSION=$(shell head -1 VERSION)
# IMAGE_NAME=$(DOCKER_REGISTRY)/$(DOCKER_ID)/$(DOCKER_REPOSITORY):$(DOCKER_IMAGE_VERSION)
IMAGE_NAME=$(DOCKER_REGISTRY)/$(DOCKER_ID)/$(DOCKER_REPOSITORY):latest

.PHONY: clean
clean: 
		rm -rf main temp

# Go
.PHONY: coverage
coverage:
		mkdir -p temp; \
		go test -coverprofile temp/cover.out ./...; \
		go tool cover -html=temp/cover.out; \
		rm -rf temp

.PHONY: test
test:
		go test ./...

.PHONY: get
get:
		go get -v -t -d ./...

.PHONY: build
build:
		go build -o $(BINARY_NAME) -v cmd/main.go

.PHONY: run
run:
		./$(BINARY_NAME)

.PHONY: start
start:
		make build
		make run

# Kubernetes

.PHONY: init-cluster
init-cluster: 
		helm install vault ./deploy/vault-helm-0.4.0	
		helm install kubernetes-dashboard stable/kubernetes-dashboard

.PHONY: reset 
reset:
		helm uninstall kubernetes-dashboard
		helm uninstall vault
		kubectl delete persistentvolumeclaims data-vault-0


.PHONY: create-secret-docker-registry 
create-secret-docker-registry:
		kubectl create secret docker-registry registrycredentials --docker-server=$(DOCKER_REGISTRY) --docker-username=$(DOCKER_ID) --docker-password=$(DOCKER_REGISTRY_PWD) --docker-email=$(DOCKER_EMAIL)

# Maybe it's possible to have some kind of a loop here 
# with a comma separated list
.PHONY: helm-install
helm-install:
		@helm install --name $(RELEASE_NAME) \
		--set IMAGE_NAME=$(IMAGE_NAME) \
		--set IMAGE_PORT=$(IMAGE_PORT) \
		--set REDIS_IMAGE=$(REDIS_IMAGE) \
		--set REDIS_PORT=$(REDIS_PORT) \
		--set APP_NAME=$(APP_NAME) \
		--set APP_VERSION=$(APP_VERSION) \
		--set CHART_NAME=$(CHART_NAME) \
		--set CHART_DESCRIPTION=$(CHART_DESCRIPTION) \
		deploy/charts 
		
.PHONY: del
del:
		helm del --purge $(RELEASE_NAME) 

.PHONY: token
token:
		@sh scripts/token.sh

.PHONY: dashboard
dashboard:
		@sh scripts/dashboard.sh


.PHONY: vault-ui 
vault-ui:
		@kubectl port-forward vault-0 8200:8200

# Docker
.PHONY: docker-login
docker-login:
		@echo $(DOCKER_REGISTRY_PWD) | docker login $(DOCKER_REGISTRY) -u $(DOCKER_ID) --password-stdin

.PHONY: docker-build
docker-build:
		docker build --tag $(IMAGE_NAME) . 

.PHONY: docker-run
docker-run: 
		docker run --rm -it --name $(DOCKER_REPOSITORY) $(IMAGE_NAME)

.PHONY: docker-push
docker-push:
		docker push $(IMAGE_NAME)

.PHONY: docker-pull
docker-pull:
		docker pull $(IMAGE_NAME)

.PHONY: docker-update-version
docker-update-version:
	date '+%Y%m%d.%H%M.%S' > VERSION

.PHONY: docker-update
docker-update: docker-login docker-update-version docker-build docker-push

