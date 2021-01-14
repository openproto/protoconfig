# ProtoConfig 1.0

[![golang docs](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/protoconfig/protoconfig/go) [![Latest Release](https://img.shields.io/github/release/protoconfig/protoconfig.svg?style=flat-square)](https://github.com/protoconfig/protoconfig/releases/latest)

The `ProtoConfig 1.0` is an open, language agnostic specification that describes a process of using, defining, and consuming software configuration input in a unified way.

*Like OpenAPI but for Configuration!*

Designed to support all popular configuration patterns: Cloud Native `Configuration as a Code`, robust command line interfaces (CLI), static or dynamic configurability as well as remote configuration APIs.

### TL;DR

The `ProtoConfig 1.0` standardises application configuration, by describing the process as follows:

1. Application developer creates a [.proto](https://developers.google.com/protocol-buffers) file with `Configuration Proto Definition`, that defines what exactly options application has.
2. Thanks to the single source of truth for your app configuration, a developer can generate a data structure in the language they need. Such structs/classes can be then used to parse incoming input in native proto or JSON encoding (for example from `stdin`, CLI flag or HTTP API, etc).
* `ProtoConfig` allows generation extensibility, so such definition can be also used to generate documentation, `--help`, `man`, custom types, or even whole CLI flag & args parsing code. Options are unlimited!
3. Following the `ProtoConfig` convention, a human or program can now configure the application either by specifying JSON manually or by generating (again) data structure from `Configuration Proto Definition` in the programming language they want. Such struct/class is typed (if a language is typed), have a valid format, have help commentaries and can be encoded to `protobuf` supported format and passed to the `configurable` application!

`ProtoConfig 1.0` standard specifies that every `ProtoConfig 1.0` compatible application accepts encoded `protobuf` in proto or JSON format as the way of configuring itself. Either as the only way or in addition to other conventions (e.g args and flags).

![Open Config 1.0](https://docs.google.com/drawings/d/e/2PACX-1vSANZkljSiDgV-o0a-dL0ryZz19p3Hblt5V_qozhBcY5ILq8j3T2GEAdCCHFHoSGT9h2H4LDqJ9bCn_/pub?w=1440&h=1080)

### Goals

Configure software in a way that is:

* **Discoverable** and **Type-Safe**: *We don't need to guess or read complex documentations to know application configuration format, invariants, and assumptions Configuration is defined by a single, typed [proto](https://developers.google.com/protocol-buffers) package.*
* **Allows Auto Generation**: *We don't need to manually implement application client or data format to encode and similarly data format to parse. All can be safely generated thanks to `protobuf` in almost any programming language, thanks to the wide plugin ecosystem. Additionally, you can generate a full CLI flag paring code or even documentation(!). It also means the smallest, possible impact in the event of application upgrade or modification. Type, Default, or even Help changed can be easily tracked in any language.*
* **Easy to Verify**: *Non-semantic verification and validation (e.g ) do not require executable participation. No need for CLI --dry-run logic, or even invoking anything external. All can be validated from the point of `protobuf`.* https://github.com/envoyproxy/protoc-gen-validate
* **Backward & Forward Compatible**: *Smooth version migrations, thanks to `protobuf` guarantee and safety checkers like [`buf check breaking`](https://docs.buf.build/breaking-usage)*
* **Extensible**: *On top of [**Proto Config Extensions Format 1.0.**](proto/protoconfig/v1/extensions.proto) this specification allows CLI or language-specific extensions (e.g see [`kingpinv2` Go package](go/kingpinv2) and its [extensions](go/kingpinv2/proto/protoconfig/kingpinv2/v1/extensions.proto))*

### Motivation

See "Configuration in 2021 is still broken" blog post (TBD).

### Principles

See [./specification#principles](specification.md#principles).

### Why Protocol Buffers and What is it

See [./specification#why-protocol-buffers](specification.md#why-protocol-buffers)

### Example

See [example in Go](go/examples/README.md).

If you are not familiar with Go, this is still a useful example, as the flow and pattern will be similar to any language.

## This Repository

| Item   | What | Status |
|--------|------|--------|
| [`proto/protoconfig/v1/extensions.proto`](proto/protoconfig/v1/extensions.proto)  |  *Proto Config Extensions Format 1.0* | Alpha |
| [`proto/examples/helloworld/v1/helloworld.proto`]( proto/examples/helloworld/v1/helloworld.proto) |  Example *Configuration Proto Definitions (CPD)* | Alpha |
| [`go/`](go)  | Go module with (optional) *Proto Config Extensions Format 1.0* bindings | Alpha |
| [`go/protoc-gen-go-protoconfig`](go/protoc-gen-go-protoconfig/README.md)  | Go module with (optional) protogen go plugin supporting *Proto Config Extensions Format 1.0* Go language  | Alpha |
| [`go/examples`](go/examples/README.md) | Go module with (optional) protogen go plugin supporting *Proto Config Extensions Format 1.0* Go language  | Alpha |

### Contributing

Any help wanted. Do you have an idea, want to help, don't know how to start helping?

Put an issue on this repo, create PR or ping us on the CNCF Slack (`@bwplotka`, `@brancz`)! 🤗

## Other related projects

* [OpenConfig](https://www.openconfig.net/): OpenConfig aims for something very similar, however aimed for network management configuration. Defined in `YANG` (some sort of `XML`) instead of more modern `protobuf` which was designed for code generation and IDL.
* [protoconf](https://github.com/protoconf/protoconf): Amazing project inspired by Facebook's Configurator. While sounding similar it's essentially a protobuf based dynamic configuration with its own GitOps pipeline. `ProtoConfig` does not specify how to 
  do dynamic configuration, leaving that to extensions that can be added to specification itself (e.g language specific generator that generates lib for dynamic reloading).
* [OpenAPI/Swagger](https://swagger.io/specification/): OpenAPI aims for API definitions. Configuration can be treated as some part of the API (payload?), however scoping definition purely for configuration, simplifies the whole topic a lot. `ProtoConfig` does not need to care about status codes, HTTP specifics or actually even OS/Process details. `ProtoConfig` can be applied   

### Roadmap

Help wanted!

* [ ] Documentation for using ProtoConfig 1.0 by different language than supported by this repo.
* [ ] Documentation for writing ProtoConfig 1.0 plugin for a different language than supported by this repo.
* [ ] Publish formal RFC-compatible specification, finish it, add recommendations (e.g package names).
* [ ] Document extensibility.
* [ ] Add example if any configuration templating engine.
* [ ] Large `protobuf` are hard to use. Standard should specify some solution when `Encoded Configuration Message` is larger than Megabytes.
* [ ] Unit tests.

## Initial Authors

* Bartek Płotka @bwplotka
* Frederic Branczyk @brancz
