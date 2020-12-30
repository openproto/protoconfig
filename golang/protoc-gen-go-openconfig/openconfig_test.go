package main

import (
	"os"
	"testing"

	"github.com/efficientgo/tools/core/pkg/testutil"
)

const protocPathEnvVar = "PROTOC_PATH"

func TestIntegration(t *testing.T) {
	protocPath := os.Getenv(protocPathEnvVar)
	testutil.Assert(t, protocPath != "", "envvar %v not specified", protocPathEnvVar)

}