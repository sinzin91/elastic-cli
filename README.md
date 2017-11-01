# elastic-cli
A command-line interface for Elasticsearch

## Installation

The following installs the program given Go is [installed](https://golang.org/doc/install) and `$GOPATH` is set:
```
go get -u github.com/sinzin91/elastic-cli
```
Make sure that `$GOPATH/bin` is in your `$PATH`.

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
     cat, c             Call cat API
     cluster, cl        Get cluster information
     nodes, n           Get node information
     query, q           Perform any ES API GET query
     recovery, re       Get recovery
     search_shards, sh  Get search_shards
     settings, se       Get index settings
     stats, st          Get index stats
     tasks, t           Get tasks
     help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --baseurl value  Base API URL (default: "http://localhost:")
   --port value     port that Elasticsearch is listening on (default: "9200")
   --help, -h       show help
   --version, -v    print the version
```

## Examples

### Cat nodes
```
❯ elastic-cli --port "6666" c nodes
http://localhost:6666/_cat/nodes?v
host                       ip         heap.percent ram.percent load node.role master name
ip-10-0-1.555.ec2.internal 10.0.1.555           39          53 0.65 d         -      esd2
ip-10-0-1.555.ec2.internal 10.0.1.555            6          33 0.06 -         m      esm0
ip-10-0-1.555.ec2.internal 10.0.1.555           11          93 0.07 -         *      esm1
ip-10-0-1.555.ec2.internal 10.0.1.555           11          54 1.54 d         -      esd0
ip-10-0-1.555.ec2.internal 10.0.1.555           15          33 0.00 -         -      esc0
ip-10-0-1.555.ec2.internal 10.0.1.555           11          54 1.58 d         -      esd3
```

### Stats for specific node
```
❯ elastic-cli nodes stats esd2
http://localhost:9200/_nodes/esd2/stats
{
  "cluster_name": "elasticsearch_cli",
  "nodes": {
    "t812hutougccaHU": {
      "attributes": {
        "aws_availability_zone": "us-west-1d",
        "master": "false",
        "max_local_storage_nodes": "1"
      },
      ...
```

### Query
```
❯ elastic-cli q '/twitter/user/test_user'
{
  "_id": "test_user",
  "_index": "twitter",
  "_source": {
    "age": "21",
    "email": "test_user@gmail.com",
  },
  "_type": "user",
  "_version": 85,
  "found": true
}
```
