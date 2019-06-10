DOCKER_ID = mathieux51
REPOSITORY = gotodo
VERSION = $(shell head -1 VERSION)
DOCKER_REGISTRY = cloud.canister.io:5000

.PHONY: clean
clean: 
		rm -rf main temp

# Docker
.PHONY: docker-tag 
docker-tag:
		docker tag $(DOCKER_ID):$(VERSION) $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION)

.PHONY: docker-build
docker-build:
		docker build --tag $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION) . 

.PHONY: docker-run
docker-run: 
		docker run --rm -it --name $(REPOSITORY) -p 3000:3000 $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION)

.PHONY: docker-push
docker-push:
		docker push $(DOCKER_REGISTRY)/$(DOCKER_ID)/$(REPOSITORY):$(VERSION)

.PHONY: docker-update-version
docker-update-version:
	date '+%Y%m%d.%H%M.%S' > VERSION

.PHONY: docker-update
docker-update: docker-update-version docker-build docker-tag docker-push

# .PHONY: login
# login:
# 	docker login -u $$DOCKER_USERNAME -p $$DOCKER_PASSWORD

# Go
.PHONY: go-build
go-build:
		go build cmd/main.go

.PHONY: go-run
go-run:
		main

.PHONY: start
start: go-build go-run

# Helm
.PHONY: helm-debug
helm-debug:
		 helm install --dry-run --debug ./chart