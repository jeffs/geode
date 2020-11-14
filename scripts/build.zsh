#!/usr/bin/env zsh

set -euo pipefail

cd $(dirname $0)/..
export GOPATH=~/.cache/geode/go
go install ./...
