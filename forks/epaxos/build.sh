#!/bin/bash


path=`pwd -P`

echo $path

export GOPATH=$path

go get github.com/stretchr/testify
go get github.com/boltdb/bolt/...
go get github.com/golang/glog
