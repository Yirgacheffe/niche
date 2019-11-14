#!/bin/sh
#
# Build images for all greeting service
#

# --
readonly -a srvs=(a b c d e f g h)
readonly env_prefix=dev

last_commit=$(git rev-parse HEAD | cut -c1-7)

for i in "${srvs[@]}"
do
    cp -f build_srv.sh "service-${i}"
    cp -f Dockerfile   "service-${i}"
    
    pushd "service-${i}"

    # Build go linuxc binary & docket it
    ./build_srv.sh
    docker build -t "soline74/niche-greetings-${i}:${env_prefix}-${last_commit}" --no-cache .
    
    # Clean up
    rm -rf Dockerfile
    rm -rf build_srv.sh
    rm -rf greetings

    popd
done

docker image ls | grep 'soline74/niche-greetings'
