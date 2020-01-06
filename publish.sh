#!/bin/bash

find . -name "*~" -delete
rsync -avC ./htdocs/ /var/www/www.sogvin.com/
