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
						query := cmdCat(c, port, "allocation", "")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "shards",
					Aliases: []string{"sh"},
					Usage:   "get shards (optionally by index)",
					Action: func(c *cli.Context) error {
						index := ""
						if c.Args().Get(0) != "" {
							index = c.Args().Get(0)
						}
						query := cmdCat(c, port, "shards", index)
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "master",
					Aliases: []string{"m"},
					Usage:   "get master",
					Action: func(c *cli.Context) error {
						index := ""
						query := cmdCat(c, port, "master", index)
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "nodes",
					Aliases: []string{"n"},
					Usage:   "get nodes",
					Action: func(c *cli.Context) error {
						index := ""
						query := cmdCat(c, port, "nodes", index)
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "indices",
					Aliases: []string{"i"},
					Usage:   "get indices",
					Action: func(c *cli.Context) error {
						index := ""
						if c.Args().Get(0) != "" {
							index = c.Args().Get(0)
						}
						query := cmdCat(c, port, "indices", index)
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "segments",
					Aliases: []string{"s"},
					Usage:   "get segments",
					Action: func(c *cli.Context) error {
						index := ""
						if c.Args().Get(0) != "" {
							index = c.Args().Get(0)
						}
						query := cmdCat(c, port, "segments", index)
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "count",
					Aliases: []string{"c"},
					Usage:   "get count",
					Action: func(c *cli.Context) error {
						index := ""
						if c.Args().Get(0) != "" {
							index = c.Args().Get(0)
						}
						query := cmdCat(c, port, "count", index)
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
