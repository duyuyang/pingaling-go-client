# Pingaling go-client

[![GoDoc](https://godoc.org/github.com/duyuyang/pingaling-go-client/pkg/pingaling?status.svg)](https://godoc.org/github.com/duyuyang/pingaling-go-client/pkg/pingaling)
[![codecov](https://codecov.io/bb/pingaling-monitoring/client/branch/master/graph/badge.svg)](https://codecov.io/bb/pingaling-monitoring/client)
[![ci](https://img.shields.io/bitbucket/pipelines/pingaling-monitoring/client.svg)](https://github.com/duyuyang/pingaling-go-client/addon/pipelines/home#!/)
[![Go Report Card](https://goreportcard.com/badge/github.com/duyuyang/pingaling-go-client)](https://goreportcard.com/report/github.com/duyuyang/pingaling-go-client)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/duyuyang/pingaling-go-client/src/master/LICENSE)


## To start developing

```shell
$ go get -d github.com/duyuyang/pingaling-go-client
$ cd $GOPATH/src/github.com/duyuyang/pingaling-go-client
```

This project use [dep](https://github.com/golang/dep) to manage dependencies.
Run `dep ensure` to add new dependencies.

## CMD document

* [pingaling](doc/pingaling.md) - command line user guide

## Build binary

```shell
$ go build -o pingaling .
$ mv pingaling /usr/local/bin/
$ pingaling -h
```
