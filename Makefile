# Import .env file and export all make variables as environment variables
include .env
export

DOCKER_ID= mathieux51
REPOSITORY = gotodo
VERSION = $(shell head -1 VERSION)
DOCKER_REGISTRY = https://registry.gitlab.com
BINARY_NAME = main
IMAGE_NAME = $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION)

.PHONY: clean
clean: 
		rm -rf main temp

# Docker
.PHONY: docker-login
docker-login:
		@docker login $(DOCKER_REGISTRY) -u $(DOCKER_ID) -p $(DOCKER_REGISTRY_PWD) 

.PHONY: docker-build
docker-build:
		docker build --tag $(IMAGE_NAME) . 

.PHONY: docker-run
docker-run: 
		docker run --rm -it --name $(REPOSITORY) $(IMAGE_NAME)

.PHONY: docker-push
docker-push:
		docker push $(IMAGE_NAME)

.PHONY: docker-update-version
docker-update-version:
	date '+%Y%m%d.%H%M.%S' > VERSION

.PHONY: docker-update
docker-update: docker-login docker-update-version docker-build docker-push

# Go
.PHONY: go-build
go-build:
		go build -o $(BINARY_NAME) -v cmd/main.go

.PHONY: start
start:
		make go-build
		./$(BINARY_NAME)

# Kubernetes
.PHONY: k8s-create-secret
k8s-create-secret:
		@kubectl create secret docker-registry regcred --docker-server=$(DOCKER_REGISTRY) --docker-username=$(DOCKER_ID) --docker-password=$(DOCKER_REGISTRY_PWD) -o yaml --dry-run > deploy/charts/secret.yaml

.PHONY: k8s-init
k8s-init:
		make k8s-create-secret; \
		kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/master/aio/deploy/recommended/kubernetes-dashboard.yaml; \
		kubectl apply -f deploy/config/tiller-clusterrolebinding.yaml; \
		helm --service-account tiller --upgrade; \

.PHONY: install
install:
		helm install deploy/charts

.PHONY: token
token:
		kubectl -n kube-system describe secret default