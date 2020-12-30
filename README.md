# OpenConfig

Prototype of open standard for generating, defining and parsing configuration input.

## Motivation

## Principles

## Quick Start

### Go

#### Prerequisites

* [Go](https://golang.org/)

For installation instructions, see [Go’s Getting Started guide](https://golang.org/doc/install).

* [Protocol buffer compiler](https://developers.google.com/protocol-buffers) (`protoc`), [version 3](https://developers.google.com/protocol-buffers/docs/proto3).

For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).

* Go plugins for the protocol compiler:
    
  1. Install the protocol compiler plugins for Go and OpenConfig using the following commands:

    ```bash 
    export GO111MODULE=on  # Enable module mode
    go get google.golang.org/protobuf/cmd/protoc-gen-go \
      github.com/thanos-io/OpenConfig/protoc-gen-go-openconfig
    ```
     
  2. Update your PATH so that the protoc compiler can find the plugins:

    ```bash 
    export PATH="$PATH:$(go env GOPATH)/bin
    ```

#### Run example

1. Change directory to the quick start example directory:

`cd examples/helloword`

TBD

#### Generate example config

1. Change directory to the quick start example directory:

`cd examples/helloword`

2. From the `examples/helloworld` directory, generate Go code for using and parsing configuration defined in proto file.

```bash
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
helloworld/helloworld.proto
```



$ go run greeter_server/main.go
From a different terminal, compile and execute the client code to see the client output:

$ go run greeter_client/main.go
Greeting: Hello world

## Initial Authors

Bartek Płotka @bwplotka
Frederic Branczyk @brancz