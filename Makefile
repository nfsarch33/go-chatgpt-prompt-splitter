.PHONY: build test run-docker

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=main

all: test build

build:
	go mod tidy
	# go build -o main -v ./cmd/go-chatgpt-prompt-splitter/main.go
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/go-chatgpt-prompt-splitter/main.go

test:
	# go test -v ./...
	$(GOTEST) -v ./...

lint:
	go mod tidy
	golangci-lint run ./...

run:
	./$(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run-docker-build:
	docker-compose --verbose up --build

run-docker:
	docker-compose --verbose up
