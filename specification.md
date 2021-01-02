# OpenConfig, a software configuration convention

## Principles

Software `configuration` is a process of passing information to the software in order to change its behavior without recompilation or sometimes even without restarting (dynamic configuration).

Two main sides exists during software configuration process:

* `Configurable`: Software we want to configure, statically (during startup) or dynamically (on the run)
* `Configurator`: Another software or human user that desire to configure the `Configurable` automatically or manually (human).

The key element of `OpenConfig` specification is a strict format and rules of the `Configuration Definition`. This definition is the true single language that both `Configurator` and `Configurable` can use as shown on below diagram:

<img style="background-color:white" src="https://docs.google.com/drawings/d/e/2PACX-1vSANZkljSiDgV-o0a-dL0ryZz19p3Hblt5V_qozhBcY5ILq8j3T2GEAdCCHFHoSGT9h2H4LDqJ9bCn_/pub?w=1440&h=1080"/>

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

## Why Protocol Buffers?

> Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

On top of incredibly efficient and fast serialization capabilities it also works well [Interface Definition Language](https://en.wikipedia.org/wiki/Interface_description_language) and it is the most used and widely adopted IDL in the industry. For example, protobuf is the definition language and serialization format for [gRPC](https://grpc.io/).

Main reasons are that `proto` (version 3) language is **strongly typed, simple, consistent**, but also **backward and forward compatible**. In addition to it's stable and fast binary serialization format it be also easily decoded or encoded into human readable formats like YAML or JSON.

Last, but the least there exists variety of encoders for almost every coding language. You can read more [here](https://developers.google.com/protocol-buffers/docs/cpptutorial).

If that is not enough, `protobuf` ecosystem is [constantly evolving](https://twitter.com/bwplotka/status/1345076383190556676).

## Scope

Configure software in a way that is:

* **Discoverable** and **Type Safe**: *We don't need to guess or read complex documentations in order to know application configuration format, invariants and assumptions Configuration is defined by a single, typed [proto](https://developers.google.com/protocol-buffers) package.*
* **Allows Auto Generation**: *We don't need manually implement application client or data format to encode and similarly data format to parse. All can be safely generated thanks to protobuf in almost any programming language, thanks to wide plugin ecosystem. Additionally, you can generate a full CLI flag paring code or even documentation(!). It also means the smallest, possible impact in the event of application upgrade of modification. Type, Default or even Help changed can be easily tracked in any language.*
* **Easy to Verify**: *Non-semantic verification and validation (e.g ) does not require executable participation. No need for CLI --dry-run logic, or even invoking anything external. All can be validated from the point of protobuf.*
* **Backward & Forward Compatible**: *Smooth version migrations, thanks to protobuf guarantee and safety checkers like [`buf check breaking`](https://docs.buf.build/breaking-usage)*
* **Extensible**: *On top of [**OpenConfig Proto Extensions Format 1.0.**](proto/openconfig/v1/extensions.proto) this specification allows CLI or language-specific extensions (e.g see [`kingpinv2` Go package](golang/kingpinv2) and its [extensions](golang/kingpinv2/proto/openconfig/kingpinv2/v1/extensions.proto))*

## Not in the scope

* Runtime APIs, RPCs and Invocation semantics (e.g error codes, output, input).
* Configuration or Invocation API that will accept (e.g dynamically or statically) for configuration in `OpenConfig` format.
* Discoverability of `Configuration Proto Definitions` details.