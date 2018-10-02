package main

import (
	"log"
	"os"
	"sort"

	"github.com/ocelotsloth/nflscores"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nflscores"
	app.Usage = "Scrapes nfl game scores."
	app.Version = "1.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Specify `FILE` location.",
			Value: ".",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "dump csv contents to stdout",
			Action: func(c *cli.Context) error {
				nflscores.WriteStdout()
				return nil
			},
		},
		{
			Name:    "csv",
			Aliases: []string{"c"},
			Usage:   "Write csv to `FILE",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

} // main()
