ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}

# Set a destination for COPY
WORKDIR /app

# Upgrade base image for security
RUN apt-get update && apt-get upgrade -y

# Download Go modules if there are any
COPY go.mod ./

# Install Go modules
RUN go mod download

# Environment variables
ENV GOPATH /golang-server/routes

# Application
COPY main.go ./
COPY routes/ routes/

# Build the binary
RUN go build -o /golang-server

# Expose port to host
EXPOSE 3333

# Run application
CMD [ "/golang-server" ]