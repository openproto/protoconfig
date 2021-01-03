package openconfig

import (
	"os/exec"
)

type Configurator interface {
	Marshal() ([]byte, error)
	MarshalJSON() ([]byte, error)

	Metadata() Metadata
}

type Commander interface {
	NewExecCmd(name string) (*exec.Cmd, error)
}
