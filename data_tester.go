/*
The Data Tester tests data in a database
*/
package main

import (
	"os"
	"fmt"
	"sort"
	"github.com/mixo/gosql"
	"github.com/mixo/data-tester/tester"
	"github.com/urfave/cli/v2"
)

func main() {
	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func createApp() *cli.App {
	app := &cli.App{
		Description: "The Data Tester tests data in a database",
		Usage: " ",
		UsageText: "data-tester test test-options...",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:    "data_tester_db_driver",
				Aliases: []string{"dbd"},
				Usage:   "The database driver",
				EnvVars: []string{"data_tester_db_driver"},
			},
			&cli.StringFlag{
				Name:    "data_tester_db_host",
				Aliases: []string{"dbh"},
				Usage:   "The database host",
				EnvVars: []string{"data_tester_db_host"},
			},
			&cli.StringFlag{
				Name:    "data_tester_db_port",
				Aliases: []string{"dbt"},
				Usage:   "The database port",
				EnvVars: []string{"data_tester_db_port"},
			},
			&cli.StringFlag{
				Name:    "data_tester_db_user",
				Aliases: []string{"dbu"},
				Usage:   "The database user",
				EnvVars: []string{"data_tester_db_user"},
			},
			&cli.StringFlag{
				Name:    "data_tester_db_password",
				Aliases: []string{"dbp"},
				Usage:   "The database password",
				EnvVars: []string{"data_tester_db_password"},
			},
			&cli.StringFlag{
				Name:    "data_tester_db_name",
				Aliases: []string{"dbn"},
				Usage:   "The database name",
				EnvVars: []string{"data_tester_db_name"},
			},
		},
		Before: func(c *cli.Context) error {
			driver := c.String("data_tester_db_driver")
			host := c.String("data_tester_db_host")
			port := c.String("data_tester_db_port")
			user := c.String("data_tester_db_user")
			pass := c.String("data_tester_db_password")
			name := c.String("data_tester_db_name")
			if driver == "" || host == "" || port == "" || user == "" || pass == "" || name == "" {
				return cli.Exit("You must specify the database params", 1)
			}

			c.App.Metadata["db"] = gosql.DB{driver, host, port, user, pass, name}

			return nil
		},
		Commands: []*cli.Command{
			(tester.DayFluctuationTester{}).GetCliCommand(),
		},
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "Unknown test '%s'\n", command)
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
