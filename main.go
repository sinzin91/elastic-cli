package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"

	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
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
			Name:    "health",
			Aliases: []string{"h"},
			Usage:   "get cluster health",
			Action: func(c *cli.Context) error {
				fmt.Println(getJSON("http://localhost:9200/_cluster/health"))
				return nil
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
