# KrakenD Sandbox

This repository is my sandbox to play with KrakenD in the context of an HTTP to gRPC gateway.
This project depends on Go, Protobuf, and Docker.

> Look at the [official playground](https://github.com/krakendio/playground-enterprise) if you're looking for a complete exposition of all KrakenD features.

## Features

This repository expose the following features from KrakenD:

- REST to gRPC transcoding and proxying
- Key-based authorization
- Backend For Frontend aggregating values from two different services
- Input mapping from URL to gRPC body
- OpenAPI doc generation with SwaggerUI served by KrakenD

## Installation

```bash
# Get the repository
git clone https://github.com/arnaudmorisset/krakend-sandbox

# For ASDF users - Install Go (or get it using your favorite package manager)
asdf install

# Install Protoc and the relevant Go plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"

# Compile proto files
make proto

# Build Go service Docker image
make build

# Start the whole environnement
make start

# Play with different endpoints
make grpcurl-greeter
make grpcurl-user
make grpcurl-profile
make curl-users

# Raise an error at the moment, looks like a bug on KrakenD side.
make curl-recurse
```
