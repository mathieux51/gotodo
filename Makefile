.PHONY: build
build:
		go build cmd/main.go

.PHONY: run
run:
		./main

.PHONY: start
start: build run

.PHONY: clean
clean: 
		rm -rf main temp 

.PHONY: helm_debug
helm_debug:
		 helm install --dry-run --debug ./chart

.PHONY: docker_build
docker_build:
		docker build -t mathieux51/gotodo:0.1.0 .

.PHONY: docker_run
docker_run: 
		docker run --rm -it --name gotodo -p 3000:3000 mathieux51/gotodo:0.1.0