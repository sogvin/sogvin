#!/bin/bash

rsync -avC --exclude-from=rsync.excludes ./se/ /var/www/www.7de.se/go-learn/SE/
