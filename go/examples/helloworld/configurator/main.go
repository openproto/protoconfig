package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	protoconfig "github.com/openproto/protoconfig/go"
	helloworldpb "github.com/openproto/protoconfig/go/examples/helloworld"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func runCmdAndPrintResult(dir string, args ...string) {
	c := exec.Command(args[0], args[1:]...)
	c.Dir = dir
	b := bytes.Buffer{}
	c.Stdout = &b
	c.Stderr = &b
	if err := c.Run(); err != nil {
		log.Fatal(err, b.String())
	}
	log.Printf("[%v] => %s\n", args, b.String())
}

// Command invoking go run ./configurable-kingpinv2 hello --world="my" --year=2077 --name="Alt C" --lang=ENGLISH --add-really
// using a few different methods.
func main() {
	log.SetFlags(0)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	runUsingArgs(pwd)
	runUsingProtoConfig1_0_NoPlugin(pwd)
	runUsingProtoConfig1_0(pwd)
}

func runUsingArgs(dir string) {
	log.Println("invoking configurable using args:")

	// Prone to errors, not generated, not maintainable, not safely typed, etc..
	runCmdAndPrintResult(dir, "go", "run", "./configurable-kingpinv2", "hello", "--world=my", "--year=2077", "--name=Alt C", "--lang=ENGLISH", "--add-really")
}

func runUsingProtoConfig1_0_NoPlugin(dir string) {
	log.Println("invoking configurable using ProtoConfig 1.0 (with just using native protogen-go nothing else!):")

	c := &helloworldpb.HelloWorldConfiguration{
		Command: &helloworldpb.HelloWorldConfiguration_Hello{
			Hello: &helloworldpb.HelloCommand{
				World:     "my",
				Year:      2077,
				Name:      "Alt C",
				Lang:      helloworldpb.Lang_ENGLISH,
				AddReally: true,
			},
		},
	}

	// Prone to runtime errors, but better than args.
	appName := proto.GetExtension(c.ProtoReflect().Descriptor().Options(), protoconfig.E_Metadata).(*protoconfig.Metadata).Name

	b, err := protojson.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	runCmdAndPrintResult(dir, "go", "run", "./"+appName, "--protoconfigv1="+string(b))
	runCmdAndPrintResult(dir, "go", "run", "./"+appName+"-kingpinv2", "--protoconfigv1="+string(b))
}

func runUsingProtoConfig1_0(dir string) {
	c := &helloworldpb.HelloWorldConfiguration{
		Command: helloworldpb.NewHelloCommand(&helloworldpb.HelloCommand{
			World:     "my",
			Year:      2077,
			Name:      "Alt C",
			Lang:      helloworldpb.Lang_ENGLISH,
			AddReally: true,
		}),
	}

	log.Println("invoking configurable using ProtoConfig 1.0 (EncodeJSON):")
	b, err := c.EncodeJSON()
	if err != nil {
		log.Fatal(err)
	}

	runCmdAndPrintResult(dir, "go", "run", "./"+c.Metadata().Name, "--protoconfigv1="+string(b))
	runCmdAndPrintResult(dir, "go", "run", "./"+c.Metadata().Name+"-kingpinv2", "--protoconfigv1="+string(b))

	log.Println("invoking configurable using ProtoConfig 1.0 (Marshal):")

	b, err = c.Encode()
	if err != nil {
		log.Fatal(err)
	}

	runCmdAndPrintResult(dir, "go", "run", "./"+c.Metadata().Name, "--protoconfigv1="+string(b))
	runCmdAndPrintResult(dir, "go", "run", "./"+c.Metadata().Name+"-kingpinv2", "--protoconfigv1="+string(b))

	log.Println("invoking configurable using ProtoConfig 1.0 (CommandLineArgument()):")

	arg, err := c.CommandLineArgument()
	if err != nil {
		log.Fatal(err)
	}

	runCmdAndPrintResult(dir, "go", "run", "./"+c.Metadata().Name, arg)
	runCmdAndPrintResult(dir, "go", "run", "./"+c.Metadata().Name+"-kingpinv2", arg)
}
