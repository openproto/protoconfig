module github.com/thanos-io/OpenConfig/golang/protoc-gen-go-openconfig

go 1.15

require (
	github.com/pkg/errors v0.9.1
	github.com/thanos-io/OpenConfig/golang v0.0.0-00010101000000-000000000000
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.25.0
)

// TODO(bwplotka): Remove once dev period completes.
replace github.com/thanos-io/OpenConfig/golang => ../
