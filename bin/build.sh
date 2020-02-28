#!/bin/bash

export PATH=$PATH:/usr/local/go/bin
export GO111MODULE=on

if [ ! -f go.mod ]; then
    go mod init gzh
fi
go mod tidy
go build -o ./bin/gzh