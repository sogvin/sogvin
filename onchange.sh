#!/bin/bash -e
path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
        goimports -w $path
        gofmt -w $path
        ;;
esac

cat testing/inline_test.go | awk '{if(NR>6)print}' | ud -i inlinetest -c -w page/inline_test_helpers.html
go test -coverprofile /tmp/c.out ./...
uncover /tmp/c.out
