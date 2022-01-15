#!/bin/bash

set -e
set -o pipefail

dist=/tmp/sogvin

case $1 in
    -h)
	echo "Usage: $0 build|publish"
	;;
    publish)
	rsync -avC $dist/docs/ www.7de.se:/var/www/www.sogvin.com/
	;;
    build)
	pushd ../
	tree -P "*.go" -I "*_test.go" navstar | grep -v directories > sogvin/example/navstar.tree
	popd
	go build ./...
	
	mkdir -p $dist
	go run ./cmd/mksite -p $dist/docs
	rsync -avC ./docs $dist/
	;;
    clean)
	rm -rf $dist
	;;
    test)
	go test -coverprofile /tmp/c.out ./... 2>&1 | sed 's|github.com/gregoryv|.|g'
	;;
    *)
	$0 build test
	;;
esac


# Run next target if any
shift
[[ -z "$@" ]] && exit 0
$0 $@

