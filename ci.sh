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
	mkdir -p $dist
	go run ./cmd/mksite -p $dist/docs
	rsync -avC ./docs $dist/
	;;
    clean)
	rm -rf $dist
	;;
    *)
	$0 build
	;;
esac


# Run next target if any
shift
[[ -z "$@" ]] && exit 0
$0 $@

