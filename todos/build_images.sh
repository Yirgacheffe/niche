#!/bin/sh
#
# Build image for todo service
#

# --

./build_todos.sh
./docker_it.sh

rm -rf ./niche-todos

docker image ls | grep 'soline74/niche-todos'
