-include .env
export
#
# docker
DOCKER_ID ?= 
DOCKER_REPOSITORY ?=
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
VERSION = $(shell head -1 VERSION)
IMAGE_NAME=$(DOCKER_REGISTRY)/$(DOCKER_ID)/$(APP_NAME)

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
# helm init --service-account tiller --history-max 200 --upgrade --wait
.PHONY: init-cluster
init-cluster:
		kubectl create serviceaccount tiller --namespace kube-system
		kubectl create clusterrolebinding tiller-admin-binding --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
		helm init --service-account=tiller --history-max=200 --wait

.PHONY: reset-tiller
reset-tiller:
		kubectl -n kube-system delete deployment tiller-deploy
		kubectl -n kube-system delete service/tiller-deploy

.PHONY: create-secret-docker-registry 
create-secret-docker-registry:
		kubectl create secret docker-registry registrycredentials --docker-server=$(DOCKER_REGISTRY) --docker-username=$(DOCKER_ID) --docker-password=$(DOCKER_REGISTRY_PWD) --docker-email=$(DOCKER_EMAIL)

.PHONY: gcloud-get-credentials
gcloud-get-credentials:
		echo $(GCLOUD_SERVICE_KEY) | gcloud auth activate-service-account --key-file=-
		gcloud container clusters get-credentials $(CLUSTER_NAME) --zone $(GOOGLE_COMPUTE_ZONE)  --project $(GOOGLE_PROJECT_ID) 
	 
# Maybe it's possible to have some kind of a loop here 
# with a comma separated list
.PHONY: helm-install
helm-install:
		@helm install --name $(RELEASE_NAME) \
		--set IMAGE_NAME=$(IMAGE_NAME) \
		--set VERSION=$(VERSION) \
		--set IMAGE_PORT=$(IMAGE_PORT) \
		--set REDIS_IMAGE=$(REDIS_IMAGE) \
		--set REDIS_PORT=$(REDIS_PORT) \
		--set APP_NAME=$(APP_NAME) \
		--set CHART_NAME=$(CHART_NAME) \
		--set CHART_DESCRIPTION=$(CHART_DESCRIPTION) \
		deploy/charts 
		
.PHONY: del
del:
		helm del --purge $(RELEASE_NAME) 

.PHONY: token
token:
		kubectl -n kube-system describe secret default

.PHONY: dashboard
dashboard:
	export POD_NAME=$(kubectl get pods -n default -l "app=kubernetes-dashboard,release=foppish-wasp" -o jsonpath="{.items[0].metadata.name}"); \
	echo https://127.0.0.1:8443/; \
	kubectl -n default port-forward $(POD_NAME) 8443:8443

# Docker
.PHONY: docker-login
docker-login:
		@echo $(DOCKER_REGISTRY_PWD) | docker login $(DOCKER_REGISTRY) -u $(DOCKER_ID) --password-stdin

.PHONY: docker-build
docker-build:
		docker build --tag $(IMAGE_NAME) . 

.PHONY: docker-run
docker-run: 
		docker run --rm -it --name $(DOCKER_REPOSITORY) $(IMAGE_NAME):$(VERSION)

.PHONY: docker-push
docker-push:
		docker push $(IMAGE_NAME):$(VERSION)

.PHONY: docker-pull
docker-pull:
		docker pull $(IMAGE_NAME):$(VERSION)

.PHONY: docker-tag
docker-tag:
	@echo '$(shell docker inspect --format='{{index .Id}}' $(IMAGE_NAME) | awk '{split($$0, a, ":"); print a[2]}')' > VERSION; \
	docker tag $(IMAGE_NAME):latest '$(IMAGE_NAME):$(shell head -1 VERSION)'

.PHONY: docker-update
docker-update: docker-build docker-tag docker-push

