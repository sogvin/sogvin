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
go test -coverprofile /tmp/c.out ./...
uncover /tmp/c.out
browser=$(xdotool search --name "@cgDISVMf")
back=$(xdotool getactivewindow)
xdotool windowactivate --sync $browser
xdotool key --window $browser --clearmodifiers "CTRL+R"
xdotool windowactivate --sync $back
