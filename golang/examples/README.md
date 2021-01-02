## Go: OpenConfg 1.0 Examples

### Prerequisites

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

    1. Update your PATH so that the protoc compiler can find the plugins:

  ```bash
  export PATH="$PATH:$(go env GOPATH)/bin
  ```

#### Run Example

1. Read the OpenConfig spec for `examples/helloworld` project (from repo root): `examples/helloworld/helloworld.proto`
2. This example has already generated Go code from this spec, you can read it: `golang/examples/helloworld/hellworld.pb.go` and `golang/examples/helloworld/hellworld_openconfig.pb.go`
3. Change directory to the quick start example directory for Go `cd golang/examples/helloworld`
4. Go example contains two binaries that are already importing and using generated code `./executor` and `./executable`
5. Run `go run ./executable --help` and `go run ./executable hello --help` to see the example executable help output of the available configuration options. Now if you look on Go code for this (`./executable/main.go`) you will that most of the parsing code is actually generated from OpenConfig proto.
6. Feel free to run executable with any parameters you want. Executable will execute all as spec and help specifies. For example you can run:

```bash
go run ./executable hello --world="my" --year=2021 --name="Kate B" --lang=ENGLISH --add-really
```

The above should print ``

1. This is not everything. Spec and this Go binding allows to generate nice executor (Go client of executable). You can see executor using generated executor logic in `./executor/main.go`.
2. Feel free to play with `./executor/main.go` and run it with `go run ./executor`. It will invoke the executable with chosen (typed!) parameters,

#### Updating & Generating example OpenConfig Spec

Let's now try to change the spec and re-generate Go code. This allows our Go executor and executable implementations to use it immediately

1. Modify `examples/helloworld/helloworld.proto`. Let's remove X
2. From the `go` directory run `make proto`.

To see what it does you can check `golang/Makefile` `proto` target:

```Makefile
	@echo ">> generating $(REPO_ROOT_DIR)/examples/helloworld/helloworld.proto in $(REPO_ROOT_DIR)/golang/examples/helloworld/"
	@PATH=$(GOBIN):$(TMP_GOBIN) $(PROTOC) -I $(REPO_ROOT_DIR)/examples/helloworld \
		--go_out=./examples/helloworld/ --go_opt=paths=source_relative \
	   	--go-openconfig_out=./examples/helloworld/ --go-openconfig_opt=paths=source_relative \
	    $(REPO_ROOT_DIR)/examples/helloworld/helloworld.proto
```

This makefile snippet generates the code from proto to our `golang/example/helloworld` directory. It has to have `protoc`, `protoc-gen-go` and `protoc-gen-openconfig` binaries installed (I know that's bit a lot 💩 - complains should go to profobuf ecosystem)

1. Once generated repeat same process with running executable and executor as previously.
