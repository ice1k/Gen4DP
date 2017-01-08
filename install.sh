#!/usr/bin/env bash

echo compiling...

if [ ! -f install ]; then
echo 'install must be run within its container folder' 1>&2
exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w ./

cd src

go install -o Gen4DP
mv Gen4DP ../

cd ..

export GOPATH="$OLDGOPATH"

echo 'finished'
