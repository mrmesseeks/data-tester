/*
The Data Tester tests data in a database
Временный вариант, до переноса в gitLab
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
				Name:    "bidder_db_driver",
				Aliases: []string{"dbd"},
				Usage:   "The database driver",
				EnvVars: []string{"bidder_db_driver"},
			},
			&cli.StringFlag{
				Name:    "bidder_db_host",
				Aliases: []string{"dbh"},
				Usage:   "The database host",
				EnvVars: []string{"bidder_db_host"},
			},
			&cli.StringFlag{
				Name:    "bidder_db_port",
				Aliases: []string{"dbt"},
				Usage:   "The database port",
				EnvVars: []string{"bidder_db_port"},
			},
			&cli.StringFlag{
				Name:    "bidder_db_user",
				Aliases: []string{"dbu"},
				Usage:   "The database user",
				EnvVars: []string{"bidder_db_user"},
			},
			&cli.StringFlag{
				Name:    "bidder_db_pass",
				Aliases: []string{"dbp"},
				Usage:   "The database password",
				EnvVars: []string{"bidder_db_pass"},
			},
			&cli.StringFlag{
				Name:    "bidder_db_name",
				Aliases: []string{"dbn"},
				Usage:   "The database name",
				EnvVars: []string{"bidder_db_name"},
			},
		},
		Before: func(c *cli.Context) error {
			driver := c.String("bidder_db_driver")
			host := c.String("bidder_db_host")
			port := c.String("bidder_db_port")
			user := c.String("bidder_db_user")
			pass := c.String("bidder_db_pass")
			name := c.String("bidder_db_name")
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
