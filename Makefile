.DEFAULT_GOAL := default

default: build

build: build-example build-gen

build-example: cmd/example/main.go
	@go build ./cmd/example

build-gen: cmd/gen/main.go
	@go build ./cmd/gen

generate-models:
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@swagger generate model -f schema/swagger.json

generate-client:
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@swagger generate client -f schema/swagger.json

generate-methods: build-gen
	@$(PWD)/gen > rpc_subscriptions.go
	goimports -w rpc_subscriptions.go
	@gofmt -w rpc_subscriptions.go

.PHONY: generate-models generate-client generate-methods