BINARY_NAME=fizzbuzz
VERSION?=latest

.PHONY: all lint test vendor build docker-build docker-release

lint:
	mkdir -p ./bin
	test -f ./bin/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
	./bin/golangci-lint run ./...

test:
	go test -v -race ./...

coverage:
	GO111MODULE=off go get golang.org/x/tools/cmd/cover
	GO111MODULE=off go get github.com/mattn/goveralls
	go test -v -covermode=count -coverprofile=coverage.out ./...
	goveralls -coverprofile=coverage.out -service=travis-ci

vendor:
	go mod vendor

build:
	mkdir -p ./out
	GO111MODULE=on go build -mod vendor -o out/bin/$(BINARY_NAME) .

docker-build:
	docker build --tag ftahmed/$(BINARY_NAME) .

docker-release:
	docker tag ftahmed/$(BINARY_NAME) ftahmed/$(BINARY_NAME):$(VERSION)
	docker push ftahmed/$(BINARY_NAME):$(VERSION)
