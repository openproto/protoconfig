module github.com/thanos-io/OpenConfig/golang/protoc-gen-go-openconfig

go 1.15

require (
	github.com/efficientgo/tools/core v0.0.0-20201228212819-69909db83cda
	github.com/pkg/errors v0.9.1
	github.com/thanos-io/OpenConfig/golang v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.25.0
)

// TODO(bwplotka): Remove once dev period completes.
replace github.com/thanos-io/OpenConfig/golang => ../
