#!/usr/bin/env bash
set -ex

ROOTDIR="$( dirname "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" )"
BINDIR=$ROOTDIR/bin

export GOPATH=$ROOTDIR
export GOOS=linux

apt-get install build-essential
apt-get install cmake git libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev

go build -ldflags="-s -w" -o $BINDIR/supply nodejs/supply/cli
go build -ldflags="-s -w" -o $BINDIR/finalize nodejs/finalize/cli
