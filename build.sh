#/bin/bash

path=`pwd -P`

echo $path

export GOPATH=$path

# create binaries

go install master

go install server

go install client

echo "done!"