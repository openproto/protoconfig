package main

import (
	"fmt"
	"log"
	"os/exec"

	openconfig "github.com/thanos-io/OpenConfig/golang"
	helloworldpb "github.com/thanos-io/OpenConfig/golang/examples/helloworld"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Command invoking go run ./configurable hello --world="my" --year=2021 --name="Kate B" --lang=ENGLISH --add-really
// using a few different methods.
func main() {
	runUsingArgs()
	runUsingOpenConfig1_0()
	runUsingOpenConfig1_0_NoPlugin()
}

func runUsingArgs() {
	log.Print("invoking configurable using args:")

	// Prone to errors, not generated, not maintainable, not safely typed.
	res, err := exec.Command("go run ../configurable", "hello", "--world=my", "--year=2021", "--name=Kate B", "--lang=1", "--add-really").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q\n", string(res))
}

func runUsingOpenConfig1_0() {
	c := &helloworldpb.HelloWorld{
		Command: helloworldpb.NewHelloCommand(&helloworldpb.HelloCommand{
			World:     "my",
			Year:      2021,
			Name:      "Kate B",
			Lang:      helloworldpb.Lang_ENGLISH,
			AddReally: true,
		}),
	}

	log.Print("invoking configurable using OpenConfig 1.0 (MarshalJSON):")
	b, err := c.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	res, err := exec.Command(fmt.Sprintf("go run ../%s", c.Metadata().Name), fmt.Sprintf("--openconfigv1=%s", b)).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q\n", string(res))
	log.Print("invoking configurable using OpenConfig 1.0 (Marshal):")

	b, err = c.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	res, err = exec.Command(fmt.Sprintf("go run ../%s", c.Metadata().Name), fmt.Sprintf("--openconfigv1=%s", b)).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q\n", string(res))
	log.Print("invoking configurable using OpenConfig 1.0 (NewExecCommand()):")

	cmd, err := c.NewExecCmd(fmt.Sprintf("go run ../%s", c.Metadata().Name))
	if err != nil {
		log.Fatal(err)
	}
	res, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%q\n", string(res))
	log.Printf("%q\n", string(res))
}

func runUsingOpenConfig1_0_NoPlugin() {
	log.Print("invoking configurable using OpenConfig 1.0 (with just using native protogen-go nothing else):")

	c := &helloworldpb.HelloWorld{
		Command: &helloworldpb.HelloWorld_Hello{
			Hello: &helloworldpb.HelloCommand{
				World:     "my",
				Year:      2021,
				Name:      "Kate B",
				Lang:      helloworldpb.Lang_ENGLISH,
				AddReally: true,
			},
		},
	}

	// Prone to runtime errors, but better than args.
	appName := proto.GetExtension(c.ProtoReflect().Descriptor().Options(), openconfig.E_Metadata).(*openconfig.Metadata).Name

	b, err := protojson.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	res, err := exec.Command(fmt.Sprintf("go run ../%s", appName), fmt.Sprintf("--openconfigv1=%s", b)).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q\n", string(res))
}
