# elastic-cli
A command-line interface for Elasticsearch

## Installation

The following installs the program given Go is [installed](https://golang.org/doc/install) and `$GOPATH` is set:
```
go get -u github.com/sinzin91/elastic-cli
```
Make sure that $GOPATH/bin is in your $PATH.

Optionally, save yourself some key strokes by renaming the app to `es`:
```
mv `which elastic-cli` $GOPATH/bin/es
```

## Usage
```
❯ elastic-cli                                                  
NAME:
   elastic-cli - A Robust CLI for Elasticsearch

USAGE:
   elastic-cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Tenzin Wangdhen <sinzin91@gmail.com>

COMMANDS:
     query, q     Perform any ES API GET query
     nodes, n     Get node information
     cluster, cl  Get cluster information
     cat, c       Call cat API
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --baseurl value  Base API URL (default: "http://localhost:")
   --port value     port that Elasticsearch is listening on (default: "9200")
   --help, -h       show help
   --version, -v    print the version
```