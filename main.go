package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var port string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "baseurl",
			Value: "http://localhost:",
			Usage: "Base API URL",
		},
		cli.StringFlag{
			Name:        "port",
			Value:       "9200",
			Usage:       "port that Elasticsearch is listening on",
			Destination: &port,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "greet",
			Aliases: []string{"g"},
			Usage:   "greet user",
			Action: func(c *cli.Context) error {
				name := "Nefertiti"
				if c.NArg() > 0 {
					name = c.Args().Get(0)
				}
				if c.String("lang") == "spanish" {
					fmt.Println("Hola", name)
				} else {
					fmt.Println("Hello", name)
				}
				return nil
			},
		},
		{
			Name:      "cluster",
			ShortName: "cl",
			Usage:     "Get cluster information",
			Subcommands: []cli.Command{
				{
					Name:    "health",
					Aliases: []string{"he"},
					Usage:   "get cluster health",
					Action: func(c *cli.Context) error {
						query := cmdCluster(c, port, "health")
						fmt.Println(query)
						fmt.Println(getJSON(query))
						return nil
					},
				},
				{
					Name:    "state",
					Aliases: []string{"s"},
					Usage:   "get cluster state",
					Action: func(c *cli.Context) error {
						query := cmdCluster(c, port, "state")
						fmt.Println(query)
						fmt.Println(getJSON(query))
						return nil
					},
				},
			},
		},
		{
			Name:      "cat",
			ShortName: "c",
			Usage:     "Call cat API",
			Subcommands: []cli.Command{
				{
					Name:    "allocation",
					Aliases: []string{"al"},
					Usage:   "get allocation",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "allocation")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "shards",
					Aliases: []string{"sh"},
					Usage:   "get shards",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "shards")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
