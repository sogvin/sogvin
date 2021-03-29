#!/bin/bash

case $1 in
    publish)
	# cleanup
	find . -name "*~" -delete

	# deploy
	rsync -avC ./docs/ www.7de.se:/var/www/www.sogvin.com/
	;;
    *)
	echo "Usage: $0 deploy"
	;;
esac
