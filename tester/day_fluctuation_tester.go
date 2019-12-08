package tester

import (
	"github.com/mixo/gosql"
	c "github.com/mixo/gocmd"
	"math"
	"time"
	"strings"
	"fmt"
)

type DayFluctuationTester struct{}

func (this DayFluctuationTester) GetName() string {
	return "day-fluctuation"
}

func (this DayFluctuationTester) GetDescription() string {
	return "Checks that a day number of rows doesn't differ" +
		"from an average number of rows per day"
}

func (this DayFluctuationTester) GetArgs() c.ArgCollection {
	return c.NewArgCollection([]c.Arg{
		c.NewArg("table-name", "", true, "Specify the table name you want to test"),
		c.NewArg("date-column", "", true, "Specify the date column that will be used to count rows"),
		c.NewArg("day-indent", 0, true, "For which day, check a row count. For example, 1 - tomorrow, 0 - today, -1 - yesterday, -2 the day before yesterday and so on"),
		c.NewArg("max-diff", 0, true, "Specify the maximum difference in percents between the yesterday row count and the average row count "),
		c.NewArg("day-count", 0, true, "Specify the number of days for which to calculate the average row count"),
		c.NewArg("numeric-columns", "", true, "Specify comma-separated numeric columns which should be checked. For example, amount,price,profit"),
		c.NewArg("group-column", "", true, "Specify the group column that will be used to count and sum columns within groups"),
		c.NewArg("filtered-groups", "", false, "Specify comma-separated groups that should be checked. Leave the argument empty '', if you want to check all groups"),
	})
}

func (this DayFluctuationTester) Test(db gosql.DB, args c.ArgCollection) TestResult {
	quantityColumn := "_quantity"
	tableName := args.G("table-name").(string)
	dateColumn := args.G("date-column").(string)
    dayCount := args.G("day-count").(int)
	day := time.Now().AddDate(0, 0, args.G("day-indent").(int))
	startDate := day.AddDate(0, 0, -dayCount)
	endDate := day.AddDate(0, 0, -1)
	numericColumns := strings.Split(args.G("numeric-columns").(string), ",")
	groupColumn := args.G("group-column").(string)
	maxDiff := args.G("max-diff").(int)
	filteredGroupStrings := strings.Split(args.G("filtered-groups").(string), ",")
	filteredGroups := make([]interface{}, 0)
	for _, filteredGroupString := range filteredGroupStrings {
		if filteredGroupString != ""  {
			filteredGroups = append(filteredGroups, filteredGroupString)
		}
	}

	avgParams := db.GetAvgRowParamsPerDay(tableName, dateColumn, startDate, endDate, dayCount, quantityColumn, numericColumns, groupColumn, filteredGroups)
	dayParams := db.GetRowParamsOnDate(tableName, dateColumn, day, quantityColumn, numericColumns, groupColumn, filteredGroups)
	groupsAvgParams := this.groupMaps(avgParams, groupColumn)
	groupsDayParams := this.groupMaps(dayParams, groupColumn)

	columns := append(numericColumns, quantityColumn)
	groupsDiffs, ok := this.compareGroupsParams(columns, maxDiff, groupsAvgParams, groupsDayParams)

	return DayFluctuationTesterResult{ok, groupsAvgParams, groupsDayParams, groupsDiffs, maxDiff, startDate, endDate, day, quantityColumn, groupColumn}
}

func (this DayFluctuationTester) groupMaps(maps []map[string]interface{},
	groupColumn string) map[string]map[string]interface{} {
	var group string

	groupMaps := make(map[string]map[string]interface{}, 0)
	for _, singleMap := range maps {
        switch singleMap[groupColumn].(type) {
            case int:
                group = fmt.Sprintf("%d", singleMap[groupColumn].(int))
            case string:
                group = singleMap[groupColumn].(string)
            default:
                panic(fmt.Sprintf("Undefined type %T", singleMap[groupColumn]))
        }
		groupMaps[group] = singleMap
	}

	return groupMaps
}

func (this DayFluctuationTester) compareGroupsParams(columns []string, maxDiff int,
	groupsAvgParams, groupsDayParams map[string]map[string]interface{}) (map[string]map[string]float64, bool) {
	ok := true
	groupsDiffs := make(map[string]map[string]float64, 0)
	for group, groupAvgParams := range groupsAvgParams {
		groupsDiffs[group] = make(map[string]float64, 0)
		for _, param := range columns {
		    var dayValue float64
			if groupsDayParams[group] == nil {
			    groupsDayParams[group] = make(map[string]interface{}, 0)
			}
			if groupsDayParams[group][param] == nil {
			    groupsDayParams[group][param] = 0.0
			}

            switch groupsDayParams[group][param].(type) {
                case int:
                    dayValue = float64(groupsDayParams[group][param].(int))
                case int8:
                    dayValue = float64(groupsDayParams[group][param].(int8))
                case int16:
                    dayValue = float64(groupsDayParams[group][param].(int16))
                case int32:
                    dayValue = float64(groupsDayParams[group][param].(int32))
                case int64:
                    dayValue = float64(groupsDayParams[group][param].(int64))
                default:
                    dayValue = groupsDayParams[group][param].(float64)
            }

			diff := this.calculateDiffInPercents(dayValue, groupAvgParams[param].(float64))
			groupsDiffs[group][param] = diff

			ok = ok && int(math.Abs(diff)) <= maxDiff
		}
	}

	return groupsDiffs, ok
}

func (this DayFluctuationTester) calculateDiffInPercents(dayRowCount, avgRowCountPerDay float64) float64 {
    return math.Round((dayRowCount * 100 / avgRowCountPerDay) - 100)
}
