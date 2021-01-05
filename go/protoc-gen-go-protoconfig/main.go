//  protoc-gen-go-protoconfig is a plugin for the Google protocol buffer compiler to
// generate Go code. Install it by building this program and making it
// accessible within your PATH with the name:
//	 protoc-gen-go-protoconfig
//
// The 'go-protoconfig' suffix becomes part of the argument for the protocol compiler,
// such that it can be invoked as:
//	protoc --go-protoconfig_out=. path/to/file.proto
//
// This generates Go configuration definitions for the protocol buffer defined by
// file.proto. With that input, the output will be written to path/to/file_protoconfig.pb.go
package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "0.1.0"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Fprintf(os.Stdout, "protoc-gen-go-grpc %v\n", version)
		return
	}

	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := generateGoProtoConfig(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}
