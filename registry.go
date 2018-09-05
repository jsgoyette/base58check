package main

import (
	"github.com/urfave/cli"
)

var (
	dataFlag = cli.StringFlag{
		Name:  "data, d",
		Value: "",
		Usage: "Data to be processed",
	}
	versionFlag = cli.StringFlag{
		Name:  "version, v",
		Value: "00",
		Usage: "Version hex string",
	}

	commands = []cli.Command{
		{
			Name:        "encode",
			Usage:       "Encode data",
			UsageText:   "base58check encode [options]",
			Description: "Encode the provided data with the base58check algorithm",
			Action:      encode,
			Flags:       []cli.Flag{dataFlag, versionFlag},
		},
		{
			Name:        "decode",
			Usage:       "Decode data",
			UsageText:   "base58check decode [options]",
			Description: "Decode the provided data with the base58check algorithm",
			Action:      decode,
			Flags:       []cli.Flag{dataFlag, versionFlag},
		},
	}
)
