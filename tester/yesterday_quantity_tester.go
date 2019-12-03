package tester

import (
	"github.com/mixo/gosql"
	c "github.com/mixo/gocmd"
	"math"
	"time"
)

type YesterdayQuantityTester struct{}

func (this YesterdayQuantityTester) GetName() string {
	return "yesterday-quantity"
}

func (this YesterdayQuantityTester) GetDescription() string {
	return "Checks that an yesterday number of rows doesn't differ" +
		"from an average number of rows per day"
}

func (this YesterdayQuantityTester) GetArgs() c.ArgCollection {
	return c.NewArgCollection([]c.Arg{
		c.NewArg("table-name", "", true, "Specify the table name you want to test"),
		c.NewArg("date-column", "", true, "Specify the date column that will be used to count rows"),
		c.NewArg("max-diff", 0, true, "Specify the maximum difference in percents between the yesterday row count and the average row count "),
		c.NewArg("day-count", 0, true, "Specify the number of days for which to calculate the average row count"),
	})
}

func (this YesterdayQuantityTester) Test(db gosql.DB, args c.ArgCollection) TestResult {
	yesterday := time.Now().AddDate(0, 0, -1)
	startDate := yesterday.AddDate(0, 0, -10)
	endDate := yesterday.AddDate(0, 0, -1)

	avgRowCountPerDay := db.GetAvgRowCountPerDay(args.G("table-name").(string), args.G("date-column").(string), startDate, endDate, args.G("day-count").(int))
	yesterdayRowCount := db.GetRowCountOnDate(args.G("table-name").(string), args.G("date-column").(string), yesterday)

	// percents
	diff := math.Round((float64(yesterdayRowCount) * 100 / float64(avgRowCountPerDay)) - 100)

	ok := int(math.Abs(diff)) <= args.G("max-diff").(int)

	return YesterdayQuantityTesterResult{ok, avgRowCountPerDay, yesterdayRowCount, int(diff), args.G("max-diff").(int), startDate, endDate, yesterday}
}
