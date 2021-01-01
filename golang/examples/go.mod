module github.com/thanos-io/OpenConfig/golang/examples

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/thanos-io/OpenConfig/golang v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
)

// TODO(bwplotka): Remove once dev period completes.
replace github.com/thanos-io/OpenConfig/golang => ../
