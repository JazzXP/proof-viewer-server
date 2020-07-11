run:
	go run server.go

test:
	go test

graphql:
	go run github.com/99designs/gqlgen generate

build:
	docker build -t samdickinson/proof-viewer-server:latest .

.PHONY: run build test graphql