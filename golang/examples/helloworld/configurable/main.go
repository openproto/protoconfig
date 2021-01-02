package main

import (
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "kənˈtrīv",
		Version:  "v19.99.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Example Human",
				Email: "human@example.com",
			},
		},
		Copyright: "(c) 1999 Serious Enterprise",
		HelpName:  "contrive",
		Usage:     "demonstrate available API",
		UsageText: "contrive - demonstrating the available API",
		ArgsUsage: "[args and such]",
		Commands:  []*cli.Command{},
	}
	app.RunContext()
}
