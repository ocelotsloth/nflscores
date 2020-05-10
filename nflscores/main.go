package main

import (
	"log"
	"os"
	"sort"

	"git.sr.ht/~ocelotsloth/nflscores"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nflscores"
	app.Usage = "Scrapes nfl game scores."
	app.Version = "1.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "dump csv contents to stdout",
			Action: func(c *cli.Context) error {
				nflscores.WriteCSV(os.Stdout)
				return nil
			},
		},
		{
			Name:    "csv",
			Aliases: []string{"c"},
			Usage:   "Write csv to `FILE`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "Specify `FILE` location.",
				},
			},
			Action: func(c *cli.Context) error {
				file := c.String("file")
				if file == "" {
					cli.ShowCommandHelp(c, "csv")
				} else {
					nflscores.WriteFile(file)
				}
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
