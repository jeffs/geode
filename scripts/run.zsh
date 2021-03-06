#!/usr/bin/env zsh

set -euo pipefail

declare argsfile=testdata/args

cd $(dirname $0)/..
scripts/build.zsh
./go/bin/geode "$@" $(test -f $argsfile && cat $argsfile)
