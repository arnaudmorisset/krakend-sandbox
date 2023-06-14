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

## Installation

```bash
# Get the repository
git clone https://github.com/arnaudmorisset/krakend-sandbox

# For ASDF users - Install Go (or get it using your favorite package manager)
asdf install

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
```

## Known Issue

### Unable to parse recursive protobuf data structure

As of today, the gateway cannot start due to the `RecursiveDataStructure` we can found in the proto definition file.
KrakenD have a bug making the gateway crash when starting due to an issue in the way it wrap proto data structures, making circular relations unaivailable.
Check the proto file (Recurse Service) and try to start the gateway to see by yourself.
