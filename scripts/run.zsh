#!/usr/bin/env zsh

set -euo pipefail

cd $(dirname $0)/..
scripts/build.zsh
./go/bin/geode "$@" $(test -f testdata/args && cat testdata/args)
