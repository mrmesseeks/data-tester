package tester

import (
	"github.com/mixo/gosql"
	"math"
	"time"
)

type FluctuationTester Tester

// maxDiff in percents
func (this FluctuationTester) TestYesterdayData(
	driver, host, port, user, password, database, tableName, dateColumn string,
	maxDiff, dayCount uint) FluctuationTestResult {
	db := gosql.DB{driver, host, port, user, password, database}

	yesterday := time.Now().AddDate(0, 0, -1)
	startDate := yesterday.AddDate(0, 0, -10)
	endDate := yesterday.AddDate(0, 0, -1)

	avgRowCountPerDay := db.GetAvgRowCountPerDay(tableName, dateColumn, startDate, endDate, dayCount)
	yesterdayRowCount := db.GetRowCountOnDate(tableName, dateColumn, yesterday)

	// percents
	diff := math.Round((float64(yesterdayRowCount) * 100 / float64(avgRowCountPerDay)) - 100)

	ok := uint(math.Abs(diff)) <= maxDiff

	return FluctuationTestResult{ok, avgRowCountPerDay, yesterdayRowCount, int(diff), maxDiff, startDate, endDate, yesterday}
}
