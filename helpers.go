package main

import (
	"encoding/json"
	"errors"
  "fmt"
	"os"
	"io/ioutil"
	"mime"
	"net/http"
	
	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/urfave/cli"
)

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
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
	case "stats":
		arg = "/stats"
	default:
		arg = ""
	}
	return url + port + route + arg
}

func cmdCat(c *cli.Context, port string, subCmd string, index string) string {
	route := "/_cat"
	url := c.GlobalString("baseurl")

	var arg string
	switch subCmd {
	case "allocation":
		arg = "/allocation"
	case "shards":
		arg = "/shards/"
	case "master":
		arg = "/master"
	case "nodes":
		arg = "/nodes"
	case "indices":
		arg = "/indices/"
	case "segments":
		arg = "/segments/"
	case "count":
		arg = "/count/"
	case "recovery":
		arg = "/recovery/"
	case "health":
		arg = "/health/"
	case "pending_tasks":
		arg = "/pending_tasks/"
	case "aliases":
		arg = "/aliases/"
	case "thread_pool":
		arg = "/thread_pool/"
	case "plugins":
		arg = "/plugins/"
	case "fielddata":
		arg = "/fielddata/"
	}
	return url + port + route + arg + index + "?v"
}

func cmdQuery(c *cli.Context, port string) string {
	route := c.Args().First()
	url := c.GlobalString("baseurl")
	return url + port + route
}