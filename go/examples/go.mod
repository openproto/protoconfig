module github.com/openproto/protoconfig/go/examples

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/openproto/protoconfig/go v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

// TODO(bwplotka): Remove once dev period completes.
replace github.com/openproto/protoconfig/go => ../
