package main

import (
	"flag"
	"log"
	"os"

	helloworldpb "github.com/thanos-io/OpenConfig/golang/examples/helloworld"
)

func main() {
	var configEncoded string

	f := flag.NewFlagSet("configurable", flag.ExitOnError)
	f.StringVar(&configEncoded, "openconfigv1", "", "configuration protobuf compatible with OpenConfig 1.0 standard")

	if err := f.Parse(os.Args); err != nil {
		log.Fatal(err)
	}

	conf := &helloworldpb.HelloWorld{}
	if configEncoded != "" {
		if err := conf.UnmarshalString(configEncoded); err != nil {
			log.Fatal(err)
		}
	}
}
