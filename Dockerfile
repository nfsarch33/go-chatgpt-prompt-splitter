# First stage: build the Go binary
FROM golang:1.20.4-alpine3.18

# Install make
RUN apk update && apk add --no-cache make

# Install golangci-lint
RUN apk add --no-cache curl && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Run linter
RUN make lint

# Run unit test
RUN make test

# Build the Go app
RUN go build -o main ./cmd/go-chatgpt-prompt-splitter/main.go

# List the contents of /app
RUN ls /app

# Make the binary executable
RUN chmod +x main

# Expose port for the application
EXPOSE 8080

# Run the executable
CMD ["/app/main"]
