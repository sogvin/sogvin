#!/bin/bash

find . -name "*~" -delete
rsync -avC ./docs/ kard.local:/var/www/www.sogvin.com/
