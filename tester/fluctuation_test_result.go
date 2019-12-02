package tester

import (
	"fmt"
	"github.com/mixo/godt"
	"math"
	"time"
)

type FluctuationTestResult struct {
	OK                                   bool
	AvgRowCountPerDay, YesterdayRowCount uint
	Diff                                 int
	MaxDiff                              uint
	StartDate, EndDate, Yesterday        time.Time
}

func (r FluctuationTestResult) ToString() string {
	var okString, inequality string
	if r.OK {
		okString = "passed :)"
	} else {
		okString = "failed :("
	}

	from := godt.ToString(r.StartDate)
	to := godt.ToString(r.EndDate)
	dt := godt.ToString(r.Yesterday)

	if r.Diff == 0 {
		inequality = "equal to"
	} else if r.Diff < 0 {
		inequality = fmt.Sprintf("%d%% less than", int(math.Abs(float64(r.Diff))))
	} else {
		inequality = fmt.Sprintf("%d%% greater than", r.Diff)
	}

	resultString := "Fluctuation test results\n"
	resultString += fmt.Sprintf("Test %s\n", okString)
	resultString += fmt.Sprintf("Average row count: %d (%s - %s)\n", r.AvgRowCountPerDay, from, to)
	resultString += fmt.Sprintf("Yesterday row count: %d (%s) \n", r.YesterdayRowCount, dt)
	resultString += fmt.Sprintf("Max difference: %d%%\n", r.MaxDiff)
	resultString += fmt.Sprintf("Difference: %d%%\n", r.Diff)
	resultString += fmt.Sprintf("Description: the yesterday row count is %s the average row count\n", inequality)

	return resultString
}
