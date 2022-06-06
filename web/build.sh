#!/bin/sh
#
# Build images for all greeting service
#

# --
./build_web.sh
./docker_it.sh

rm -rf ./niche-web

docker image ls | grep 'soline74/niche-web'
