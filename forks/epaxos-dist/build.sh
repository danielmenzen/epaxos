#!/bin/bash

#go get github.com/stretchr/testify
#go get code.google.com/p/leveldb-go/leveldb
#go get github.com/golang/glog

path=`pwd -P`

echo $path

export GOPATH=$path

#go get github.com/golang/leveldb

go get github.com/syndtr/goleveldb

go build demo/server.go



