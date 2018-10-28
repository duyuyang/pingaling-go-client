# Pingaling go-client

[![GoDoc](https://godoc.org/bitbucket.org/pingaling-monitoring/client/pkg/pingaling?status.svg)](https://godoc.org/bitbucket.org/pingaling-monitoring/client/pkg/pingaling)
[![codecov](https://codecov.io/bb/pingaling-monitoring/client/branch/master/graph/badge.svg)](https://codecov.io/bb/pingaling-monitoring/client)
[![ci](https://img.shields.io/bitbucket/pipelines/pingaling-monitoring/client.svg)](https://bitbucket.org/pingaling-monitoring/client/addon/pipelines/home#!/)
[![Go Report Card](https://goreportcard.com/badge/bitbucket.org/pingaling-monitoring/client)](https://goreportcard.com/report/bitbucket.org/pingaling-monitoring/client)

## To start developing

```shell
$ go get -d bitbucket.org/pingaling-monitoring/client
$ cd $GOPATH/src/bitbucket.org/pingaling-monitoring/client
```

## CMD document

* [pingaling](doc/pingaling.md) - command line user guide

## Build binary

```shell
$ go build -o pingaling .
$ mv pingaling /usr/local/bin/
$ pingaling -h
```

## Troubleshooting

Bitbucket access 403

```shell
$ git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

```
