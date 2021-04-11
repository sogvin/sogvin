#!/bin/bash
path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"


pushd ../
tree -P "*.go" -I "*_test.go" navstar | grep -v directories > sogvin/example/navstar.tree
popd

case $extension in
    go)
        goimports -w $path
	go test -coverprofile /tmp/c.out ./... 
	#uncover /tmp/c.out
        ;;
esac

./ci.sh build
