#!/usr/bin/env zsh

set -euo pipefail

cd $(dirname $0)/..
export GOPATH=$PWD/go
go install ./...
