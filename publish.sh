#!/bin/bash

rsync -avC --exclude-from=rsync.excludes ./htdocs/ /var/www/www.sogvin.com/
