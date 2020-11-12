#!/usr/bin/env zsh

set -uo pipefail

local flat=geode-testdata-groovy
local name=geode/testdata-groovy

if docker image inspect $name >&/dev/null; then
    docker rmi --no-prune $name
fi

if docker volume inspect $flat >&/dev/null; then
    docker volume rm $flat
fi
