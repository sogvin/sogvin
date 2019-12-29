#!/bin/bash

rsync -avC --exclude-from=rsync.excludes ./page/book/internal/www/ /var/www/www.7de.se/go-learn/SE/
