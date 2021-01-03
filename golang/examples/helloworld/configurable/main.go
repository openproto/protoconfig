package main

import (
	"flag"
	"log"
	"os"

	helloworldpb "github.com/thanos-io/OpenConfig/golang/examples/helloworld"
)

func main() {
	log.SetFlags(0)

	var configEncoded string

	conf := &helloworldpb.HelloWorldConfiguration{}
	f := flag.NewFlagSet(conf.Metadata().Name, flag.ExitOnError)
	f.StringVar(&configEncoded, "openconfigv1", "", "configuration protobuf compatible with OpenConfig 1.0 standard")

	if err := f.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	if err := conf.DecodeString(configEncoded); err != nil {
		log.Fatal(err)
	}

	log.Println(helloworldpb.PrepareMessage(conf))
}
