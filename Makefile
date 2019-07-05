# Import .env file and export all make variables as environment variables
include .env
export

DOCKER_IMAGE_VERSION=$(shell head -1 VERSION)
IMAGE_NAME=$(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(DOCKER_IMAGE_VERSION)

.PHONY: clean
clean: 
		rm -rf main temp

# Docker
.PHONY: docker-login
docker-login:
		@echo $(DOCKER_REGISTRY_PWD) | docker login $(DOCKER_REGISTRY) -u $(DOCKER_ID) --password-stdin

.PHONY: docker-build
docker-build:
		docker build --tag $(IMAGE_NAME) . 

.PHONY: docker-run
docker-run: 
		docker run --rm -it --name $(REPOSITORY) $(IMAGE_NAME)

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

# Go
.PHONY: coverage
coverage:
		mkdir -p temp; \
		go test -coverprofile temp/cover.out ./...; \
		go tool cover -html=temp/cover.out; \
		rm -rf temp

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
		@kubectl create secret docker-registry regcred --docker-server=$(DOCKER_REGISTRY) --docker-username=$(DOCKER_ID) --docker-password=$(DOCKER_REGISTRY_PWD) -o yaml --dry-run > deploy/charts/templates/secret.yaml

.PHONY: k8s-init
k8s-init:
		helm reset; \
		make k8s-create-secret; \
		helm init --service-account tiller --history-max 200
		kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml; \
		kubectl apply -f deploy/config/tiller-clusterrolebinding.yaml; \
		helm version

.PHONY: install
install:
		helm install --name $(RELEASE_NAME) \
		--set imageName=$(IMAGE_NAME) \
		deploy/charts 
# --dry-run --debug \
		
.PHONY: del
del:
		helm del --purge $(RELEASE_NAME) 

.PHONY: token
token:
		kubectl -n kube-system describe secret default