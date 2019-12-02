package tester

import (
	"fmt"
	"github.com/mixo/data-tester/fixture"
	"testing"
	"time"
)

var (
	startDate          = time.Now().AddDate(0, 0, -11)
	endDate            = startDate.AddDate(0, 0, 10)
	rowCountPerDayFrom = 100
	rowCountPerDayTo   = 110
	maxDiff            = uint(15)
	dayCount           = uint(10)
	fluctuations       = []fixture.Fluctuation{
		fixture.Fluctuation{endDate, 1, maxDiff},
	}
	driver     = "mysql"
	host       = "127.0.0.1"
	port       = "3306"
	user       = "root"
	password   = "root"
	database   = "test"
	tableName  = "datatester_fixture"
	dateColumn = "date"
	columns    = []string{"date"}
	columnsSql = []string{"`date` date"}
	//columnsSql = []string{"\"date\" date"}
)

func TestTestYesterdayDataFailed(t *testing.T) {
	dataLoader := fixture.DataLoader{tableName, columns, columnsSql}
	dataLoader.Load(startDate, endDate, rowCountPerDayFrom, rowCountPerDayTo, fluctuations)
	defer dataLoader.Unload()

	var tester FluctuationTester
	r := tester.TestYesterdayData(driver, host, port, user, password, database, tableName, dateColumn, maxDiff, dayCount)
	fmt.Println(r.ToString())
	if r.OK {
		t.Error("The test must be failed")
	}
}

func TestTestYesterdayDataPass(t *testing.T) {
	dataLoader := fixture.DataLoader{tableName, columns, columnsSql}
	dataLoader.Load(startDate, endDate, rowCountPerDayFrom, rowCountPerDayTo, []fixture.Fluctuation{})
	//defer dataLoader.Unload()

	var tester FluctuationTester
	r := tester.TestYesterdayData(driver, host, port, user, password, database, tableName, dateColumn, maxDiff, dayCount)
	fmt.Println(r.ToString())
	if !r.OK {
		t.Error("The test must be passed")
	}
}
