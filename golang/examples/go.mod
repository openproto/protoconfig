module github.com/thanos-io/OpenConfig/golang/examples

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/thanos-io/OpenConfig/golang v0.0.0-20201230221506-feea1e3c8bd7 // indirect
	google.golang.org/protobuf v1.25.0
)

// TODO(bwplotka): Remove once dev period completes.
replace github.com/thanos-io/OpenConfig/golang => ../
