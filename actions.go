package main

import (
	"fmt"

	"github.com/anaskhan96/base58check"
	"github.com/urfave/cli"
)

func encode(c *cli.Context) error {

	data := c.String("data")
	version := c.String("version")

	if data == "" {
		return cli.NewExitError("No data provided", 1)
	}

	encoded, err := base58check.Encode(version, data)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(encoded)

	return nil
}

func decode(c *cli.Context) error {

	data := c.String("data")

	if data == "" {
		return cli.NewExitError("No data provided", 1)
	}

	decoded, err := base58check.Decode(data)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(decoded)

	return nil
}
