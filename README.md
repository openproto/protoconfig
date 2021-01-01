# OpenConfig

Open standard for generating, defining and using software configuration input.

## Motivation

See ["Configuration in 2021 is still broken"](https://deploy-preview-26--bwplotka.netlify.app/2020/configuring-sw-is-broken/)

## Spec Tl;DR

### Principles

Software `configuration` is a process of passing information to the software in order to change its behavior without recompilation or sometimes even without restarting (dynamic configuration).  

Two main sides exists during software configuration process:

* `Configurable`: Software we want to configure, statically (during startup) or dynamically (on the run)
* `Configurator`: Another software or human user that desire to configure the `Configurable` automatically or manually (human).

The key element of `OpenConfig` specification is a strict format and rules of the `Configuration Definition`. This definition is the true single language that both `Configurator` and `Configurable` can use as shown on below diagram:

<div class="bg-gray">
<img src="https://docs.google.com/drawings/d/e/PACX-1vSANZkljSiDgV-o0a-dL0ryZz19p3Hblt5V_qozhBcY5ILq8j3T2GEAdCCHFHoSGT9h2H4LDqJ9bCn_/pub?w=1440&h=1080"/>
</div>

`OpenConfig 1.0` standard specifies that, in order for a `Configurable` (Software) to be `OpenConfig 1.0` compliant, it shall have one `Configuration Proto Definition`, which shall be the [source of truth](https://en.wikipedia.org/wiki/Single_source_of_truth) of its configuration input. Such a `Configuration Proto Definition` (called further `CPD`) shall meet following requirements:

* Shall be stored and versioned together with the software itself (e.g in the same repository).
* Shall be compatible with [Protocol Buffer Version 3](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec) specification.
* Shall define all the configuration options the `Configurable` exposes.
* Can use [**OpenConfig Proto Extensions Format 1.0.**](proto/openconfig/v1/extensions.proto) which might be useful for more advanced configuration options.
* Can use any other [proto](https://developers.google.com/protocol-buffers) extension or important one or more `CPD` from other software.

Furthermore, additionally to `CPD`, the `OpenConfig 1.0` specifies the `Encoded Configuration Message` (called further `ECM`), which is a configuration information serialized into stream of bytes in [`protobuf "byte" encoding`](https://developers.google.com/protocol-buffers/docs/encoding) or [`json encoding`](https://developers.google.com/protocol-buffers/docs/proto3#json) format based on specific `Configuration Proto Definition (CPD)`.

`OpenConfig 1.0` specifies `Configuration Process` as passing the `Encoded Configuration Message (ECM)` to `Configurable` in order to control programmed options it exposes.
If the `ECM` is passed and parsed on the start of the application before anything else (e.g starting functionalities) it's called `Static Configuration Process (SCP)`. For example reading `ECM` from file on start or reading `ECM` from CLI flag. If `ECM` is passed and configuration is performed during further execution of the `Configurable` execution it's called `Dynamic Configuration Process (DCP)` Watching the file with `ECM` for modify events and reloading internal components when those occur based on new `ECM` input, or accepting `ECM` on HTTP endpoint are examples of `DCP`.

The `Configurable` shall implement either `SCP` or `DCP` or both. `SCP` is recommended for simplicity, explicitness and lower complexity on `Configurable` side reasons.
In any case, the `Configurable` shall accept the `ECM` based on its `CPD`, decode it based on its `CPD` and use parsed data to control programmed options. Note that this specification does not force any consumption mechanism as long as there is one.

Additionally:

* [`json encoding`](https://developers.google.com/protocol-buffers/docs/proto3#json) shall be supported.
* [`protobuf "byte" encoding`](https://developers.google.com/protocol-buffers/docs/encoding) can be supported, but it's recommended to not use it for most of the cases. Reason is that configuration has to be verbose, simple and explicit. Ideally, software is configured statically on the start and all its configuration is visible with a quick glance on CLI invocation. See further rationales [here](https://www.bwplotka.dev/2020/flagarize/#flags-ftw).
* `Configurable` should accept empty or partially empty `ECM` by providing safe and minimalistic defaults.

Overall, `OpenConfig 1.0` recommends allowing `SCP` rather `SPD`, by accepting `json` based `ECM` through the `--openconfig.v1` CLI flag. 

Full standard will be published soon (help wanted if you want to contribute 🤗)

### Why Protocol Buffers?

> Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

On top of incredibly efficient and fast serialization capabilities it also works well [Interface Definition Language](https://en.wikipedia.org/wiki/Interface_description_language) and it is the most used and widely adopted IDL in the industry. For example, protobuf is the definition language and serialization format for [gRPC](https://grpc.io/).

Main reasons are that `proto` (version 3) language is **strongly typed, simple, consistent**, but also **backward and forward compatible**. In addition to it's stable and fast binary serialization format it be also easily decoded or encoded into human readable formats like YAML or JSON.

Last, but the least there exists variety of encoders for almost every coding language. You can read more [here](https://developers.google.com/protocol-buffers/docs/cpptutorial).

If that is not enough, `protobuf` ecosystem is [constantly evolving](https://twitter.com/bwplotka/status/1345076383190556676).

### Scope

`OpenConfig 1.0` standard specifies a process of configuring, define configuration and consume configuration by software applications in a way that is: 

* Discoverable and Type Safe: **We don't need to guess or read complex documentations in order to know executable configuration format, invariants and assumptions**
* Auto Generatable: **We don't need manually write application client and executable config definitions**
* Easy to Verify: **Non-semantic verification and validation (e.g ) does not require executable participation**
* Maintainable: **Smallest possible impact in the event of executable upgrade of modification**
* Backward & Forward Compatible: **Smooth version migrations**

#### Not in the scope

* Runtime APIs, RPCs and Invocation semantics (e.g error codes, output, input).
* Configuration or Invocation API that will accept (e.g dynamically or statically) for configuration in `OpenConfig` format.
* Discoverability of `Configuration Proto Definitions` details.

## Content

| Item   | What | Status |
|--------|------|--------|
| `proto/openconfig/v1/extensions.proto`  |  *OpenConfig Proto Extensions Format 1.0* | Beta |
| `proto/examples`  |  Example *Configuration Proto Definitions (CPD)* | Beta |
| golang |      | WIP |
|        |      |     |


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

  1. Update your PATH so that the protoc compiler can find the plugins:

  ```bash
  export PATH="$PATH:$(go env GOPATH)/bin
  ```

#### Run example

1. Read the OpenConfig spec for helloworld project (from repo root): `examples/helloworld/helloworld.proto`
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

## Contributing

Any help wanted. Do you have idea, want to help, don't know how to start helping? 

Put an issue on this repo, create PR or ping us on the CNCF Slack (`@bwplotka`, `@brancz`)! 🤗

## Initial Authors

* Bartek Płotka @bwplotka
* Frederic Branczyk @brancz

## Open Questions:

### Block protobuf bytes encoding completely? or actually make it mandatory?

This is tricky. From our experience configuration **has to be** human-readable, especially static one. However, there might be cases of configuration being large or maybe configuration process being latency critical. In those cases well compressible proto byte format shall be useful.

### Large `ECM`

Large protobuf are really hard to use. Standard should specify some solution when `Encoded Configuration Message` is larger than Megabytes.

### Standard should specify discoverability?

Discoverability of the protobuf is a big problem in the ecosystem (e.g `gRPC`). Ideally binaries are self-describing. There might be huge value in bringing
this topic here, but it could be also subject of another specification (as per `Focus on one thing well`).

For example:

> `CPD` Shall be embedded in the software executable/binary and available on demand. Note that this specification does not force any delivery mechanism as long as there is one.
>  * It's recommended to expose `openconfig.v1` argument or flag that will output embedded `CPD`  in `.proto` format.
>  * For web applications it's recommended to also expose `/openconfig/v1` endpoint that returns `CPD`  `.proto` format.

The above assumes `CPD` is self-contained which conflicts with below points about extensibility.
This could then looks like: 

<div class="bg-gray">
<img src="https://docs.google.com/drawings/d/e/2PACX-1vSI4Z7dP3TfnoFFAohwCATtC_JOcc1TPgGgaPrkcs2SdNjLPfwMUAQa2D5DmlbG_sI_s7HqJzv4VrAM/pub?w=1440&h=1080"/>
</div>

