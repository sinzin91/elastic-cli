package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"

	prettyjson "github.com/hokaccha/go-prettyjson"
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
			},
		},
	}

	app.Run(os.Args)
}

func getJSON(route string) (string, error) {
	r, err := http.Get(route)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status code: %s", r.Status)
	}

	mediatype, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}
	if mediatype == "" {
		return "", errors.New("mediatype not set")
	}
	if mediatype != "application/json" {
		return "", fmt.Errorf("mediatype is '%s', 'application/json' expected", mediatype)
	}

	var b interface{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return "", err
	}
	out, err := prettyjson.Marshal(b)
	return string(out), err
}

func getRaw(route string) (string, error) {
	r, err := http.Get(route)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status code: %s", r.Status)
	}

	out, err := ioutil.ReadAll(r.Body)
	return string(out), err
}

func cmdCluster(c *cli.Context, port string, subCmd string) string {
	route := "/_cluster"
	url := "http://localhost:"

	var arg string
	switch subCmd {
	case "health":
		arg = "/health"
	case "state":
		arg = "/state"
	default:
		arg = ""
	}
	return url + port + route + arg
}

func cmdCat(c *cli.Context, port string, subCmd string) string {
	route := "/_cat"
	url := c.GlobalString("baseurl")

	var arg string
	switch subCmd {
	case "allocation":
		arg = "/allocation?v"
	}
	return url + port + route + arg
}
