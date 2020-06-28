run:
	go run main.go

test:
	go test

build:
	docker build -t samdickinson/proofviewer:latest .

.PHONY: run build test