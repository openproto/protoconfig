## Go: OpenConfg 1.0 Examples

### Prerequisites

* [Go](https://golang.org/)

For installation instructions, see [Go’s Getting Started guide](https://golang.org/doc/install).

* [Protocol buffer compiler](https://developers.google.com/protocol-buffers) (`protoc`), [version 3](https://developers.google.com/protocol-buffers/docs/proto3).

For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).

* Go plugins for the protocol compiler:

    1. Install the protocol compiler plugins for Go and ProtoConfig using the following commands:

  ```bash
  export GO111MODULE=on  # Enable module mode
  go get google.golang.org/protobuf/cmd/protoc-gen-go \
    github.com/protoconfig/protoconfig/go/protoc-gen-go-protoconfig
  ```

    1. Update your PATH so that the protoc compiler can find the plugins:

  ```bash
  export PATH="$PATH:$(go env GOPATH)/bin
  ```

### Steps

1. Start by reading example application `Configuration Proto Definition` complainant with `ProtoConfig 1.0` (from repo root):  [`proto/examples/helloworld/v1/helloworld.proto`](/proto/examples/helloworld/v1/helloworld.proto)

This file is just `.proto` with a couple of options from few extensions:

* `protoconfig` defined in [`proto/protoconfig/v1/extensions.proto`](/proto/protoconfig/v1/extensions.proto) is an `ProtoConfig 1.0` extension which adds optional context for given data fields like: `required` `hidden` `default` and allows to add metadata that indicates the entry point(s) for the configuration. 
* `kingpin` defined in [`proto/go/kingpinv2/v1/extensions.proto`](/proto/go/kingpinv2/v1/extensions.proto) which adds `Go`, `kingpin` specific options allowing even more Go or [`kingpin`](https://github.com/alecthomas/kingpin) library specific context like custom types (existing file, regexp, IP) etc!

The power of `ProtoConfig 1.0` comes from `protobuf` superpowers: Those options are fully ignored if your `protoc` does not have plugins supporting them (for example you generate data structures for C++ or Java!). This allows ultimate extensibility.

2. This example has already generated Go code from this `helloworld` application `Configuration Proto Definition`, and you can see it [go/examples/helloworld/helloworld.pb.go](/go/examples/helloworld/helloworld.pb.go) and [go/examples/helloworld/helloworld_protoconfig.pb.go](/go/examples/helloworld/helloworld_protoconfig.pb.go). Since it's generated code,
   it's not readable much. Check [go.dev](https://pkg.go.dev/github.com/protoconfig/protoconfig/go/examples) instead! (It's go code after all and supports `godoc`!)
   
What you see is the `protobuf` Go code that allows to marshal and unmarshal filled types in to `ProtoConfig 1.0` (and proto) compliant encoding.
   
3. Change directory to our quick start example directory for Go `cd go/examples/helloworld`
   
4. In this directory, on top of generated Go code for `helloworld` `Configuration Proto Definition` we see a couple of directories:

* `./configurable` is an `ProtoConfig 1.0` compliant `helloword` application that just allows `ProtoConfig` for configuration. 
* `./configurable-kingpinv2` is an `ProtoConfig 1.0` compliant `helloword` application that allows `ProtoConfig` as well as standard POSIX flags (that were manually implemented based on [`kingpin`](https://github.com/alecthomas/kingpin) library ) for configuration.
* `./configurable-kingpinv2-gen` is an `ProtoConfig 1.0` compliant `helloword` application that allows `ProtoConfig` as well as standard POSIX flags that were generated thanks to [`../kingpinv2/`](/go/kingpinv2) extension.
* `./configurator` is `helloworld` application client that configures (by executing it in another process) in couple different ways in order
to get `(really!) Hello my world for "Alt C" in 2077 year!` string output.

5. Feel free to run `./configurable-kingpinv2` with any parameters you want. Check all flags via `go run ./configurable-kingpinv2 --help` Executable will execute all as `Configuration Proto Definition` defines and help specifies. For example you can run:

```bash
go run ./configurable-kingpinv2 hello --world="my" --year=2077 --name="Alt C" --lang=ENGLISH --add-really
```

But you can also use `ProtoConfig 1.0` convention!:

```bash
go run ./configurable-kingpinv2 --protoconfigv1='{"hello": {"name": "Alt C", "year": 2077, "world": "my", "add_really": true}}'
```

All the above should print `(really!) Hello my world for "Alt C" in 2077 year!`

6. Run configurator to check if all applications returns expected message by running: `go run ./configurator`

7. Read [`./configurator`](/go/examples/helloworld/configurator/main.go) to check all different ways to configure `helloworld` `configurable`
with and without `ProtoConfig 1.0`!

#### Updating Configuration Proto Definition

TBD (tl;dr: `make proto`)