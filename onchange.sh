#!/bin/bash -e
path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
        goimports -w $path
        ;;
esac
cat cmd/graceful/graceful.go | \
    awk '{if(NR>10)print}' | \
    ud -i graceful -c -w page/graceful_server_shutdown.html

go test -coverprofile /tmp/c.out ./...
uncover /tmp/c.out
