#!/bin/bash

rsync -avC --exclude-from=rsync.excludes ./www/ /var/www/www.7de.se/go-learn/SE/
