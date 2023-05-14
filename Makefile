.PHONY: build test run-docker

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=main

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/go-chatgpt-prompt-splitter/main.go

test:
	$(GOTEST) -v ./...

lint:
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
