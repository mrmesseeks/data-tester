/*
The Data Tester tests data in a database
*/
package main

import (
	"fmt"
	"github.com/mixo/gocmd"
	t "github.com/mixo/data-tester/tester"
	"os"
)

const (
	usage = "Usage: dataTester table-name date-column max-diff day-count"
)

func main() {
	var (
		tester  t.FluctuationTester
		maxDiff uint
	)

	driver, host, port, user, password, database, tableName, dateColumn, maxDiff, dayCount := getParams()
	r := tester.TestYesterdayData(driver, host, port, user, password, database, tableName, dateColumn, maxDiff, dayCount)
	fmt.Println(r.ToString())

	if r.OK {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func getParams() (driver, host, port, user, password, database, tableName, dateColumn string, maxDiff, dayCount uint) {
	prefix := "data_tester_"
	driver = getEnvVar(prefix + "db_driver")
	host = getEnvVar(prefix + "db_host")
	port = getEnvVar(prefix + "db_port")
	user = getEnvVar(prefix + "db_user")
	password = getEnvVar(prefix + "db_password")
	database = getEnvVar(prefix + "db_name")
	tableName = gocmd.GetArg(0, usage)
	dateColumn = gocmd.GetArg(1, usage)
	maxDiff = gocmd.GetUintArg(2, usage)
	dayCount = gocmd.GetUintArg(3, usage)

	return
}

func getEnvVar(name string) (envVar string) {
	envVar, exists := os.LookupEnv(name)
	if !exists {
		panic("You must set env var " + name)
	}
	return
}
