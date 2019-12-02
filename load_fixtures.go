package main

import (
	"fmt"
	"github.com/mixo/data-tester/fixture"
	"time"
)

var (
	startDate          = time.Now().AddDate(0, 0, -11)
	endDate            = startDate.AddDate(0, 0, 10)
	rowCountPerDayFrom = 100
	rowCountPerDayTo   = 110
	maxDiff            = uint(15)
	fluctuations       = []fixture.Fluctuation{
		fixture.Fluctuation{endDate, 1, maxDiff},
	}
	tableName  = "datatester_fixture"
	columns    = []string{"date"}
	columnsSql = []string{"\"date\" date"}
)

func main() {
	fmt.Println("Load fixtures")
	dataLoader := fixture.DataLoader{tableName, columns, columnsSql}
	dataLoader.Load(startDate, endDate, rowCountPerDayFrom, rowCountPerDayTo, fluctuations)
}
