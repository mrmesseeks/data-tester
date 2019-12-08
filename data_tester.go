/*
The Data Tester tests data in a database
*/
package main

import (
	"fmt"
	"github.com/mixo/gosql"
	"github.com/mixo/gocmd"
	t "github.com/mixo/data-tester/tester"
	"os"
	"strings"
	"sort"
)

const (
	basicUsage = "Usage: data-tester <test> [args]\n" +
		"There are the following tests:\n"
)

var (
	testers = map[string]t.Tester{
		t.DayFluctuationTester{}.GetName(): t.DayFluctuationTester{},
	}
)

func main() {
	driver, host, port, user, password, database := getParams()
	db := gosql.DB{driver, host, port, user, password, database}

	usage := composeUsage(basicUsage, testers)

	argIndex := 0
	testName := gocmd.GetArg(argIndex, usage)
	tester := getTestByName(testName)

	argIndex++
	args := gocmd.InjectArgs(tester.GetArgs(), argIndex, usage)

	r := testers[testName].Test(db, args)
	r.Show()

	if r.IsOk() {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func getTestByName(testName string) t.Tester {
	tester := testers[testName]
	if tester == nil {
		fmt.Printf("A '%s' test doesn't exist. There are the following tests: %s\n", testName, getTesterNames(testers))
		os.Exit(1)
	}
	return tester
}

func composeUsage(basicUsage string, testers map[string]t.Tester) string {
	names := []string{}
	for name, _ := range testers {
		names = append(names, name)
	}

	usages := []string{basicUsage}
	sort.Strings(names)
	for _, name := range names {
		tester := testers[name]

		usage := fmt.Sprintf("\t%s\t%s\n", tester.GetName(), tester.GetDescription())
		usage += fmt.Sprintf("\t\t\t\t%s\n", createUsageForArgs(tester.GetArgs(), "\t\t\t\t"))

		usages = append(usages, usage)
	}

	return strings.Join(usages, "\n")
}

func createUsageForArgs(args gocmd.ArgCollection, argPrefix string) (usage string) {
	usageArgs := make([]string, 0)
	for _, arg := range args.GetAll() {
		usageArgs = append(usageArgs, fmt.Sprintf("<%s>", arg.GetName()))
	}

	usage = "Arguments: " + strings.Join(usageArgs, " ") + "\n"
	argDescriptions := make([]string, 0)
	for _, arg := range args.GetAll() {
		argDescriptions = append(argDescriptions, fmt.Sprintf("%s<%s>\t%s", argPrefix, arg.GetName(), arg.GetDescription()))
	}
	usage += strings.Join(argDescriptions, "\n")

	return
}

func getParams() (driver, host, port, user, password, database string) {
	prefix := "data_tester_"
	driver = getEnvVar(prefix + "db_driver")
	host = getEnvVar(prefix + "db_host")
	port = getEnvVar(prefix + "db_port")
	user = getEnvVar(prefix + "db_user")
	password = getEnvVar(prefix + "db_password")
	database = getEnvVar(prefix + "db_name")

	return
}

func getEnvVar(name string) (envVar string) {
	envVar, exists := os.LookupEnv(name)
	if !exists {
		fmt.Println("You must set env var " + name)
		os.Exit(1)
	}
	return
}

func getTesterNames(testers map[string]t.Tester) string {
	names := []string{}
	for name, _ := range testers {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
