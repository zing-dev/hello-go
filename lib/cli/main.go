package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
		Commands: cli.Commands{
			{
				Name: "test",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "bool",
						Value: true,
					},
					&cli.Int64Flag{
						Name:  "int64",
						Value: 1,
					},
				},
				Action: func(c *cli.Context) error {
					for _, v := range c.FlagNames() {
						fmt.Println(v, c.Value(v))
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
