module github.com/thanos-io/OpenConfig/go/examples

go 1.15

require (
	github.com/thanos-io/OpenConfig/golang feea1e3c8bd738431c63b174060c2ce9e054890f
	github.com/golang/protobuf v1.4.3
	google.golang.org/protobuf v1.25.0
)

// TODO(bwplotka): Remove once dev period completes.
replace github.com/thanos-io/OpenConfig/golang => ../