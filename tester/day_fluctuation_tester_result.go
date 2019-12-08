package tester

import (
	"fmt"
	"github.com/mixo/godt"
	"time"
	"os"
	"github.com/olekukonko/tablewriter"
)

type DayFluctuationTesterResult struct {
	OK bool
	GroupsAvgParams, GroupsDayParams map[string]map[string]interface{}
	GroupsDiffs map[string]map[string]float64
	MaxDiff int
	StartDate, EndDate, Day time.Time
	QuantityColumn, GroupColumn string
}

func (r DayFluctuationTesterResult) IsOk() bool {
	return r.OK
}

func (r DayFluctuationTesterResult) Show() {
    headerRow := []string{"Group", "Period", "Row Quantity"}
    rows := [][]string{}
	isFirstGroup := true
	for group, avgParams := range r.GroupsAvgParams {
	    dayParams := r.GroupsDayParams[group]
	    diffParams := r.GroupsDiffs[group]

        avgRow := []string{group, "avg", fmt.Sprintf("%.2f", avgParams[r.QuantityColumn])}
        dayRow := []string{group, "day", fmt.Sprintf("%v", dayParams[r.QuantityColumn])}
        diffRow := []string{group, "diff", fmt.Sprintf("%.2f%%", diffParams[r.QuantityColumn])}

		for param, avgValue := range avgParams {
			if param == r.QuantityColumn || param == r.GroupColumn {
				continue
			}
			if isFirstGroup {
			    headerRow = append(headerRow, param)
			}

			dayValue := dayParams[param]
			diffValue := diffParams[param]

			avgRow = append(avgRow, fmt.Sprintf("%.2f", avgValue))
			dayRow = append(dayRow, fmt.Sprintf("%v", dayValue))
			diffRow = append(diffRow, fmt.Sprintf("%.2f%%", diffValue))
		}

		isFirstGroup = false
        rows = append(rows, avgRow)
        rows = append(rows, dayRow)
        rows = append(rows, diffRow)
	}

	rows = append([][]string{headerRow}, rows...)

    table := tablewriter.NewWriter(os.Stdout)
    table.SetAutoMergeCells(true)
    table.SetRowLine(true)
    table.AppendBulk(rows)

	var okString string
	if r.OK {
		okString = "passed :)"
	} else {
		okString = "failed :("
	}
	resultString := "Day Fluctuation test results\n"
	resultString += fmt.Sprintf("Test %s\n", okString)
	resultString += fmt.Sprintf("Max difference: %d%%\n", r.MaxDiff)
	resultString += fmt.Sprintf("Average period: %s - %s\n", godt.ToString(r.StartDate), godt.ToString(r.EndDate))
	resultString += fmt.Sprintf("Day: %s\n", godt.ToString(r.Day))

	fmt.Println(resultString)
    table.Render()
}
