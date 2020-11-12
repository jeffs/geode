#!/usr/bin/env zsh

cd $(dirname $0)/..
export GOPATH=$PWD/go

declare argsfile=testdata/args

declare tmp=$(mktemp -d)
trap "rm -rf $tmp" EXIT

declare a=$tmp/a b=$tmp/b

# Print timestamps of regular files under the current directory.
ts() {
    fd --type=file --exclude=go --no-ignore --exec-batch stat -f %m
}

args() {
    test -f $argsfile && cat $argsfile
}

clear-run() {
    if (( $# > 0 )); then
        local run="./go/bin/geode $@ $(args)"
    else
        local run="./go/bin/geode $(args)"
    fi
    clear
    echo -e "\e[2m[$(date +%T)] go install ./...\e[22m\n"
    if ! go install ./...; then
        return
    fi
    clear
    echo -e "\e[2m[$(date +%T)] $run \e[22m\n"
    ./go/bin/geode "$@" $(args) >& $tmp/Dockerfile
    local code=$?
    bat --pager=never --plain $tmp/Dockerfile
    echo -ne "\n\e[2m[$(date +%T)] $code\e[22m"
}

ts >$a
clear-run "$@"

while true; do
    sleep 0.5
    ts >$b
    if ! cmp --silent $a $b; then
        mv $b $a
        clear-run "$@"
    fi
done
