.PHONY: start
start:
		go build cmd/main.go && ./main

.PHONY: clean
clean: 
		rm -rf main temp 