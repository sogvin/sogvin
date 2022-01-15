#!/bin/bash
path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

set -e
set -o pipefail

case $extension in
    go)
        goimports -w $path
        ;;
esac

./ci.sh build test
