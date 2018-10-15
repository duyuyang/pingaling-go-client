# Pingaling go-client

## To start developing

```shell
$ go get -d bitbucket.org/pingaling-monitoring/client
$ cd $GOPATH/src/bitbucket.org/pingaling-monitoring/client
$ go get ./...              # instal dependencies
$ go build -o pingaling .   # build binary
```

## Troubleshooting

Bitbucket access 403

```shell
$ git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

```
