#!/bin/bash

find . -name "*~" -delete
rsync -avC ./htdocs/ kard.local:/var/www/www.sogvin.com/
