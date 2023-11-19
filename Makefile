GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test ./...
SERVER=cmd/server
CLIENT=cmd/client

generate:
	go generate ./...

client:
	cd $(CLIENT) && $(GOBUILD) -o main
	cd $(CLIENT) && ./main

server:
	cd $(SERVER) && $(GOBUILD) -o main
	cd $(SERVER) && ./main

run-server: generate server

run-client: client