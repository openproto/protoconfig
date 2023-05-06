module github.com/openproto/protoconfig/go/examples

go 1.19

require (
	github.com/alecthomas/kingpin/v2 v2.3.2
	github.com/openproto/protoconfig/go v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
)

require (
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
)

// TODO(bwplotka): Remove once dev period completes.
replace github.com/openproto/protoconfig/go => ../
