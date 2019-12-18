package tester

import (
	"github.com/urfave/cli/v2"
)

type Tester interface{
    GetCliCommand() *cli.Command
}

type TestResult interface{
	IsOk() bool
	Show()
}
