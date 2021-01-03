package main

import (
	"bytes"
	"log"
	"os"

	helloworldpb "github.com/thanos-io/OpenConfig/golang/examples/helloworld"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	log.SetFlags(0)

	var (
		configEncoded      string
		helloCmd           helloworldpb.HelloCommand
		byeConfigurableCmd helloworldpb.ByeConfigurableCommand
		langEnum           string
		usage              bytes.Buffer
	)

	conf := &helloworldpb.HelloWorldConfiguration{}
	f := kingpin.New(conf.Metadata().Name, conf.Metadata().Description)
	f.UsageWriter(&usage)
	f.Terminate(func(i int) {
		if configEncoded != "" {
			if err := conf.DecodeString(configEncoded); err != nil {
				log.Fatal(err)
			}
			log.Println(helloworldpb.PrepareMessage(conf))
			os.Exit(0)
		}
		os.Exit(1)
	})

	// First input method: OpenConfig.
	f.Flag("openconfigv1", "configuration protobuf compatible with OpenConfig 1.0 standard").StringVar(&configEncoded)

	// Second: Manually written CLI. It's a bit tedious and prone to errors, so see `configurable-kingpinv2-gen` for better approach.
	c := f.Command("hello", "TODO")
	c.Flag("name", "TODO").StringVar(&helloCmd.Name)
	c.Flag("year", "TODO").Int64Var(&helloCmd.Year)
	c.Flag("world", "TODO").StringVar(&helloCmd.World)
	c.Flag("lang", "TODO").EnumVar(&langEnum, "ENGLISH", "POLISH", "GERMAN", "UNSPECIFIED")
	c.Flag("please-add-really", "TODO").BoolVar(&helloCmd.PleaseAddReally) //nolint (deprecated, but we want to be backward compatible).
	c.Flag("add-really", "TODO").BoolVar(&helloCmd.AddReally)

	b := f.Command("bye", "TODO")
	b.Flag("lang", "TODO").EnumVar(&langEnum, "ENGLISH", "POLISH", "GERMAN", "UNSPECIFIED")
	_ = b.Command("just", "TODO")

	// TODO(bwplotka): Actually implement.
	_ = b.Command("configurable", "TODO")

	cmd, err := f.Parse(os.Args[1:])
	if err != nil {
		log.Println(usage)
		log.Fatal(err)
	}

	switch cmd {
	case "hello":
		// YOLO!
		helloCmd.Lang = helloworldpb.Lang(helloworldpb.Lang_value[langEnum])
		conf.Command = helloworldpb.NewHelloCommand(&helloCmd)
	case "bye just":
		conf.Command = helloworldpb.NewByeCommand(&helloworldpb.ByeCommand{
			Lang:    helloworldpb.Lang(helloworldpb.Lang_value[langEnum]),
			Command: helloworldpb.NewByeJustCommand(&helloworldpb.ByeJustCommand{}),
		})
	case "bye configurable":
		conf.Command = helloworldpb.NewByeCommand(&helloworldpb.ByeCommand{
			Lang:    helloworldpb.Lang(helloworldpb.Lang_value[langEnum]),
			Command: helloworldpb.NewByeConfigurableCommand(&byeConfigurableCmd),
		})
	}

	log.Println(helloworldpb.PrepareMessage(conf))
}
