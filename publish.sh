#!/bin/bash

find . -name "*~" -delete
rsync -avC ./htdocs/ warmachine.local:/var/www/www.sogvin.com/
