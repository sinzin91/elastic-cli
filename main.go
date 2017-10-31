package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	var port string

	app := cli.NewApp()
	app.Name = "elastic-cli"
	app.Usage = "A Robust CLI for Elasticsearch"
	app.Version = "1.0.0"
	app.Author = "Tenzin Wangdhen"
	app.Email = "sinzin91@gmail.com"
	app.Flags = []cli.Flag{
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
			Name:      "query",
			ShortName: "q",
			Usage:     "Perform any ES API GET query",
			Action: func(c *cli.Context) {
				var out string
				var err error
				if strings.Contains(c.Args().First(), "_cat/") {
					out, err = getRaw(cmdQuery(c, port))
				} else {
					out, err = getJSON(cmdQuery(c, port))
				}
				if err != nil {
					fatal(err)
				}
				fmt.Println(out)
			},
		},
		{
			Name:      "tasks",
			ShortName: "t",
			Usage:     "get tasks",
			Action: func(c *cli.Context) {
				query := cmdTasks(c, port)
				fmt.Println(query)
				fmt.Println(getJSON(query))
			},
		},
		{
			Name:      "search_shards",
			ShortName: "sh",
			Usage:     "get search_shards",
			Action: func(c *cli.Context) {
				query := cmdSearchShards(c, port)
				fmt.Println(query)
				fmt.Println(getJSON(query))
			},
		},
		{
			Name:      "nodes",
			ShortName: "n",
			Usage:     "Get node information",
			Action: func(c *cli.Context) {
				query := cmdNodes(c, port, "")
				fmt.Println(query)
				fmt.Println(getJSON(query))
			},
			Subcommands: []cli.Command{
				{
					Name:    "stats",
					Aliases: []string{"s"},
					Usage:   "get node stats",
					Action: func(c *cli.Context) {
						query := cmdNodes(c, port, "stats")
						fmt.Println(query)
						fmt.Println(getJSON(query))
					},
				},
				{
					Name:    "hot_threads",
					Aliases: []string{"h"},
					Usage:   "get hot_threads (optionally by node)",
					Action: func(c *cli.Context) {
						query := cmdNodes(c, port, "hot_threads")
						fmt.Println(query)
						fmt.Println(getRaw(query))
					},
				},
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
				{
					Name:    "stats",
					Aliases: []string{"st"},
					Usage:   "get cluster stats",
					Action: func(c *cli.Context) error {
						query := cmdCluster(c, port, "stats")
						fmt.Println(query)
						fmt.Println(getJSON(query))
						return nil
					},
				},
				{
					Name:    "settings",
					Aliases: []string{"se"},
					Usage:   "get cluster settings",
					Action: func(c *cli.Context) error {
						query := cmdCluster(c, port, "settings")
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
					Usage:   "get shards (optionally by index)",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "shards")
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
						query := cmdCat(c, port, "master")
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
						query := cmdCat(c, port, "nodes")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "indices",
					Aliases: []string{"i"},
					Usage:   "get indices (optionally by index)",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "indices")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "segments",
					Aliases: []string{"s"},
					Usage:   "get segments (optionally by index)",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "segments")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "count",
					Aliases: []string{"c"},
					Usage:   "get count (optionally by index)",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "count")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "recovery",
					Aliases: []string{"r"},
					Usage:   "get recovery (optionally by index)",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "recovery")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "health",
					Aliases: []string{"he"},
					Usage:   "get health",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "health")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "pending_tasks",
					Aliases: []string{"pt"},
					Usage:   "get pending_tasks",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "pending_tasks")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "aliases",
					Aliases: []string{"as"},
					Usage:   "get aliases (optionally by index)",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "aliases")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "thread_pool",
					Aliases: []string{"th"},
					Usage:   "get thread_pool",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "thread_pool")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "plugins",
					Aliases: []string{"pl"},
					Usage:   "get plugins",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "plugins")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "fielddata",
					Aliases: []string{"f"},
					Usage:   "get fielddata (optionally by index)",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "fielddata")
						fmt.Print(query + "\n")
						fmt.Println(getRaw(query))
						return nil
					},
				},
				{
					Name:    "tasks",
					Aliases: []string{"t"},
					Usage:   "get tasks",
					Action: func(c *cli.Context) error {
						query := cmdCat(c, port, "tasks")
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
