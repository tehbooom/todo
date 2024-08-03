BINARY := todo
SHELL := /bin/bash -o pipefail

build:
	go build -o bin/$(BINARY)

clean:
	rm -rf bin/$(BINARY)

all: clean build 

lint: deps
	golangci-lint run -v --timeout 10m 

format: deps
	go get
	golangci-lint run --fix

deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)bin v1.57.2
	go install github.com/google/addlicense@latest

test:
	go vet
	go test -coverprofile=coverage.txt -covermode count

