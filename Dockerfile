# Use a base image with Xvfb and Google Chrome installed
FROM selenium/standalone-chrome:latest

# Switch to root user to install packages
USER root

# Install Golang
RUN apt-get update && \
    apt-get -y install wget && \
    wget https://dl.google.com/go/go1.20.linux-amd64.tar.gz && \
    tar -xvf go1.20.linux-amd64.tar.gz && \
    mv go /usr/local

# Set environment variables for Go
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Install make
RUN apt-get update && apt-get -y install make

# Install curl
RUN apt-get update && apt-get -y install curl

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1

# Install dbus and configure it for use with Chrome
RUN apt-get update && apt-get -y install dbus-x11 && \
    mkdir -p /var/run/dbus && chown seluser:seluser /var/run/dbus

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/app

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
RUN ls $GOPATH/src/app

# Make the binary executable
RUN chmod +x main

RUN mkdir -p /home/seluser/.cache/google-chrome/Default/Cache/Cache_Data
RUN chown -R seluser:seluser /home/seluser/.cache

# Switch back to the non-root user
USER seluser

# Expose port for the application
EXPOSE 8080

# Start D-Bus, Xvfb, Chrome, and the Go app
CMD dbus-daemon --system && \
    Xvfb :99 -screen 0 1024x768x16 & \
    export DISPLAY=:99 && \
    google-chrome-stable --disable-gpu --disable-software-rasterizer --no-sandbox --disable-dev-shm-usage & \
    $GOPATH/src/app/main
