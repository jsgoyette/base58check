package main

import (
	"os"
	"sort"

	"github.com/urfave/cli"
)

/*
 * Resources:
 * https://en.bitcoin.it/wiki/List_of_address_prefixes
**/

func main() {

	// set up cli
	app := cli.NewApp()

	app.Usage = "Base58Check utility"
	app.UsageText = "base58check"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{}
	app.Before = beforeApp
	app.Commands = commands

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}

func beforeApp(c *cli.Context) error {
	return nil
}
