package tester

import (
	"github.com/mixo/gosql"
	c "github.com/mixo/gocmd"
)

type Tester interface{
    GetName() string
    GetDescription() string
    GetArgs() c.ArgCollection
	Test(gosql.DB, c.ArgCollection) TestResult
}

type TestResult interface{
	IsOk() bool
	Show()
}
