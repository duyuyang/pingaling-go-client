# Pingaling go-client

[![GoDoc](https://godoc.org/bitbucket.org/pingaling-monitoring/client/pkg/pingaling?status.svg)](https://godoc.org/bitbucket.org/pingaling-monitoring/client/pkg/pingaling)

## To start developing

```shell
$ go get -d bitbucket.org/pingaling-monitoring/client
$ cd $GOPATH/src/bitbucket.org/pingaling-monitoring/client
$ go get ./...              # install dependencies
$ go build -o pingaling .   # build binary
```

## Troubleshooting

Bitbucket access 403

```shell
$ git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

```
