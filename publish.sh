#!/bin/bash

# cleanup
find . -name "*~" -delete

# deploy
rsync -avC ./docs/ www.7de.se:/var/www/www.sogvin.com/
