#!/bin/bash

set -e
set -o pipefail

dist=/tmp/website

case $1 in
    -h)
	echo "Usage: $0 build|publish"
	;;
    publish)
	go run ./cmd/mksite -c # guard
	rsync -avC $dist/docs/ www.7de.se:/var/www/www.website.com/
	;;
    build)
	pushd ../
	tree -P "*.go" -I "*_test.go" navstar | \
	    grep -v directories > website/example/navstar.tree
	popd
	go build ./...
	
	mkdir -p $dist
	go run ./cmd/mksite -p $dist/docs
	rsync -aC ./docs $dist/
	;;
    clean)
	rm -rf $dist
	;;
    test)
	go test -coverprofile /tmp/c.out ./... 2>&1 | \
	    sed -e 's| of statements||g' \
		-e 's|coverage: ||g' \
		-e 's|github.com/gregoryv/website|.|g' | \
	    grep -v "no test"
	;;
    *)
	$0 build test
	;;
esac


# Run next target if any
shift
[[ -z "$@" ]] && exit 0
$0 $@

