#!/bin/sh
#
# Build images for all greeting service
#
# readonly -a srvs=(a b c d e f g h)

readonly -a srvs=(e)
# readonly env=dev

last_commit=$(git rev-parse HEAD | cut -c1-7)

for i in "${srvs[@]}"
do
    cp -f Dockerfile "service-${i}"
    pushd "service-${i}"
    docker build -t "soline74/niche-greeting-${i}:${last_commit}" --no-cache .
    rm -rf Dockerfile
    popd
done

docker image ls | grep 'soline74/niche-greeting'
